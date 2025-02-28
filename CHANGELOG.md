# Changelog

## [[v0.10.0]](https://github.com/mlange-42/arche-model/compare/v0.9.0...v0.10.0)

### Breaking changes

* Changes parameters of constructor `model.New` (#68)

### Other

* Upgrades Arche to v0.15.0 (#68)

## [[v0.9.0]](https://github.com/mlange-42/arche-model/compare/v0.8.1...v0.9.0)

### Breaking changes

* Renames `reporter.Callback` to `reporter.RowCallback` (#66)

### Features

* Adds option `Final` to `reporter.RowCallback` to report data once on finalize instead of on ticks (#65)
* Adds `reporter.TableCallback` for direct retrieval of table observer output in Go code (#66)

## [[v0.8.1]](https://github.com/mlange-42/arche-model/compare/v0.8.0...v0.8.1)

### Other

* Improved float formatting in CSV output (#64)

## [[v0.8.0]](https://github.com/mlange-42/arche-model/compare/v0.7.0...0.8.0)

### Features

* Adds `reporter.Callback` for direct retrieval of row observer output in Go code (#61)

### Bugfixes

* Fix typo in error message when adding UI system as normal system (#60)
* Fix reporters did not work with unspecified `UpdateInterval` (#61)

### Other

* Upgrade to Arche v0.12.0 (#60)

## [[v0.7.0]](https://github.com/mlange-42/arche-model/compare/v0.6.0...v0.7.0)

### Features

* The model can be stepped manually, instead of relying on `Model.Run()` (#52, #53).

## [[v0.6.0]](https://github.com/mlange-42/arche-model/compare/v0.5.0...v0.6.0)

### Breaking changes

* Upgrade to Arche v0.10.0 (#51)

## [[v0.5.0]](https://github.com/mlange-42/arche-model/compare/v0.4.1...v0.5.0)

### Infrastructure

* Upgrade to Go 1.21 and Arche 0.9.0 (#48)

## [[v0.4.1]](https://github.com/mlange-42/arche-model/compare/v0.4.0...v0.4.1)

### Bugfixes

* Fixes `Systems` spinning at 100% CPU load despite low TPS and FPS (#47, see mlange-42/arche#304)

## [[v0.4.0]](https://github.com/mlange-42/arche-model/compare/v0.3.1...v0.4.0)

### Breaking changes

* Upgrade to Arche 0.8 (#44, #46)

### Features

* Resource `model.Systems` has methods to get the list of systems, for inspection (#45)

## [[v0.3.1]](https://github.com/mlange-42/arche-model/compare/v0.3.0...v0.3.1)

### Other

* Increased time precision on Windows for more consistent TPS and FPS (#42)

## [[v0.3.0]](https://github.com/mlange-42/arche-model/compare/v0.2.0...v0.3.0)

### Breaking changes

* Methods `observer.Grid.X` and `observer.Grid.Y` take an `int` argument and return one value instead of all (#38)

### Features

* Observer constructor `RowToTable` as adapter from `Row` to `Table` observer (#37)
* Observer constructor `MatrixToGrid` as adapter from `Matrix` to `Grid` observer (#38)
* Observer interfaces `MatrixLayers` and `GridLayers` for multi-layered matrices and grids (#39)
* Observer constructors `MatrixToLayers`, `GridToLayers` and `LayersToLayers` as adapters (#39)
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
