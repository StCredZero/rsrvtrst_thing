package fib

import "sync"

const (
	InstantIter = 10000 // 10000 is a good default value, for a number of iterations runnable by a system, "instantly"
)

type Fibber struct {
	mutex  sync.Mutex
	cache  map[uint64]uint64
	maxN   uint64
	maxFib uint64
}

func NewFibber() *Fibber {
	f := new(Fibber)
	f.Initialize()
	return f
}

// Initialize synchronized to be reused to clear cache
func (f *Fibber) Initialize() {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	f.cache = make(map[uint64]uint64)
	f.cache[0] = 0
	f.cache[1] = 1
	f.maxN = 1
	f.maxFib = 1
}

// SyncGetCached synchronized method got getting cached result
func (f *Fibber) SyncGetCached(n uint64) (uint64, bool) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	fib, ok := f.cache[n]
	return fib, ok
}

// SyncGetCached synchronized method for setting cached result
func (f *Fibber) SyncSetCache(n uint64, fib uint64) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	f.cache[n] = fib
	if n > f.maxN {
		f.maxN = n
		f.maxFib = fib
	}
}

// bareExtend should only be used inside synchronized code
// It exists to eliminate duplicate code
func (f *Fibber) bareExtend(condition func(uint64) bool) {
	for i := f.maxN + 1; condition(i); i++ {
		fib := f.cache[i-1] + f.cache[i-2]
		f.cache[i] = fib
		f.maxN = i
		f.maxFib = fib
	}
}

// SyncExtendToN extends the cache to ordinal n
// This has the same complexity as the recursive algorithm, without using O(n) stack
func (f *Fibber) SyncExtendToN(n uint64) bool {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	limit := f.maxN + InstantIter
	if limit > n {
		limit = n
	}

	f.bareExtend(func(i uint64) bool {
		return i <= limit
	})

	return f.maxN == n
}

// SyncExtendToN extends the cache to a fibonaci value larger than x
// This has the same complexity as the recursive algorithm, without using O(n) stack
func (f *Fibber) SyncExtendToX(x uint64) bool {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	limit := f.maxN + InstantIter

	f.bareExtend(func(i uint64) bool {
		return i <= limit && f.maxFib < x
	})

	return f.maxFib >= x
}

// SyncIncludesX is required to query cache state without race conditions mixing values
func (f *Fibber) SyncIncludesX(x uint64) (uint64, uint64, bool) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return f.maxN, f.maxFib, f.maxFib >= x
}

// FibonaciOrdinal what it says on the tin
func (f *Fibber) FibonaciOrdinal(n uint64) uint64 {
	fibCached, foundCached := f.SyncGetCached(n)
	if foundCached {
		return fibCached
	}

	// extend the cache until we have calculated for ordinal n
	// we will do this InstantIter at a time, to prevent the Fibber from being locked for very long
	for !f.SyncExtendToN(n) {
	}

	fib, _ := f.SyncGetCached(n)
	return fib
}

// CardinalityLessThan answers sequence cardinality for items less than X
func (f *Fibber) CardinalityLessThan(x uint64) uint64 {

	// I just special case the starter ordinals, so I don't have to think too hard about the min boundary condition
	switch x {
	case 0:
		return 0
	case 1:
		return 1
	}

	maxN, maxFib, includes := f.SyncIncludesX(x)
	if includes {
		if maxFib == x {
			return maxN
		} else { // search for it
			return f.SearchForOrdinalLessThan(x, 0, maxN-1) + 1
		}
	}

	// extend the cache until we have calculated a sequence iteration greater than x
	// we will do this InstantIter at a time, to prevent the Fibber from being locked for very long
	for !f.SyncExtendToX(x) {
	}

	maxN, maxFib, _ = f.SyncIncludesX(x)
	return f.SearchForOrdinalLessThan(x, 0, maxN) + 1
}

// SearchForOrdinalLessThan implements binary search, with a few cache specific features
func (f *Fibber) SearchForOrdinalLessThan(x, min, max uint64) uint64 {

	halfway := (min + max) / 2
	fibHalf := f.FibonaciOrdinal(halfway)

	// cases x < fibHalf and x == fibHalf
	if fibHalf == x {
		return halfway - 1
	} else if x < fibHalf {
		return f.SearchForOrdinalLessThan(x, min, halfway)
	}

	// case x > fibhalf
	nextFib := f.FibonaciOrdinal(halfway + 1)
	if x <= nextFib {
		return halfway
	} else {
		return f.SearchForOrdinalLessThan(x, halfway+1, max)
	}
}
