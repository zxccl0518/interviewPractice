package main

import "fmt"

type people interface {
	Speak(string) string
}

type student struct {
}

func (s *student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "okokokokok"
	} else {
		talk = "xxxxxxxxxxx"
	}

	return talk
}

func main() {
	stu := student{}

	result := stu.Speak("bitch")
	fmt.Println("result = ", result)
}
