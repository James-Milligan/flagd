package dapr

import (
	"context"
	"fmt"

	daprSDK "github.com/dapr/go-sdk/client"
	"github.com/open-feature/flagd/core/pkg/logger"
	"github.com/open-feature/flagd/core/pkg/sync"
	isync "github.com/open-feature/flagd/core/pkg/sync"
)

type Sync struct {
	client daprSDK.Client
	ready  bool

	URI       string
	StoreName string
	Logger    *logger.Logger
	Address   string
}

func (d *Sync) Init(ctx context.Context) error {
	var client daprSDK.Client
	var err error
	if d.Address != "" {
		client, err = daprSDK.NewClientWithAddress(d.Address)
	} else {
		client, err = daprSDK.NewClient()
	}
	if err != nil {
		return err
	}

	d.client = client
	d.ready = true

	return nil
}

func (d *Sync) Sync(ctx context.Context, dataSync chan<- isync.DataSync) error {
	// initial fetch
	config, err := d.fetchConfig(ctx)
	if err != nil {
		return err
	}
	dataSync <- sync.DataSync{
		FlagData: config,
		Source:   fmt.Sprintf("%s/%s", d.StoreName, d.URI),
		Type:     sync.ALL,
	}

	// subscription handle
	err = d.client.SubscribeConfigurationItems(ctx, d.StoreName, []string{d.URI}, func(id string, config map[string]*daprSDK.ConfigurationItem) {
		if len(config) == 0 {
			d.Logger.InfoWithID(id, "subscribed to config changes")
			return
		}
		d.Logger.InfoWithID(id, "configuration change detected")
		conf, ok := config[d.URI]
		if !ok {
			d.Logger.InfoWithID(id, fmt.Sprintf("URI %s missing from configuration change event", d.URI))
			return
		}
		dataSync <- sync.DataSync{
			FlagData: conf.Value,
			Source:   fmt.Sprintf("%s/%s", d.StoreName, d.URI),
			Type:     sync.ALL,
		}
	})
	if err != nil {
		return fmt.Errorf("unable to subscribe to configuration %s/%s, err: %w", d.StoreName, d.URI, err)
	}
	<-ctx.Done()
	return nil
}

func (d *Sync) fetchConfig(ctx context.Context) (string, error) {
	config, err := d.client.GetConfigurationItem(ctx, d.StoreName, d.URI)
	if err != nil {
		err = fmt.Errorf("Could not get config item, err: %w", err)
		return "", err
	}
	return string(config.Value), nil
}

func (d *Sync) ReSync(ctx context.Context, dataSync chan<- isync.DataSync) error {
	config, err := d.fetchConfig(ctx)
	if err != nil {
		return err
	}
	dataSync <- sync.DataSync{
		FlagData: config,
		Source:   fmt.Sprintf("%s/%s", d.StoreName, d.URI),
		Type:     sync.ALL,
	}
	return nil
}

func (d *Sync) IsReady() bool {
	return d.ready
}
