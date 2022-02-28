package vector_test

import (
	"testing"

	"github.com/derekmwright/aegis/vector"
)

func TestNew(t *testing.T) {
	v := new(vector.Vector2)
	if v.X != 0.0 || v.Y != 0.0 {
		t.Fail()
	}
}

func TestSet(t *testing.T) {
	v := new(vector.Vector2)
	v.Set(0.5, 0.2)
	if v.X != 0.5 {
		t.Fail()
	}
	if v.Y != 0.2 {
		t.Fail()
	}
}

func TestSetScalar(t *testing.T) {
	v := new(vector.Vector2)
	v.SetScalar(0.4)
	if v.X != 0.4 {
		t.Fail()
	}
	if v.Y != 0.4 {
		t.Fail()
	}
}

func TestSetX(t *testing.T) {
	v := new(vector.Vector2)
	v.SetX(0.1)
	if v.X != 0.1 {
		t.Fail()
	}
	if v.Y != 0.0 {
		t.Fail()
	}
}

func TestSetY(t *testing.T) {
	v := new(vector.Vector2)
	v.SetY(0.2)
	if v.X != 0.0 {
		t.Fail()
	}
	if v.Y != 0.2 {
		t.Fail()
	}
}

func TestClone(t *testing.T) {
	v := new(vector.Vector2)
	v.SetX(0.2)
	v.SetY(2.2)
	v2 := v.Clone()
	if v2.X != 0.2 {
		t.Fail()
	}
	if v2.Y != 2.2 {
		t.Fail()
	}
}

func TestCopy(t *testing.T) {
	v := new(vector.Vector2)
	v.SetX(0.2)
	v.SetY(2.2)
	v2 := new(vector.Vector2)
	v.Copy(v2)
	if v.X != 0.0 {
		t.Fail()
	}
	if v.Y != 0.0 {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	v := new(vector.Vector2)
	v.SetX(0.2)
	v.SetY(2.2)
	v2 := new(vector.Vector2)
	v2.SetScalar(1.0)
	v.Add(v2)
	if v.X != 1.2 {
		t.Fail()
	}
	if v.Y != 3.2 {
		t.Fail()
	}
}

func TestAddScalar(t *testing.T) {
	v := new(vector.Vector2)
	v.SetX(0.2)
	v.SetY(2.2)
	v.AddScalar(1.0)
	if v.X != 1.2 {
		t.Fail()
	}
	if v.Y != 3.2 {
		t.Fail()
	}
}

func TestFloor(t *testing.T) {
	v := new(vector.Vector2)
	v.SetX(0.9)
	v.SetY(2.2)
	v.Floor()
	if v.X != 0.0 {
		t.Fail()
	}
	if v.Y != 2.0 {
		t.Fail()
	}
}

func TestCeil(t *testing.T) {
	v := new(vector.Vector2)
	v.SetX(0.9)
	v.SetY(2.2)
	v.Ceil()
	if v.X != 1.0 {
		t.Fail()
	}
	if v.Y != 3.0 {
		t.Fail()
	}
}

func TestRound(t *testing.T) {
	v := new(vector.Vector2)
	v.SetX(0.9)
	v.SetY(2.2)
	v.Round()
	if v.X != 1.0 {
		t.Fail()
	}
	if v.Y != 2.0 {
		t.Fail()
	}
}

func TestAngle(t *testing.T) {
	v := new(vector.Vector2)
	v.SetX(1.0)
	v.SetY(1.0)
	if v.Angle() != 0.78539816339744827899949086713604629039764404296875 {
		t.Fail()
	}
}
