package vector

import (
	"math"
)

// Add adds vector u to vector v and returns the vector
func Add(v, u *Vec) *Vec {
	return &Vec{
		v.X + u.X,
		v.Y + u.Y,
	}
}

// Divide divides the vector by div and returns the vector
func Divide(v *Vec, div float64) *Vec {
	return &Vec{
		v.X / div,
		v.Y / div,
	}
}

// Multiply scales a vector by mul and returns the vector
func Multiply(v *Vec, mul float64) *Vec {
	return &Vec{
		v.X / mul,
		v.Y / mul,
	}
}

// Normalize normalizes the vector and returns the vector
func Normalize(v *Vec) *Vec {
	return Divide(v, v.Magnitude())
}

// Project projects vector u onto vector v and returns the vector
func Project(v, u *Vec) *Vec {
	l := v.Dot(u) / u.Magnitude()
	result := Normalize(u)
	result.Multiply(l)
	return result
}

// Rotate rotates the vector by angle in radians and returns the vector
func Rotate(v *Vec, angle float64) *Vec {
	sin, cos := math.Sincos(angle)
	return &Vec{
		v.X*cos - v.Y*sin,
		v.X*sin + v.Y*cos,
	}
}

// Subtract subtracts vector u from vector v and returns the vector
func Subtract(v, u *Vec) *Vec {
	return &Vec{
		v.X - u.X,
		v.Y - u.Y,
	}
}
