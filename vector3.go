package plugify

import (
	"fmt"
	"math"
)

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

// Add another vector
func (v Vector3) Add(vector Vector3) Vector3 {
	return Vector3{v.X + vector.X, v.Y + vector.Y, v.Z + vector.Z}
}

// Subtract another vector
func (v Vector3) Subtract(vector Vector3) Vector3 {
	return Vector3{v.X - vector.X, v.Y - vector.Y, v.Z - vector.Z}
}

// Scale by a scalar
func (v Vector3) Scale(scalar float32) Vector3 {
	return Vector3{v.X * scalar, v.Y * scalar, v.Z * scalar}
}

// Calculate the magnitude (length) of the vector
func (v Vector3) Magnitude() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
}

// Normalize the vector to a unit vector
func (v Vector3) Normalize() Vector3 {
	magnitude := v.Magnitude()
	if magnitude == 0 {
		return Vector3{0, 0, 0}
	}
	return v.Scale(1 / magnitude)
}

// Dot product with another vector
func (v Vector3) Dot(vector Vector3) float32 {
	return v.X*vector.X + v.Y*vector.Y + v.Z*vector.Z
}

// Cross product with another vector
func (v Vector3) Cross(vector Vector3) Vector3 {
	return Vector3{
		v.Y*vector.Z - v.Z*vector.Y,
		v.Z*vector.X - v.X*vector.Z,
		v.X*vector.Y - v.Y*vector.X,
	}
}

// Calculate the distance to another vector
func (v Vector3) DistanceTo(vector Vector3) float32 {
	return float32(math.Sqrt(float64((v.X-vector.X)*(v.X-vector.X) + (v.Y-vector.Y)*(v.Y-vector.Y) + (v.Z-vector.Z)*(v.Z-vector.Z))))
}

// Return a string representation
func (v Vector3) ToString() string {
	return fmt.Sprintf("Vector3(%f, %f, %f)", v.X, v.Y, v.Z)
}

// Static method to create a zero vector
func Vector3Zero() Vector3 {
	return Vector3{0, 0, 0}
}

// Static method to create a unit vector
func Vector3Unit() Vector3 {
	return Vector3{1, 1, 1}
}
