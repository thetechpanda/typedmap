# Benchmarks
The benchmarks aim to compare the performance of `TypedMap` with that of `sync.Map`. Since `sync.Map` is not typed, the benchmarks focus on comparing the performance of operations that are common to both maps.

### Concurrent Benchmarks
Concurrent benchmark use all the same function body, so that each bench has the behaviour except for TypedMap and sync.Map operations.

Check `benchmarkConcurrentInt` in [benchmarks/benchmarks.go](benchmarks/benchmarks.go) for how the concurrent benchmark is implemented.

## Apple M2 Max 

```
BenchmarkNativeSyncMapStoreAndDelete-12           	 1396048	       790.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkNativeSyncMapRange-12                    	 2658678	       594.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkNativeSyncMapLoad-12                     	 1425985	       810.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkNativeSyncMapSimulateEntries-12          	 1000000	      1197 ns/op	      83 B/op	       0 allocs/op
BenchmarkNativeSyncMapSimulateKeys-12             	 2823246	       693.2 ns/op	      45 B/op	       0 allocs/op
BenchmarkNativeSyncMapSimulateValues-12           	 1882360	       642.8 ns/op	      43 B/op	       0 allocs/op
BenchmarkNativeSyncMapSimulateUpdate-12           	  831982	      1818 ns/op	      47 B/op	       2 allocs/op
BenchmarkNativeSyncMapSimulateUpdateRange-12      	 1000000	      1300 ns/op	      47 B/op	       2 allocs/op
BenchmarkNativeSyncMapConcurrentOperations-12     	      79	  16636238 ns/op	  231689 B/op	   10365 allocs/op
BenchmarkNativeSyncMapConcurrentStore-12          	      85	  23952053 ns/op	  182394 B/op	   10212 allocs/op
BenchmarkNativeSyncMapConcurrentSwap-12           	      78	  22730509 ns/op	  176194 B/op	   10204 allocs/op
BenchmarkNativeSyncMapConcurrentLoadOrStore-12    	     176	  10466779 ns/op	  122115 B/op	    5380 allocs/op
BenchmarkTypedSyncMapStoreAndDelete-12            	 1244832	      1017 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedSyncMapRange-12                     	 2125144	       788.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedSyncMapLoad-12                      	 1348383	       903.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedSyncMapSimulateEntries-12           	 1476582	      1097 ns/op	      88 B/op	       0 allocs/op
BenchmarkTypedSyncMapSimulateKeys-12              	 1000000	      1096 ns/op	      41 B/op	       0 allocs/op
BenchmarkTypedSyncMapSimulateValues-12            	 1000000	      1050 ns/op	      41 B/op	       0 allocs/op
BenchmarkTypedSyncMapSimulateUpdate-12            	  730126	      1992 ns/op	      47 B/op	       2 allocs/op
BenchmarkTypedSyncMapSimulateUpdateRange-12       	 1000000	      1770 ns/op	      47 B/op	       2 allocs/op
BenchmarkTypedSyncMapConcurrentOperations-12      	      75	  17230769 ns/op	  186672 B/op	   10280 allocs/op
BenchmarkTypedSyncMapConcurrentStore-12           	      79	  26353848 ns/op	  176047 B/op	   10202 allocs/op
BenchmarkTypedSyncMapConcurrentSwap-12            	      76	  26858961 ns/op	  176096 B/op	   10203 allocs/op
BenchmarkTypedSyncMapConcurrentLoadOrStore-12     	     184	   7612887 ns/op	   44179 B/op	    1771 allocs/op
BenchmarkTypedMapStoreAndDelete-12                	 1401651	       820.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedMapRange-12                         	10887088	       104.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedMapLoad-12                          	 1601234	       718.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedMapEntries-12                       	13415784	        93.26 ns/op	      16 B/op	       0 allocs/op
BenchmarkTypedMapKeys-12                          	20117828	        58.31 ns/op	       8 B/op	       0 allocs/op
BenchmarkTypedMapValues-12                        	14659636	        88.51 ns/op	       8 B/op	       0 allocs/op
BenchmarkTypedMapUpdate-12                        	 1261820	       853.9 ns/op	      16 B/op	       1 allocs/op
BenchmarkTypedMapUpdateRange-12                   	 7492224	       161.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedMapConcurrentOperations-12          	      42	  43040459 ns/op	   15413 B/op	     194 allocs/op
BenchmarkTypedMapConcurrentStore-12               	     100	  12164259 ns/op	   16082 B/op	     201 allocs/op
BenchmarkTypedMapConcurrentSwap-12                	     100	  12065482 ns/op	   15804 B/op	     198 allocs/op
BenchmarkTypedMapConcurrentLoadOrStore-12         	     112	  11389425 ns/op	   16060 B/op	     201 allocs/op
BenchmarkTypedMapConcurrentUpdate-12              	     100	  13105178 ns/op	   15884 B/op	     199 allocs/op
PASS
ok  	github.com/thetechpanda/typedmap/benchmarks	222.107s
```

## Intel(R) Xeon(R) CPU E5-2678 v3 @ 2.50GHz

```
goos: linux
goarch: amd64
pkg: github.com/thetechpanda/typedmap/benchmarks
cpu: Intel(R) Xeon(R) CPU E5-2678 v3 @ 2.50GHz
BenchmarkNativeSyncMapStoreAndDelete-4          	 1000000	      1230 ns/op	       0 B/op	       0 allocs/op
BenchmarkNativeSyncMapRange-4                   	 1231428	       957.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkNativeSyncMapLoad-4                    	 1000000	      1200 ns/op	       0 B/op	       0 allocs/op
BenchmarkNativeSyncMapSimulateEntries-4         	 1000000	      1216 ns/op	      83 B/op	       0 allocs/op
BenchmarkNativeSyncMapSimulateKeys-4            	 1000000	      1042 ns/op	      41 B/op	       0 allocs/op
BenchmarkNativeSyncMapSimulateValues-4          	 1000000	      1059 ns/op	      41 B/op	       0 allocs/op
BenchmarkNativeSyncMapSimulateUpdate-4          	  556580	      2680 ns/op	      47 B/op	       2 allocs/op
BenchmarkNativeSyncMapSimulateUpdateRange-4     	 1000000	      2044 ns/op	      47 B/op	       2 allocs/op
BenchmarkNativeSyncMapConcurrentOperations-4    	     126	   9543200 ns/op	  184693 B/op	   10214 allocs/op
BenchmarkNativeSyncMapConcurrentStore-4         	      63	  25804220 ns/op	  175738 B/op	   10199 allocs/op
BenchmarkNativeSyncMapConcurrentSwap-4          	      52	  26064633 ns/op	  176290 B/op	   10205 allocs/op
BenchmarkNativeSyncMapConcurrentLoadOrStore-4   	     184	   7771043 ns/op	  151012 B/op	    7739 allocs/op
BenchmarkTypedSyncMapStoreAndDelete-4           	 1000000	      1657 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedSyncMapRange-4                    	 1000000	      1043 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedSyncMapLoad-4                     	 1000000	      1368 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedSyncMapSimulateEntries-4          	 1000000	      1334 ns/op	      83 B/op	       0 allocs/op
BenchmarkTypedSyncMapSimulateKeys-4             	 1000000	      1252 ns/op	      41 B/op	       0 allocs/op
BenchmarkTypedSyncMapSimulateValues-4           	 1000000	      1265 ns/op	      41 B/op	       0 allocs/op
BenchmarkTypedSyncMapSimulateUpdate-4           	  469512	      2919 ns/op	      47 B/op	       2 allocs/op
BenchmarkTypedSyncMapSimulateUpdateRange-4      	  753475	      2373 ns/op	      47 B/op	       2 allocs/op
BenchmarkTypedSyncMapConcurrentOperations-4     	     133	   9076202 ns/op	  176005 B/op	   10199 allocs/op
BenchmarkTypedSyncMapConcurrentStore-4          	      57	  27671952 ns/op	  175992 B/op	   10202 allocs/op
BenchmarkTypedSyncMapConcurrentSwap-4           	      52	  27281447 ns/op	  175995 B/op	   10202 allocs/op
BenchmarkTypedSyncMapConcurrentLoadOrStore-4    	     195	   8536586 ns/op	  124643 B/op	    6731 allocs/op
BenchmarkTypedMapStoreAndDelete-4               	 1238575	       893.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedMapRange-4                        	11824228	       124.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedMapLoad-4                         	 1768748	       933.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedMapEntries-4                      	 8203728	       144.9 ns/op	      16 B/op	       0 allocs/op
BenchmarkTypedMapKeys-4                         	12758884	       102.8 ns/op	       8 B/op	       0 allocs/op
BenchmarkTypedMapValues-4                       	10166542	       119.0 ns/op	       8 B/op	       0 allocs/op
BenchmarkTypedMapUpdate-4                       	 1000000	      1104 ns/op	      16 B/op	       1 allocs/op
BenchmarkTypedMapUpdateRange-4                  	 7602775	       167.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkTypedMapConcurrentOperations-4         	      46	  48861308 ns/op	   15946 B/op	     198 allocs/op
BenchmarkTypedMapConcurrentStore-4              	     100	  12923046 ns/op	   15794 B/op	     198 allocs/op
BenchmarkTypedMapConcurrentSwap-4               	     100	  13534033 ns/op	   15892 B/op	     199 allocs/op
BenchmarkTypedMapConcurrentLoadOrStore-4        	      93	  12073894 ns/op	   15779 B/op	     198 allocs/op
BenchmarkTypedMapConcurrentUpdate-4             	     100	  12777332 ns/op	   15902 B/op	     199 allocs/op
PASS
ok  	github.com/thetechpanda/typedmap/benchmarks	199.901s
```