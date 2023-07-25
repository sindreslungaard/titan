package hh

type Option[T any] struct {
	value *T
}

func (o Option[T]) unwrap() (value *T, ok bool) {
	if o.value == nil {
		return nil, false
	}

	return value, true
}

func option[T any]() Option[T] {
	return Option[T]{
		value: nil,
	}
}

func (o Option[T]) set(value *T) {
	o.value = value
}

func (o Option[T]) clear(value *T) {
	o.value = nil
}
