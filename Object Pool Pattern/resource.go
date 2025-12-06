package main

// Resource represents a pooled resource
// ------------------------------------------------------------
type Resource struct {
	id int
}

func NewResource(id int) Resource {
	// NOTE: In real system avoid printing inside constructor for performance
	println("New resource is created with id", id)
	return Resource{id: id}
}
