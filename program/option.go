package program

type Option[T any] struct {
	value *T
}

func (o Option[T]) Unwrap() (value *T, ok bool) {
	if o.value == nil {
		return nil, false
	}

	return value, true
}

func NewOption[T any](value *T) Option[T] {
	return Option[T]{
		value: value,
	}
}

func EmptyOption[T any]() Option[T] {
	return Option[T]{
		value: nil,
	}
}

func (o Option[T]) Set(value *T) {
	o.value = value
}
