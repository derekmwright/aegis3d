package vector

import "math"

// Vector2 is a two dimensional representation of vectors/points in space
type Vector2 struct {
	X float64
	Y float64
}

// Width returns the value of X
func (v *Vector2) Width() (w float64) {
	w = v.X
	return
}

// Height returns the value of Y
func (v *Vector2) Height() (h float64) {
	h = v.Y
	return
}

// Set sets X,Y to quantity passed via arguments
func (v *Vector2) Set(x float64, y float64) {
	v.X = x
	v.Y = y
	return
}

// SetScalar sets X,Y to scalar quantity
func (v *Vector2) SetScalar(s float64) {
	v.X = s
	v.Y = s
	return
}

// SetX sets the value of X
func (v *Vector2) SetX(x float64) {
	v.X = x
	return
}

// SetY sets the value of Y
func (v *Vector2) SetY(y float64) {
	v.Y = y
	return
}

// Clone clones the instance Vector2 to a new instance of a Vector2
func (v *Vector2) Clone() (c Vector2) {
	c.X = v.X
	c.Y = v.Y
	return
}

// Copy copies the X/Y quantities from the Vector2 passed via argument
func (v *Vector2) Copy(c *Vector2) {
	v.X = c.X
	v.Y = c.Y
	return
}

// Add adds the passed Vector quantities to the instance Vector2
func (v *Vector2) Add(c *Vector2) {
	v.X += c.X
	v.Y += c.Y
	return
}

// AddScalar adds a scalar quantity to the instance Vector2
func (v *Vector2) AddScalar(s float64) {
	v.X += s
	v.Y += s
	return
}

// Floor floors both X and Y values of the Vector2
func (v *Vector2) Floor() {
	v.X = math.Floor(v.X)
	v.Y = math.Floor(v.Y)
	return
}

// Ceil ceils both X and Y values of the Vector2
func (v *Vector2) Ceil() {
	v.X = math.Ceil(v.X)
	v.Y = math.Ceil(v.Y)
	return
}

// Round rounds both X and Y values of the Vector2
func (v *Vector2) Round() {
	v.X = math.Floor(v.X + .5)
	v.Y = math.Floor(v.Y + .5)
	return
}

// Angle computes the angle in radians with respect to the positive x-axis
func (v *Vector2) Angle() (a float64) {
	a = math.Atan2(v.Y, v.X)
	if a < 0 {
		a += 2 * math.Pi
	}
	return
}
