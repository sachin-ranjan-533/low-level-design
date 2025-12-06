package builder

type Builder interface {
	SetName(name string)
	SetAge(age int)
	SetRollNo(rollNo int)
	SetSubjects()
	Build() Builder
	GetName() string
	GetAge() int
	GetRollNo() int
	GetSubjects() []string
}
