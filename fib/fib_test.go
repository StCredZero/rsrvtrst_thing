package fib

import (
	"reflect"
	"testing"
)

func TestFib(t *testing.T) {
	fibber := NewFibber(nil)
	fibber.testInit()
	fibber.SyncExtendToN(20)
	expected := []uint64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765}
	result := make([]uint64, 21)
	for k, v := range fibber.cache {
		result[k] = v
	}
	if !reflect.DeepEqual(expected, result) {
		t.Fatal("did not see expected sequence!")
	}
}

func TestFibOrdinal(t *testing.T) {
	fibber := NewFibber(nil)
	fibber.testInit()
	value := fibber.FibonaciOrdinal(9)
	if value != 34 {
		t.Fatal("did not see expected result for ordinal!")
	}
	value2 := fibber.FibonaciOrdinal(11)
	if value2 != 89 {
		t.Fatal("did not see expected result for ordinal!")
	}
}

func TestOrdinalSearch1(t *testing.T) {
	fibber := NewFibber(nil)
	fibber.testInit()
	fibber.SyncExtendToN(9)

	result := fibber.SearchForOrdinalLessThan(12, 0, 9)
	if result != 6 {
		t.Fatal("did not see expected result for ordinal search!")
	}
}

func TestOrdinalSearch2(t *testing.T) {
	fibber := NewFibber(nil)
	fibber.testInit()
	fibber.SyncExtendToN(9)

	result := fibber.SearchForOrdinalLessThan(34, 0, 9)
	if result != 8 {
		t.Fatal("did not see expected result for ordinal search!")
	}
}

func TestCardinality1(t *testing.T) {
	fibber := NewFibber(nil)
	fibber.testInit()
	fibber.SyncExtendToN(9)

	result := fibber.CardinalityLessThan(120)
	if result != 12 {
		t.Fatal("did not see expected result for cardinality1!")
	}
}

func itemsLessThan(x uint64, list []uint64) uint64 {
	count := uint64(0)
	for _, v := range list {
		if v < x {
			count += 1
		}
	}
	return count
}

func TestCardinality2(t *testing.T) {
	fibber := NewFibber(nil)
	fibber.testInit()
	fibber.SyncExtendToN(9)

	expected := []uint64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

	for i := uint64(0); i < 14; i++ {
		expectedCount := itemsLessThan(i, expected)
		result := fibber.CardinalityLessThan(i)
		if result != expectedCount {
			t.Fatal("did not see expected result for cardinality!")
		}
	}
}
