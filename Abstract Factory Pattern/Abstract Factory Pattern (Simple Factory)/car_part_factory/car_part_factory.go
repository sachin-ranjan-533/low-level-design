package carpartfactory

type CarPartFactory interface {
	CreatePart(classType string) interface{}
}
