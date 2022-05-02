package main

import (
	"fmt"
)

var (
	student_name             string = "Muhammad Abdulaleem"
	student_age              int    = 18
	student_hobby            string = "Programming 😍"
	student_github_user_name string = "aleemaheer"
)

var (
	num1 float32 = 23.3
	num2 float64 = 54.4
)

func main() {
	fmt.Printf("%v, %T, %v, %T", num1, num1, num2, num2)
	fmt.Println("\n     ------       Student Record      --------\nStudent Name: ", student_name, "\nStudent Age: ", student_age, "\nStudent Hobby: ", student_hobby, "\nStudent's Github: ", student_github_user_name)
}
