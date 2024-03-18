# changelog

# v0.0.1

Initial release

# v1.0.0 Stable

After some bench testing and comments from peers, it was clear that the use of the mutex to guarantee exclusive access to the map was equivalent (and in some cases worse) to using a map and a RWMutex.
The code has been changed to use this approach.

More extensive testing is introduced and new benchmarks have been added to verify these cases.

# v1.1.0 SyncMap + Code Restructure

Introduces typed SyncMap, a generic wrapped `sync.Map`.
Code has been restructured, tests and benchmark have been changed to run the same tests for both `TypedMap` and `SyncMap`.

Go version has been changed from `1.22.0` to `1.22`.