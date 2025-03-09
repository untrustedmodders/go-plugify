package plugify

import (
	"fmt"
	"math"
)

type Vector2 struct {
	X float32
	Y float32
}

// Add another vector
func (v Vector2) Add(vector Vector2) Vector2 {
	return Vector2{v.X + vector.X, v.Y + vector.Y}
}

// Subtract another vector
func (v Vector2) Subtract(vector Vector2) Vector2 {
	return Vector2{v.X - vector.X, v.Y - vector.Y}
}

// Scale by a scalar
func (v Vector2) Scale(scalar float32) Vector2 {
	return Vector2{v.X * scalar, v.Y * scalar}
}

// Calculate the magnitude (length) of the vector
func (v Vector2) Magnitude() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

// Normalize the vector to a unit vector
func (v Vector2) Normalize() Vector2 {
	magnitude := v.Magnitude()
	if magnitude == 0 {
		return Vector2{0, 0}
	}
	return v.Scale(1 / magnitude)
}

// Dot product with another vector
func (v Vector2) Dot(vector Vector2) float32 {
	return v.X*vector.X + v.Y*vector.Y
}

// Calculate the distance to another vector
func (v Vector2) DistanceTo(vector Vector2) float32 {
	return float32(math.Sqrt(float64((v.X-vector.X)*(v.X-vector.X) + (v.Y-vector.Y)*(v.Y-vector.Y))))
}

// Return a string representation
func (v Vector2) ToString() string {
	return fmt.Sprintf("Vector2(%f, %f)", v.X, v.Y)
}

// Static method to create a zero vector
func Vector2Zero() Vector2 {
	return Vector2{0, 0}
}

// Static method to create a unit vector
func Vector2Unit() Vector2 {
	return Vector2{1, 1}
}
