package plugify

import (
	"fmt"
	"math"
)

type Matrix4x4 struct {
	M [4][4]float32
}

// Constructor to initialize the matrix
func NewMatrix4x4(elements [4][4]float32) Matrix4x4 {
	return Matrix4x4{M: elements}
}

// Add another matrix
func (m Matrix4x4) Add(matrix Matrix4x4) Matrix4x4 {
	var result [4][4]float32
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i][j] = m.M[i][j] + matrix.M[i][j]
		}
	}
	return Matrix4x4{M: result}
}

// Subtract another matrix
func (m Matrix4x4) Subtract(matrix Matrix4x4) Matrix4x4 {
	var result [4][4]float32
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i][j] = m.M[i][j] - matrix.M[i][j]
		}
	}
	return Matrix4x4{M: result}
}

// Multiply by another matrix
func (m Matrix4x4) Multiply(matrix Matrix4x4) Matrix4x4 {
	var result [4][4]float32
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				result[i][j] += m.M[i][k] * matrix.M[k][j]
			}
		}
	}
	return Matrix4x4{M: result}
}

// Multiply by a vector (Vector4)
func (m Matrix4x4) MultiplyVector(vector Vector4) Vector4 {
	return Vector4{
		X: m.M[0][0]*vector.X + m.M[0][1]*vector.Y + m.M[0][2]*vector.Z + m.M[0][3]*vector.W,
		Y: m.M[1][0]*vector.X + m.M[1][1]*vector.Y + m.M[1][2]*vector.Z + m.M[1][3]*vector.W,
		Z: m.M[2][0]*vector.X + m.M[2][1]*vector.Y + m.M[2][2]*vector.Z + m.M[2][3]*vector.W,
		W: m.M[3][0]*vector.X + m.M[3][1]*vector.Y + m.M[3][2]*vector.Z + m.M[3][3]*vector.W,
	}
}

// Transpose the matrix
func (m Matrix4x4) Transpose() Matrix4x4 {
	var result [4][4]float32
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[j][i] = m.M[i][j]
		}
	}
	return Matrix4x4{M: result}
}

// Get the identity matrix
func Matrix4x4Identity() Matrix4x4 {
	return Matrix4x4{
		M: [4][4]float32{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		},
	}
}

// Create a zero matrix
func Matrix4x4Zero() Matrix4x4 {
	return Matrix4x4{
		M: [4][4]float32{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
	}
}

// Create a scaling matrix
func Scaling(sx, sy, sz float32) Matrix4x4 {
	return Matrix4x4{
		M: [4][4]float32{
			{sx, 0, 0, 0},
			{0, sy, 0, 0},
			{0, 0, sz, 0},
			{0, 0, 0, 1},
		},
	}
}

// Create a translation matrix
func Translation(tx, ty, tz float32) Matrix4x4 {
	return Matrix4x4{
		M: [4][4]float32{
			{1, 0, 0, tx},
			{0, 1, 0, ty},
			{0, 0, 1, tz},
			{0, 0, 0, 1},
		},
	}
}

// Create a rotation matrix around the X-axis
func RotationX(angle float32) Matrix4x4 {
	c := float32(math.Cos(float64(angle)))
	s := float32(math.Sin(float64(angle)))
	return Matrix4x4{
		M: [4][4]float32{
			{1, 0, 0, 0},
			{0, c, -s, 0},
			{0, s, c, 0},
			{0, 0, 0, 1},
		},
	}
}

// Create a rotation matrix around the Y-axis
func RotationY(angle float32) Matrix4x4 {
	c := float32(math.Cos(float64(angle)))
	s := float32(math.Sin(float64(angle)))
	return Matrix4x4{
		M: [4][4]float32{
			{c, 0, s, 0},
			{0, 1, 0, 0},
			{-s, 0, c, 0},
			{0, 0, 0, 1},
		},
	}
}

// Create a rotation matrix around the Z-axis
func RotationZ(angle float32) Matrix4x4 {
	c := float32(math.Cos(float64(angle)))
	s := float32(math.Sin(float64(angle)))
	return Matrix4x4{
		M: [4][4]float32{
			{c, -s, 0, 0},
			{s, c, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		},
	}
}

// Print matrix as a formatted string
func (m Matrix4x4) ToString() string {
	return fmt.Sprintf(
		"Matrix4x4[%f, %f, %f, %f]\n[%f, %f, %f, %f]\n[%f, %f, %f, %f]\n[%f, %f, %f, %f]",
		m.M[0][0], m.M[0][1], m.M[0][2], m.M[0][3],
		m.M[1][0], m.M[1][1], m.M[1][2], m.M[1][3],
		m.M[2][0], m.M[2][1], m.M[2][2], m.M[2][3],
		m.M[3][0], m.M[3][1], m.M[3][2], m.M[3][3],
	)
}
