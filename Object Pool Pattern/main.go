package main

import "fmt"

func main() {
	rpm := GetResourcePoolManager()

	for i := 0; i < 20; i++ {
		r, err := rpm.GetResource()
		if err != nil {
			fmt.Println("Unexpected error:", err)
			return
		}
		fmt.Println("Got:", r)
	}

	r, err := rpm.GetResource()
	if err != nil {
		fmt.Println("Expected exhaustion error:", err)
		return
	}

	fmt.Println("Got resource unexpectedly:", r)
}
