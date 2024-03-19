# Benchmarks
The benchmarks aim to compare the performance of `TypedMap` with that of `sync.Map`. Since `sync.Map` is not typed, the benchmarks focus on comparing the performance of operations that are common to both maps.

The following command is used to run the tests

```bash
go test -cpu=4 -bench=. -benchmem ./benchmarks/...
```

### Concurrent Benchmarks
Concurrent benchmark use all the same function body, so that each bench has the behaviour except for TypedMap and sync.Map operations.

Check `benchmarkConcurrentInt` in [benchmarks/benchmarks.go](benchmarks/benchmarks.go) for how the concurrent benchmark is implemented.

## Apple M2 Max 

```
goos: darwin
goarch: arm64
pkg: github.com/thetechpanda/typedmap/benchmarks
BenchmarkNativeSyncMapStoreAndDelete-4          	 6290401	       288.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkNativeSyncMapRange-4                   	18979533	       101.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkNativeSyncMapLoad-4                    	10645312	       248.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkNativeSyncMapSimulateEntries-4         	19292758	       207.3 ns/op	      99 B/op	       0 allocs/op
BenchmarkNativeSyncMapSimulateKeys-4            	23044294	       127.1 ns/op	      41 B/op	       0 allocs/op
BenchmarkNativeSyncMapSimulateValues-4          	10964636	       124.2 ns/op	      44 B/op	       0 allocs/op
BenchmarkNativeSyncMapSimulateUpdate-4          	 4219632	       356.4 ns/op	      31 B/op	       2 allocs/op
BenchmarkNativeSyncMapSimulateUpdateRange-4     	 4276798	       340.5 ns/op	      31 B/op	       2 allocs/op
BenchmarkNativeSyncMapConcurrentOperations-4    	    2589	    434747 ns/op	  223384 B/op	   10297 allocs/op
BenchmarkNativeSyncMapConcurrentStore-4         	     626	   2007323 ns/op	  175977 B/op	   10200 allocs/op
BenchmarkNativeSyncMapConcurrentSwap-4          	     535	   2015166 ns/op	  175978 B/op	   10200 allocs/op
BenchmarkNativeSyncMapConcurrentLoadOrStore-4   	    3452	    343587 ns/op	  172937 B/op	    9439 allocs/op
BenchmarkTypedSyncMapStoreAndDelete-4           	 8171257	       219.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedSyncMapRange-4                    	22390948	       145.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedSyncMapLoad-4                     	 6585390	       306.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedSyncMapSimulateEntries-4          	11433235	       211.6 ns/op	      86 B/op	       0 allocs/op
BenchmarkTypedSyncMapSimulateKeys-4             	15681337	       102.0 ns/op	      49 B/op	       0 allocs/op
BenchmarkTypedSyncMapSimulateValues-4           	23475115	        95.68 ns/op	      40 B/op	       0 allocs/op
BenchmarkTypedSyncMapSimulateUpdate-4           	 3557814	       492.6 ns/op	      31 B/op	       2 allocs/op
BenchmarkTypedSyncMapSimulateUpdateRange-4      	 4552926	       383.6 ns/op	      31 B/op	       2 allocs/op
BenchmarkTypedSyncMapConcurrentOperations-4     	    2198	    561213 ns/op	  176004 B/op	   10200 allocs/op
BenchmarkTypedSyncMapConcurrentStore-4          	     559	   1955571 ns/op	  175960 B/op	   10200 allocs/op
BenchmarkTypedSyncMapConcurrentSwap-4           	     625	   1955117 ns/op	  175970 B/op	   10200 allocs/op
BenchmarkTypedSyncMapConcurrentLoadOrStore-4    	    3680	    371850 ns/op	  166656 B/op	    9448 allocs/op
BenchmarkTypedMapStoreAndDelete-4               	11636006	       189.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedMapRange-4                        	88322694	        14.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedMapLoad-4                         	22004098	       124.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedMapEntries-4                      	85618215	        15.04 ns/op	      16 B/op	       0 allocs/op
BenchmarkTypedMapKeys-4                         	86954680	        13.76 ns/op	       8 B/op	       0 allocs/op
BenchmarkTypedMapValues-4                       	87782664	        13.75 ns/op	       8 B/op	       0 allocs/op
BenchmarkTypedMapUpdate-4                       	10347079	       323.5 ns/op	      16 B/op	       1 allocs/op
BenchmarkTypedMapUpdateRange-4                  	53616708	        27.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedMapConcurrentOperations-4         	     235	   5107732 ns/op	   15981 B/op	     200 allocs/op
BenchmarkTypedMapConcurrentStore-4              	     524	   2421942 ns/op	   15936 B/op	     200 allocs/op
BenchmarkTypedMapConcurrentSwap-4               	     516	   2400017 ns/op	   15939 B/op	     200 allocs/op
BenchmarkTypedMapConcurrentLoadOrStore-4        	     417	   2957191 ns/op	   15954 B/op	     200 allocs/op
BenchmarkTypedMapConcurrentUpdate-4             	     541	   2407494 ns/op	   15954 B/op	     200 allocs/op
PASS
ok  	github.com/thetechpanda/typedmap/benchmarks	389.778s
```

## Intel(R) Xeon(R) CPU E5-2678 v3 @ 2.50GHz

```
goos: linux
goarch: amd64
pkg: github.com/thetechpanda/typedmap/benchmarks
cpu: Intel(R) Xeon(R) CPU E5-2678 v3 @ 2.50GHz
BenchmarkNativeSyncMapStoreAndDelete-4          	 6117129	       249.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkNativeSyncMapRange-4                   	12032986	       149.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkNativeSyncMapLoad-4                    	 5939712	       235.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkNativeSyncMapSimulateEntries-4         	 6360204	       370.8 ns/op	      98 B/op	       0 allocs/op
BenchmarkNativeSyncMapSimulateKeys-4            	 9345627	       168.9 ns/op	      42 B/op	       0 allocs/op
BenchmarkNativeSyncMapSimulateValues-4          	 6252608	       216.8 ns/op	      40 B/op	       0 allocs/op
BenchmarkNativeSyncMapSimulateUpdate-4          	 3307431	       390.0 ns/op	      31 B/op	       2 allocs/op
BenchmarkNativeSyncMapSimulateUpdateRange-4     	 3292473	       401.8 ns/op	      31 B/op	       2 allocs/op
BenchmarkNativeSyncMapConcurrentOperations-4    	    1090	   1208427 ns/op	  220494 B/op	   10291 allocs/op
BenchmarkNativeSyncMapConcurrentStore-4         	     457	   3242486 ns/op	  175961 B/op	   10200 allocs/op
BenchmarkNativeSyncMapConcurrentSwap-4          	     422	   3307308 ns/op	  175948 B/op	   10200 allocs/op
BenchmarkNativeSyncMapConcurrentLoadOrStore-4   	    1617	    875344 ns/op	  170295 B/op	    8837 allocs/op
BenchmarkTypedSyncMapStoreAndDelete-4           	 6427506	       224.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedSyncMapRange-4                    	 9427014	       176.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedSyncMapLoad-4                     	 6493525	       222.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedSyncMapSimulateEntries-4          	 5228080	       381.1 ns/op	      96 B/op	       0 allocs/op
BenchmarkTypedSyncMapSimulateKeys-4             	 6117523	       230.7 ns/op	      41 B/op	       0 allocs/op
BenchmarkTypedSyncMapSimulateValues-4           	 6138771	       217.5 ns/op	      40 B/op	       0 allocs/op
BenchmarkTypedSyncMapSimulateUpdate-4           	 3143239	       381.8 ns/op	      31 B/op	       2 allocs/op
BenchmarkTypedSyncMapSimulateUpdateRange-4      	 3225171	       375.9 ns/op	      31 B/op	       2 allocs/op
BenchmarkTypedSyncMapConcurrentOperations-4     	    1105	   1206810 ns/op	  175872 B/op	   10199 allocs/op
BenchmarkTypedSyncMapConcurrentStore-4          	     469	   2896674 ns/op	  175955 B/op	   10200 allocs/op
BenchmarkTypedSyncMapConcurrentSwap-4           	     454	   3191634 ns/op	  175944 B/op	   10200 allocs/op
BenchmarkTypedSyncMapConcurrentLoadOrStore-4    	    1690	    943276 ns/op	  171243 B/op	    9783 allocs/op
BenchmarkTypedMapStoreAndDelete-4               	 8391085	       189.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedMapRange-4                        	73879828	        17.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedMapLoad-4                         	13276474	       166.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedMapEntries-4                      	60733093	        19.20 ns/op	      16 B/op	       0 allocs/op
BenchmarkTypedMapKeys-4                         	70526901	        17.24 ns/op	       8 B/op	       0 allocs/op
BenchmarkTypedMapValues-4                       	74341210	        17.10 ns/op	       8 B/op	       0 allocs/op
BenchmarkTypedMapUpdate-4                       	 7432243	       230.5 ns/op	      16 B/op	       1 allocs/op
BenchmarkTypedMapUpdateRange-4                  	42481010	        29.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedMapConcurrentOperations-4         	     549	   2329689 ns/op	   15956 B/op	     200 allocs/op
BenchmarkTypedMapConcurrentStore-4              	     756	   1670592 ns/op	   15960 B/op	     200 allocs/op
BenchmarkTypedMapConcurrentSwap-4               	     754	   1618879 ns/op	   15988 B/op	     200 allocs/op
BenchmarkTypedMapConcurrentLoadOrStore-4        	     738	   1591633 ns/op	   15972 B/op	     200 allocs/op
BenchmarkTypedMapConcurrentUpdate-4             	     769	   1600448 ns/op	   15977 B/op	     200 allocs/op
PASS
ok  	github.com/thetechpanda/typedmap/benchmarks	332.681s
```