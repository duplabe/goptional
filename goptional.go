package goptional

type optional[T any] struct {
	val *T
}

func None[T any]() optional[T] {
	return optional[T]{}
}

func Some[T any](val T) optional[T] {
	return optional[T]{
		val: &val,
	}
}

func Of[T any](val *T) optional[T] {
	if val == nil {
		return None[T]()
	}

	return Some[T](*val)
}

func (o optional[T]) IsPresent() bool {
	return nil != o.val
}

func (o optional[T]) IsEmpty() bool {
	return !o.IsPresent()
}

func (o optional[T]) Get() (T, bool) {
	if o.IsPresent() {
		return *o.val, true
	}

	return *new(T), false
}

func (o optional[T]) GetOr(emptyValue T) T {
	if val, ok := o.Get(); ok {
		return val
	}

	return emptyValue
}

func (o optional[T]) GetOrZero() T {
	val, _ := o.Get()

	return val
}

func (o optional[T]) GetOrPanic() T {
	if val, ok := o.Get(); ok {
		return val
	}

	panic("not present")
}

func (o optional[T]) IfPresent(consumer func(T)) {
	if o.IsPresent() {
		consumer(*o.val)
	}
}
