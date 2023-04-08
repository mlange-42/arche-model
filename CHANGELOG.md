# Changelog

## [[v0.0.4]](https://github.com/mlange-42/arche-model/compare/v0.0.4...v0.0.5)

### Features

* Simulations can be paused through the `Systems` resource (#25)

## [[v0.0.4]](https://github.com/mlange-42/arche-model/compare/v0.0.3...v0.0.4)

### Other

* Precise (average) FPS and TPS timing by using semi-cumulative time (#24)

## [[v0.0.3]](https://github.com/mlange-42/arche-model/compare/v0.0.2...v0.0.3)

### Breaking changes

* All observers moved to separate `observer` package (#20)
* Renamed `Observer` to `observer.Row` and `MatrixObserver` to `observer.Table` (#20)
* Add new `observer.Matrix` and `observer.Grid` for matrices and grids (#20)
* Observer methods like `Header` don't take a `*ecs.World` argument (#20)

### Features

* `Model.Seed()` returns the receiver's pointer to allow for method chaining (#22)

### Documentation

* Extend documentation on `Model`, `Systems` and observers (#18)
* Adds full implementation examples for all observer interfaces (#21)

## [[v0.0.2]](https://github.com/mlange-42/arche-model/compare/v0.0.1...v0.0.2)

### Breaking changes

* All resources moved to package `resource` (#16)
* Resource `Time` is now split into `Tick` and `Termination` (#16)

### Features

* Adds a system `CallbackTermination` to end the simulation based on a callback (#13)

### Bugfixes

* Fix check when removing a system that is not in `Systems` (#15)

### Documentation

* Improves examples with inline comments (#9)
* Adds a CHANGELOG.md file (#9)
* Adds examples for implementing `System` and `UISystem` (#10)

### Other

* Systems are removed immediately when `Systems.RemoveSystem` is called outside of a loop over systems (#15)
* Included systems do no longer depend on resource `Tick` (formerly `Time`) (#16)
* Upgrade to dependency to Arche v0.6.1 (#16)
