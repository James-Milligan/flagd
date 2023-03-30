# Changelog

## [0.5.0](https://github.com/james-milligan/flagd/compare/core-v0.4.5...core/v0.5.0) (2023-03-30)


### ⚠ BREAKING CHANGES

* package refactor
* unify sources configuration handling ([#560](https://github.com/james-milligan/flagd/issues/560))

### 🧹 Chore

* move credential builder for grpc sync into seperate component ([#536](https://github.com/james-milligan/flagd/issues/536)) ([7314fee](https://github.com/james-milligan/flagd/commit/7314feea8c7bc90aac0528a9e1be0759a7a60c15))
* refactor configuration handling for startup ([#551](https://github.com/james-milligan/flagd/issues/551)) ([8dfbde5](https://github.com/james-milligan/flagd/commit/8dfbde5bbffd16fb66797a750d15f0226edf54a7))
* refactor middleware setup in server ([#554](https://github.com/james-milligan/flagd/issues/554)) ([01016c7](https://github.com/james-milligan/flagd/commit/01016c7df7c5f653cdadf151539432f692f36251))
* refactor service configuration objects ([#545](https://github.com/james-milligan/flagd/issues/545)) ([c7b29ed](https://github.com/james-milligan/flagd/commit/c7b29edcfe9dab61eaa585011a690d47829601b6)), closes [#524](https://github.com/james-milligan/flagd/issues/524)
* release main ([#488](https://github.com/james-milligan/flagd/issues/488)) ([873c7b3](https://github.com/james-milligan/flagd/commit/873c7b37dbbfefa8f6448fb1dec06b2ce40fb02d))
* release main ([#512](https://github.com/james-milligan/flagd/issues/512)) ([85a0f0d](https://github.com/james-milligan/flagd/commit/85a0f0d4d97d9388dcd291b2a96e8bbbdd53c2a7))
* unify sources configuration handling ([#560](https://github.com/james-milligan/flagd/issues/560)) ([7f4888a](https://github.com/james-milligan/flagd/commit/7f4888a1676e49acecf328685623566ea057ffcf))


### 🐛 Bug Fixes

* **deps:** update kubernetes packages to v0.26.3 ([#533](https://github.com/james-milligan/flagd/issues/533)) ([6ddd5b2](https://github.com/james-milligan/flagd/commit/6ddd5b29806f3101cf122bfc4196ba7d0ef4c025))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.31 ([#527](https://github.com/james-milligan/flagd/issues/527)) ([7928130](https://github.com/james-milligan/flagd/commit/7928130b10906b10f4501630f16a71bdd8e4e365))
* **deps:** update module google.golang.org/grpc to v1.54.0 ([#548](https://github.com/james-milligan/flagd/issues/548)) ([99ba5ec](https://github.com/james-milligan/flagd/commit/99ba5ece76d98124c108bc6280bee03a5c0cd25d))
* **deps:** update module google.golang.org/protobuf to v1.29.1 [security] ([#504](https://github.com/james-milligan/flagd/issues/504)) ([59db0bb](https://github.com/james-milligan/flagd/commit/59db0bba43a9c002378fdced2fcf4729d619e28b))
* **deps:** update module google.golang.org/protobuf to v1.30.0 ([#499](https://github.com/james-milligan/flagd/issues/499)) ([f650338](https://github.com/james-milligan/flagd/commit/f650338e01e721a9d24e2ed6f6fe585dbb6beb42))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.14.6 ([#572](https://github.com/james-milligan/flagd/issues/572)) ([bed9458](https://github.com/james-milligan/flagd/commit/bed94584a30bb6752284ece152cb114a102cbe8a))
* fixing silent lint failures ([#550](https://github.com/james-milligan/flagd/issues/550)) ([30c8022](https://github.com/james-milligan/flagd/commit/30c8022e891d1d278c096dd2438137ced7552678))
* nil pointer fix + export constructors ([#555](https://github.com/james-milligan/flagd/issues/555)) ([78adb81](https://github.com/james-milligan/flagd/commit/78adb81f4eb7a5b7fdb7075fa0bf8afa6d03dc72))


### ✨ New Features

* expose Impression metric ([#556](https://github.com/james-milligan/flagd/issues/556)) ([77e0a33](https://github.com/james-milligan/flagd/commit/77e0a33be24dcd0b6e239e5ed709167167c14171))
* grpc connection options to flagd configuration options ([#532](https://github.com/james-milligan/flagd/issues/532)) ([aa74951](https://github.com/james-milligan/flagd/commit/aa74951f43b662ff2df53e68d347fc10e6d23bb8))
* Introduce flagd kube proxy ([#495](https://github.com/james-milligan/flagd/issues/495)) ([440864c](https://github.com/james-milligan/flagd/commit/440864ce87174618321c9d5146221490d8f07b24))
* Introduce kube-proxy-metrics ([#558](https://github.com/james-milligan/flagd/issues/558)) ([ad0baeb](https://github.com/james-milligan/flagd/commit/ad0baeb08fa67c94356d6a3f298283373bd5211b))
* package refactor ([b765684](https://github.com/james-milligan/flagd/commit/b7656845a80c4fd5788c03d2c36c1e23201c0275))
* refactor core module into multiple packages ([#530](https://github.com/james-milligan/flagd/issues/530)) ([9d68d0b](https://github.com/james-milligan/flagd/commit/9d68d0b45815facdf6079ffcd7864f720ccb8475))
* Restructure for monorepo setup ([#486](https://github.com/james-milligan/flagd/issues/486)) ([ed2993c](https://github.com/james-milligan/flagd/commit/ed2993cd67b8a46db3beb6bb8a360e1aa20349da))

## [0.4.5](https://github.com/open-feature/flagd/compare/core/v0.4.4...core/v0.4.5) (2023-03-20)


### 🐛 Bug Fixes

* **deps:** update kubernetes packages to v0.26.3 ([#533](https://github.com/open-feature/flagd/issues/533)) ([6ddd5b2](https://github.com/open-feature/flagd/commit/6ddd5b29806f3101cf122bfc4196ba7d0ef4c025))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.31 ([#527](https://github.com/open-feature/flagd/issues/527)) ([7928130](https://github.com/open-feature/flagd/commit/7928130b10906b10f4501630f16a71bdd8e4e365))
* **deps:** update module google.golang.org/protobuf to v1.29.1 [security] ([#504](https://github.com/open-feature/flagd/issues/504)) ([59db0bb](https://github.com/open-feature/flagd/commit/59db0bba43a9c002378fdced2fcf4729d619e28b))
* **deps:** update module google.golang.org/protobuf to v1.30.0 ([#499](https://github.com/open-feature/flagd/issues/499)) ([f650338](https://github.com/open-feature/flagd/commit/f650338e01e721a9d24e2ed6f6fe585dbb6beb42))


### ✨ New Features

* grpc connection options to flagd configuration options ([#532](https://github.com/open-feature/flagd/issues/532)) ([aa74951](https://github.com/open-feature/flagd/commit/aa74951f43b662ff2df53e68d347fc10e6d23bb8))
* Introduce flagd kube proxy ([#495](https://github.com/open-feature/flagd/issues/495)) ([440864c](https://github.com/open-feature/flagd/commit/440864ce87174618321c9d5146221490d8f07b24))

## [0.4.4](https://github.com/open-feature/flagd/compare/core-v0.4.3...core/v0.4.4) (2023-03-10)


### ✨ New Features

* Restructure for monorepo setup ([#486](https://github.com/open-feature/flagd/issues/486)) ([ed2993c](https://github.com/open-feature/flagd/commit/ed2993cd67b8a46db3beb6bb8a360e1aa20349da))
