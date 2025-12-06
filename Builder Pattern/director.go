package main

import "builder-pattern/builder"

type Director struct {
	builder builder.Builder
}

func NewDirector(b builder.Builder) *Director {
	return &Director{builder: b}
}

func (d *Director) Construct(name string, age int, roll_no int) builder.Builder {
	d.builder.SetName(name)
	d.builder.SetAge(age)
	d.builder.SetRollNo(roll_no)
	d.builder.SetSubjects()
	return d.builder.Build()
}
