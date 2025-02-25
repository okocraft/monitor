package null

type Optional[T any] struct {
	value *T
}

func (o Optional[T]) Get() T {
	if o.value != nil {
		return *o.value
	}
	var zero T
	return zero
}

func (o Optional[T]) GetOrElse(def T) T {
	if o.value != nil {
		return *o.value
	}
	return def
}

func (o Optional[T]) Valid() bool {
	return o.value != nil
}

func FromValue[T any](v T) Optional[T] {
	return Optional[T]{value: &v}
}

func Empty[T any]() Optional[T] {
	return Optional[T]{}
}

func FromPtr[T any](v *T) Optional[T] {
	if v == nil {
		return Empty[T]()
	}
	return FromValue(*v)
}
