# Changelog

## [[v0.0.2]](https://github.com/mlange-42/arche/compare/v0.0.1...v0.0.2)

### Breaking changes

* Resource `Time` is now split into `Tick` and `Termination` (#16)

### Features

* Adds a system `CallbackTermination` to end the simulation based on a callback (#13)

### Bugfixes

* Fix check when removing a system that is not in `Systems` (#15)

### Other

* Systems are removed immediately when `Systems.RemoveSystem` is called outside of a loop over systems (#15)
* Included systems do no longer depend on resource `Tick` (formerly `Time`) (#16)

### Documentation

* Improves examples with inline comments (#9)
* Adds a CHANGELOG.md file (#9)
* Adds examples for implementing `System` and `UISystem` (#10)
