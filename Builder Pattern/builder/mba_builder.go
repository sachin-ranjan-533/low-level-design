package builder

type MbaBuilder struct {
	Name     string
	Age      int
	Roll_no  int
	Subjects []string
}

func NewMbaBuilder() *MbaBuilder {
	return &MbaBuilder{}
}

func (mb *MbaBuilder) SetName(name string) {
	mb.Name = name
}

func (mb *MbaBuilder) SetAge(age int) {
	mb.Age = age
}

func (mb *MbaBuilder) SetRollNo(roll_no int) {
	mb.Roll_no = roll_no
}

func (mb *MbaBuilder) SetSubjects() {
	mb.Subjects = []string{"Accounting", "Marketing", "Finance"}
}

func (mb *MbaBuilder) Build() Builder {
	return mb
}

func (mb *MbaBuilder) GetName() string {
	return mb.Name
}

func (mb *MbaBuilder) GetAge() int {
	return mb.Age
}

func (mb *MbaBuilder) GetRollNo() int {
	return mb.Roll_no
}

func (mb *MbaBuilder) GetSubjects() []string {
	return mb.Subjects
}
