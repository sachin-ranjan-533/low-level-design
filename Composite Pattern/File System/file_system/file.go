package filesystem

import "fmt"

type File struct {
	Name string
}

func NewFile(name string) *File {
	return &File{Name: name}
}

func (f *File) GetName() {
	fmt.Println("File Name:", f.Name)
}
