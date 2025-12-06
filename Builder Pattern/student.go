package main

import "builder-pattern/builder"

type Student struct {
	Name     string
	Age      int
	RollNo   int
	Subjects []string
}

func NewStudent(b builder.Builder) *Student {
	return &Student{
		Name:     b.GetName(),
		Age:      b.GetAge(),
		RollNo:   b.GetRollNo(),
		Subjects: b.GetSubjects(),
	}
}
