package main

import (
	"builder-pattern/builder"
	"fmt"
)

func main() {
	director := NewDirector(builder.NewMbaBuilder())
	builder := director.Construct("John Doe", 25, 101)
	student := NewStudent(builder)
	fmt.Printf("%+v\n", student)
}
