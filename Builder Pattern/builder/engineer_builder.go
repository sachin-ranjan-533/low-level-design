package builder

type EngineerBuilder struct {
	Name     string
	Age      int
	Roll_no  int
	Subjects []string
}

func NewEngineerBuilder() *EngineerBuilder {
	return &EngineerBuilder{}
}

func (eb *EngineerBuilder) SetName(name string) {
	eb.Name = name
}

func (eb *EngineerBuilder) SetAge(age int) {
	eb.Age = age
}

func (eb *EngineerBuilder) SetRollNo(roll_no int) {
	eb.Roll_no = roll_no
}

func (eb *EngineerBuilder) SetSubjects() {
	eb.Subjects = []string{"Maths", "Science", "Computer"}
}

func (eb *EngineerBuilder) Build() Builder {
	return eb
}

func (eb *EngineerBuilder) GetName() string {
	return eb.Name
}

func (eb *EngineerBuilder) GetAge() int {
	return eb.Age
}

func (eb *EngineerBuilder) GetRollNo() int {
	return eb.Roll_no
}

func (eb *EngineerBuilder) GetSubjects() []string {
	return eb.Subjects
}
