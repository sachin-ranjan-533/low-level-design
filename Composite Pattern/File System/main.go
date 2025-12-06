package main

import filesystem "composite-pattern/file_system"

func main() {
	mainDIrectory := filesystem.NewDirectory("root")
	binDirectory := filesystem.NewDirectory("bin")
	usrDirectory := filesystem.NewDirectory("usr")

	binDirectory.Add(filesystem.NewFile("bash"))
	binDirectory.Add(filesystem.NewFile("ls"))

	usrDirectory.Add(filesystem.NewFile("user1"))
	usrDirectory.Add(filesystem.NewFile("user2"))

	mainDIrectory.Add(binDirectory)
	mainDIrectory.Add(usrDirectory)

	mainDIrectory.GetName()
}
