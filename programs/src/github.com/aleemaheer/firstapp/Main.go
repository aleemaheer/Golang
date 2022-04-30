package main

import (
	"fmt"
)

// 46 min
// Variables that are declared but not used, compiler will give error, as our application grows and updates the previous variables that are declared must be removed
// If we want to declare more than variables at the global scope, we can use the following approach, this will look like more convenint
var (
	actor_name   string = "Steve Jobs"
	actor_age    int    = 45
	actor_height string = "4.2 feet"
)

var (
	i int = 100
)

func main() {
	fmt.Println(i)       // If we have a variable out of the main function in the package level, this is global, but we can reasign it with another data, this is called shadowing
	var i float32 = 10.5 // We can declare variable and assign value by defining it's data type
	z := 20.5            // We can also declare variable like this, go will figure out the data type, note that if we make variable of float by using this approach, the variable data type will be float64, if we want a variable of data type float32 we have to use the upper syntax.
	fmt.Println("Hello world", i, z)
	fmt.Printf("%v, %T\n", z, z)
	fmt.Println("Actor Name: ", actor_name, "\nActor age: ", actor_age, "\nActor heging: ", actor_height)
}
