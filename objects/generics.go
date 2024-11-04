package objects

import (
	"fmt"
)

// wrap / type alias to implement the summary method
type Str string
type I int
type Bl bool
type Fl32 float32
type Fl64 float64

type Container[T any] struct {
	Val T
}

type Summarise interface {
	Summary() string
}

func (con Container[T]) Summary() string {
	// any() turns con.val to type: any.
	// .(summarise) checks if con.val impl's summarise - if it does then value = con.Val
	if value, ok := any(con.Val).(Summarise); ok {
		// ^ in essence if ok == true (ok is bool ret from type assertion above)
		return value.Summary()
	}
	return "No Summary method"
}

func (s Str) Summary() string {
	return fmt.Sprintf("String input: %s\n", s)
}

func (integer I) Summary() string {
	return fmt.Sprintf("Integer input: %d\n", integer)
}

func (boolean Bl) Summary() string {
	return fmt.Sprintf("Boolean input: %t\n", boolean)
}

func (float32 Fl32) Summary() string {
	return fmt.Sprintf("32 bit floating point input: %f\n", float32)
}

func (float64 Fl64) Summary() string {
	return fmt.Sprintf("64 bit floating point input: %f\n", float64)
}
