package gorbit

import (
	"fmt"
)

func HelloWorld(name string) error {
	fmt.Println("Starting Hello World Function")
	fmt.Println("Starting Hello World Function.")
	fmt.Println("Starting Hello World Function..")
	fmt.Println("Starting Hello World Function...")

	if len(name) == 0 {
		return fmt.Errorf("Error: No name or empty name given")
	}
	fmt.Printf("Hello %v!\n", name)
	return nil
}
