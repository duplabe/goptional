package goptional_test

import (
	"testing"

	"github.com/duplabe/goptional"
)

func TestNone(t *testing.T) {
	none := goptional.None[bool]()
	if !none.IsEmpty() {
		t.Fatalf(`None is not empty`)
	}

	if none.IsPresent() {
		t.Fatalf(`None is present`)
	}

	val, ok := none.Get()
	if ok {
		t.Fatalf(`Get is ok`)
	}

	if val {
		t.Fatalf(`Val is not false`)
	}

	if !none.GetOr(true) {
		t.Fatalf(`GetOr false`)
	}

	if none.GetOrZero() {
		t.Fatalf(`GetOrZero true`)
	}

	var flag bool
	none.IfPresent(func(_ bool) {
		flag = true
	})

	if flag {
		t.Fatalf(`IfPresent callback was called`)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	none.GetOrPanic()
}

func TestSome(t *testing.T) {
	some := goptional.Some(42)

	if some.IsEmpty() {
		t.Fatalf(`Some is empty`)
	}

	if !some.IsPresent() {
		t.Fatalf(`Some is not present`)
	}

	val, ok := some.Get()
	if !ok {
		t.Fatalf(`Get is not ok`)
	}

	if val == 0 {
		t.Fatalf(`Val is 0`)
	}

	if some.GetOr(100) != 42 {
		t.Fatalf(`GetOr is not 42`)
	}

	if some.GetOrZero() != 42 {
		t.Fatalf(`GetOrZero is not 42`)
	}

	var flag bool
	some.IfPresent(func(val int) {
		if val == 0 {
			t.Fatalf(`IfPresent val is 0`)
		}

		flag = true
	})

	if !flag {
		t.Fatalf(`IfPresent callback was not called`)
	}

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code paniced")
		}
	}()

	some.GetOrPanic()
}

func TestOf(t *testing.T) {
	val := 42
	var p *int

	if (goptional.Of(&val)).IsEmpty() {
		t.Fatalf(`empty`)
	}

	if (goptional.Of(p)).IsPresent() {
		t.Fatalf(`present`)
	}
}
