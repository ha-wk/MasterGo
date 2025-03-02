package main

import (
	"fmt"
	"math"
)

// Shape struct (Reusable Base Struct)
type Shape struct {
	color string
}

// Circle struct
type Circle struct {
	Shape  // Embedding Shape struct
	radius float64
}

// Rectangle struct
type Rectangle struct {
	Shape   // Embedding Shape struct
	length  float64
	breadth float64
}

// Metrics interface defines common behavior
type Metrics interface {
	Area() float64
	Perimeter() float64
	GetColor() string
}

// Implement Area and Perimeter for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Implement Area and Perimeter for Rectangle
func (r Rectangle) Area() float64 {
	return r.length * r.breadth
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

// Implement GetColor method for both
func (s Shape) GetColor() string {
	return s.color
}

// Factory functions for better initialization
func NewCircle(radius float64, color string) Circle {
	return Circle{Shape: Shape{color: color}, radius: radius}
}

func NewRectangle(length, breadth float64, color string) Rectangle {
	return Rectangle{Shape: Shape{color: color}, length: length, breadth: breadth}
}

// Generic function to print metrics
func printMetrics(m Metrics) {
	fmt.Printf("Color: %s | Area: %.2f | Perimeter: %.2f\n", m.GetColor(), m.Area(), m.Perimeter())
}

func main() {
	// Using factory functions to create instances
	circle := NewCircle(3.0, "Red")
	rectangle := NewRectangle(4.0, 2.0, "Blue")

	// Print metrics for both
	fmt.Println("Circle Metrics:")
	printMetrics(circle)

	fmt.Println("Rectangle Metrics:")
	printMetrics(rectangle)
}
