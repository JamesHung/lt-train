package main

import (
	"errors"
	"fmt"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	R float64
}

func (c *Circle) Area() float64 {
	if c == nil {
		fmt.Println("[Circle.Area] nil receiver")
		return 0
	}
	return 3.14 * c.R * c.R
}

type Rectangle struct {
	W, H float64
}

func (r *Rectangle) Area() float64 {
	if r == nil {
		fmt.Println("[Rectangle.Area] nil receiver")
		return 0
	}
	return r.W * r.H
}

// NewShape shows the correct factory pattern: return a true nil interface when invalid.
func NewShape(kind string) Shape {
	switch kind {
	case "circle":
		return &Circle{R: 10}
	case "rect":
		return &Rectangle{W: 5, H: 8}
	default:
		return nil // (type=nil, data=nil)
	}
}

// BadFactory demonstrates the wrapped-nil bug.
func BadFactory() Shape {
	var c *Circle = nil
	return c // interface = (type=*Circle, data=nil) => NOT nil
}

func demoInterfaceNilCore() {
	fmt.Println("== interface wrapping a nil concrete pointer ==")
	var c *Circle = nil
	var s Shape = c
	fmt.Printf("c == nil? %v\n", c == nil)                      // true
	fmt.Printf("s == nil? %v (holds type *Circle)\n", s == nil) // false
}

func demoFactoryPattern() {
	fmt.Println("\n== factory returns true nil on invalid kind ==")
	shape1 := NewShape("circle")
	shape2 := NewShape("unknown")

	if shape1 != nil {
		fmt.Printf("shape1 area = %.2f\n", shape1.Area()) // 314.00
	}
	if shape2 == nil {
		fmt.Println("shape2 is nil (invalid shape)")
	}
}

func demoWrappedNilDetection() {
	fmt.Println("\n== detecting wrapped nil from a bad factory ==")
	shape := BadFactory()
	fmt.Printf("shape == nil? %v\n", shape == nil) // false

	// Type assertion + nil check to detect wrapped nil.
	if ptr, ok := shape.(*Circle); ok && ptr == nil {
		fmt.Println("detected: interface holds a nil *Circle (wrapped nil)")
		return
	}
	fmt.Println("non-nil shape, area =", shape.Area())
}

type User struct {
	ID   string
	Name string
}

type UserRepository interface {
	FindUser(id string) (*User, error)
}

var ErrUserNotFound = errors.New("user not found")

type InMemoryUserRepo struct {
	data map[string]*User
}

func (r *InMemoryUserRepo) FindUser(id string) (*User, error) {
	u, ok := r.data[id]
	if !ok {
		return nil, ErrUserNotFound
	}
	return u, nil
}

func demoRepositoryPattern() {
	fmt.Println("\n== repository prefers (*T, error) over interface ==")
	repo := &InMemoryUserRepo{
		data: map[string]*User{
			"123": {ID: "123", Name: "Alice"},
		},
	}

	u1, err := repo.FindUser("123")
	fmt.Printf("u1: user=%+v err=%v\n", u1, err) // user=&{ID:123 Name:Alice} err=<nil>

	u2, err := repo.FindUser("999")
	fmt.Printf("u2: user=%+v err=%v (errors.Is? %v)\n", u2, err, errors.Is(err, ErrUserNotFound)) // user=<nil> err=user not found (errors.Is? true)
}

func main() {
	demoInterfaceNilCore()
	demoFactoryPattern()
	demoWrappedNilDetection()
	demoRepositoryPattern()
}
