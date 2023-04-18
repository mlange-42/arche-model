# Changelog

## [[v0.3.0]](https://github.com/mlange-42/arche-model/compare/v0.2.0...v0.3.0)

### Breaking changes

* Methods `observer.Grid.X` and `observer.Grid.Y` take an `int` argument and return one value instead of all (#38)

### Features

* Observer `RowToTable` as adapter from `Row` to `Table` observer (#37)
* Observer `MatrixToGrid` as adapter from `Matrix` to `Grid` observer (#38)
* System `PerfTimer` prints total step and average time per tick on finalization (#37)

## [[v0.2.0]](https://github.com/mlange-42/arche-model/compare/v0.1.0...v0.2.0)

### Features

* Adds a resource `SelectedEntity` for communication between UI systems, e.g. for entity inspection or manipulation by the user (#34)

### Documentation

* Extends documentation on resources (#34)
* Adds a list of features and a usage example to the README (#35)

### Other

* Upgrade to Arche v0.7.0 (#36)

## [[v0.1.0]](https://github.com/mlange-42/arche-model/compare/v0.0.5...v0.1.0)

### Other

* Brings test coverage to 95%, adds test coverage badge (#30)
* More precise TPS when simulation does not reach target TPS (#31)
* Get rid of hot loop for waiting small amounts of time (#32)
* Upgrade to Arche v0.6.3 (#33)
* Promote to v0.1.0 to reflect increased API stability (#33)

## [[v0.0.5]](https://github.com/mlange-42/arche-model/compare/v0.0.4...v0.0.5)

### Breaking changes

* Renamed `Systems.Fps` and `Systems.Tps` to `Systems.FPS` and `Systems.TPS` (#26)

### Features

* Simulations can be paused through the `Systems` resource (#25)

### Other

* Unset/zero `Model.FPS` sets to 30 FPS, as a default more useful than synced with TPS (#27)

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
