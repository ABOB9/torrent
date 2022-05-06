package typed_roaring

type BitConstraint interface {
	~int | ~uint32
}
