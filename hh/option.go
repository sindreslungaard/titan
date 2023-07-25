package hh

type Option[T any] struct {
	value *T
}

func option[T any]() *Option[T] {
	return &Option[T]{
		value: nil,
	}
}

func (o *Option[T]) unwrap() (value *T, ok bool) {
	if o.value == nil {
		return nil, false
	}

	return o.value, true
}

func (o *Option[T]) set(value *T) *Option[T] {
	o.value = value
	return o
}

func (o *Option[T]) clear() {
	o.value = nil
}

func (o *Option[T]) some() bool {
	return o.value != nil
}
