package typed_roaring

import (
	"github.com/RoaringBitmap/roaring"
)

type Iterator[T BitConstraint] interface {
	HasNext() bool
	AdvanceIfNeeded(minVal T)
	Next() T
}

type typedRoaringBitmapIterator[T BitConstraint] struct {
	roaring.IntPeekable
}

func (t typedRoaringBitmapIterator[T]) Next() T {
	return T(t.IntPeekable.Next())
}

func (t typedRoaringBitmapIterator[T]) AdvanceIfNeeded(minVal T) {
	t.IntPeekable.AdvanceIfNeeded(uint32(minVal))
}
