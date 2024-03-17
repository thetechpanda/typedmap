# changelog

# v0.0.1

Initial release

# v1.0.0

After some bench testing and comments from peers, it was clear that the use of the mutex to guarantee exclusive access to the map was equivalent (and in some cases worse) to using a map and a RWMutex.
The code has been changed to use this approach.

More extensive testing is introduced and new benchmarks have been added to verify these cases.