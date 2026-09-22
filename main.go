package main

import "fmt"

// ========================================
// Domain Layer
// ========================================

// Operation adalah abstraction untuk sebuah operasi matematika.
type Operation interface {
	Apply(a, b float64) float64
}

// Addition adalah implementasi operasi penjumlahan.
type Addition struct{}

func (Addition) Apply(a, b float64) float64 {
	return a + b
}

// Subtraction adalah implementasi operasi pengurangan.
type Subtraction struct{}

func (Subtraction) Apply(a, b float64) float64 {
	return a - b
}

// Multiplication adalah implementasi operasi perkalian.
type Multiplication struct{}

func (Multiplication) Apply(a, b float64) float64 {
	return a * b
}

// Division adalah implementasi operasi pembagian.
type Division struct{}

func (Division) Apply(a, b float64) float64 {
	return a / b
}

// ========================================
// Application / Service Layer
// ========================================

// Calculator menggunakan Operation sebagai dependency.
type Calculator struct {
	operation Operation
}

// NewCalculator adalah constructor untuk Calculator.
//
// Dependency Injection dilakukan melalui parameter operation.
func NewCalculator(operation Operation) *Calculator {
	return &Calculator{
		operation: operation,
	}
}

// Calculate menjalankan operasi yang sudah diberikan.
func (c *Calculator) Calculate(a, b float64) float64 {
	return c.operation.Apply(a, b)
}

// ========================================
// Main
// ========================================

func main() {
	addition := Addition{}

	calculator := NewCalculator(addition)

	result := calculator.Calculate(10, 5)

	fmt.Println("10 + 5 =", result)
}
