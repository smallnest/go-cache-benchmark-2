package main

import (
	"flag"
	"runtime"
	"sync"
	"time"
)

type Benchmark struct {
	Generator
	N int
}

var (
	concurrency = flag.Int("c", 4, "Number of concurrent workers")
	cacheSize   = flag.Int("s", 100e3, "Cache size")
	multiplier  = flag.Int("m", 100, "multiplier for number of keys")
	genType     = flag.String("g", "zipfian", "zipfian, hotspot, uniform")
)

var genTypes = make(map[string]NewGeneratorFunc)

func init() {
	genTypes["zipfian"] = NewScrambledZipfian
	genTypes["hotspot"] = NewHotspot
	genTypes["uniform"] = NewUniform
}

func main() {
	flag.Parse()

	// cacheSize := []int{1e3, 10e3, 100e3, 1e6}
	// multiplier := []int{10, 100, 1000}

	cacheSizes := []int{*cacheSize}
	multipliers := []int{*multiplier}

	newCache := []NewCacheFunc{
		NewTinyLFU,
		NewClockPro,
		NewARC,
		NewRistretto,
		NewDirectCache,
		NewTwoQueue,
		NewGroupCacheLRU,
		NewHashicorpLRU,
		NewS4LRU,
		NewSLRU,
		NewWTFCache,
		NewFreeCache,
		NewBigCache,
		NewFastCache,
		NewSyncMap,
		NewMutexMap,
		NewKodingCache,
		NewGcache,
	}

	var results []*BenchmarkResult

	newGen := genTypes[*genType]
	if newGen == nil {
		newGen = NewScrambledZipfian
	}

	for _, cacheSize := range cacheSizes {
		for _, multiplier := range multipliers {
			numKey := cacheSize * multiplier

			for _, newCache := range newCache {
				result := run(newGen, cacheSize, numKey, newCache)
				results = append(results, result)
			}

			if len(results) > 0 {
				printResults(results)
				results = results[:0]
			}
		}
	}

}

func run(newGen NewGeneratorFunc, cacheSize, numKey int, newCache NewCacheFunc) *BenchmarkResult {
	gen := newGen(numKey)
	b := &Benchmark{
		Generator: gen,
		N:         1e6,
	}

	alloc1 := memAlloc()
	cache := newCache(cacheSize)
	defer cache.Close()

	start := time.Now()
	hits, misses := 0, 0
	if *concurrency <= 1 {
		hits, misses = bench(b, cache)
	} else {
		hits, misses = bench_par(b, cache, 4)
	}

	dur := time.Since(start)

	alloc2 := memAlloc()

	return &BenchmarkResult{
		GenName:     gen.Name(),
		CacheName:   cache.Name(),
		CacheSize:   cacheSize,
		NumKey:      numKey,
		Concurrency: *concurrency,

		Hits:     hits,
		Misses:   misses,
		Duration: dur,
		Bytes:    int64(alloc2) - int64(alloc1),
	}
}
func bench_par(b *Benchmark, cache Cache, concurrency int) (hits, misses int) {
	var wg sync.WaitGroup
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()

			n := b.N / concurrency
			for i := 0; i < n; i++ {
				value := b.Next()
				if cache.Get(value) {
					hits++
				} else {
					misses++
					cache.Set(value)
				}
			}
		}()
	}

	wg.Wait()

	return hits, misses
}

func bench(b *Benchmark, cache Cache) (hits, misses int) {
	for i := 0; i < b.N; i++ {
		value := b.Next()
		if cache.Get(value) {
			hits++
		} else {
			misses++
			cache.Set(value)
		}
	}

	return hits, misses
}

func memAlloc() uint64 {
	runtime.GC()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc
}
