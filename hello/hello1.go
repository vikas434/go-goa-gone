package hello

import (
	"examples.com/greetings"
	"fmt"
)

func main() {
	// Get a greetings message and print it.
	message := greetings.Hello("VIkas")
	fmt.Println(message)
}
