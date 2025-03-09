package plugify

import (
	"fmt"
	"math"
)

type Vector4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

// Add another vector
func (v Vector4) Add(vector Vector4) Vector4 {
	return Vector4{v.X + vector.X, v.Y + vector.Y, v.Z + vector.Z, v.W + vector.W}
}

// Subtract another vector
func (v Vector4) Subtract(vector Vector4) Vector4 {
	return Vector4{v.X - vector.X, v.Y - vector.Y, v.Z - vector.Z, v.W - vector.W}
}

// Scale by a scalar
func (v Vector4) Scale(scalar float32) Vector4 {
	return Vector4{v.X * scalar, v.Y * scalar, v.Z * scalar, v.W * scalar}
}

// Calculate the magnitude (length) of the vector
func (v Vector4) Magnitude() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)))
}

// Normalize the vector to a unit vector
func (v Vector4) Normalize() Vector4 {
	magnitude := v.Magnitude()
	if magnitude == 0 {
		return Vector4{0, 0, 0, 0}
	}
	return v.Scale(1 / magnitude)
}

// Dot product with another vector
func (v Vector4) Dot(vector Vector4) float32 {
	return v.X*vector.X + v.Y*vector.Y + v.Z*vector.Z + v.W*vector.W
}

// Calculate the distance to another vector
func (v Vector4) DistanceTo(vector Vector4) float32 {
	return float32(math.Sqrt(float64((v.X-vector.X)*(v.X-vector.X) + (v.Y-vector.Y)*(v.Y-vector.Y) + (v.Z-vector.Z)*(v.Z-vector.Z) + (v.W-vector.W)*(v.W-vector.W))))
}

// Return a string representation
func (v Vector4) ToString() string {
	return fmt.Sprintf("Vector4(%f, %f, %f, %f)", v.X, v.Y, v.Z, v.W)
}

// Static method to create a zero vector
func Vector4Zero() Vector4 {
	return Vector4{0, 0, 0, 0}
}

// Static method to create a unit vector
func Vector4Unit() Vector4 {
	return Vector4{1, 1, 1, 1}
}
