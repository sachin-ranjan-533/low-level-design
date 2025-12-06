package filesystem

import "fmt"

type Directory struct {
	Name  string
	files []FileSystem
}

func NewDirectory(name string) *Directory {
	return &Directory{Name: name, files: []FileSystem{}}
}

func (d *Directory) Add(fs FileSystem) {
	d.files = append(d.files, fs)
}

func (d *Directory) GetName() {
	fmt.Println("Directory Name:", d.Name)
	for _, f := range d.files {
		f.GetName()
	}
}
