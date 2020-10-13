package vector

import (
	"fmt"
	"math"
)

// Vec is a two dimensional vector
type Vec struct {
	X, Y float64
}

// New returns a new vector at (x, y)
func New(x, y float64) *Vec {
	return &Vec{x, y}
}

// Zero returns a zero-vector
func Zero() *Vec {
	return &Vec{0, 0}
}

// Unit returns a unit vector at the given angle
func Unit(angle float64) *Vec {
	v := New(1, 0)
	v.Rotate(angle)
	return v
}

// Magnitude returns the magnitude of a vector
func (v *Vec) Magnitude() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// Angle returns the angle of a vector between [-PI, PI]
func (v Vec) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}

// Add adds vector u to vector v
func (v *Vec) Add(u *Vec) {
	v.X += u.X
	v.Y += u.Y
}

// Cap limits the length of a vector. If it is longer than max, it will be
// cut down to length max
func (v *Vec) Cap(max float64) {
	if max < v.Magnitude() {
		v.SetMagnitude(max)
	}
}

// Cross calculates the cross product of two vectors and returns a scalar
func (v Vec) Cross(u *Vec) float64 {
	return v.X*u.Y - u.X*v.Y
}

// Divide divides the vector by div
func (v *Vec) Divide(div float64) {
	v.X /= div
	v.Y /= div
}

// Dot calculates the dot product of two vectors and returns a scalar
func (v Vec) Dot(u *Vec) float64 {
	return v.X*u.X + v.Y*u.Y
}

// Multiply scales a vector by mul
func (v *Vec) Multiply(mul float64) {
	v.X *= mul
	v.Y *= mul
}

// Normalize normalizes the vector
func (v *Vec) Normalize() {
	v.Divide(v.Magnitude())
}

// Project projects vector u onto vector v
func (v *Vec) Project(u *Vec) {
	l := v.Dot(u) / u.Magnitude()
	*v = *u
	v.Normalize()
	v.Multiply(l)
}

// Rotate rotates the vector by angle in radians
func (v *Vec) Rotate(angle float64) {
	sin, cos := math.Sincos(angle)
	v.X = v.X*cos - v.Y*sin
	v.Y = v.X*sin + v.Y*cos
}

// SetMagnitude sets the magnitude of the vector to l
func (v *Vec) SetMagnitude(l float64) {
	v.Normalize()
	v.Multiply(l)
}

// String renders the vector as a string
func (v *Vec) String() string {
	return fmt.Sprintf("Vector(%f, %f)", v.X, v.Y)
}

// Subtract subtracts vector u from vector v
func (v *Vec) Subtract(u *Vec) {
	v.X -= u.X
	v.Y -= u.Y
}