package main

import "fmt"

type Runner interface {
	Run() string
}

type Swimmer interface {
	Swim() string
}
type Flyer interface {
	Fly() string
}

type Ducker interface {
	Run() string
	Swim() string
	Fly() string
}

func (h Human) Run() string {
	return fmt.Sprintf("Прогер %s бігає", h.Name)
}

func (h Human) writeCode() string {
	return "Програміст пише код"
}

type Duck struct {
	Name, Surname string
}

func (d Duck) run() string {
	return "Качка бігає"
}

func (d Duck) Swim() string {
	return "Качка плаває"
}

func (d Duck) Fly() string {
	return "Качка літає"
}

type Human struct {
	Name string
}

func RunLesson13() {
	var runner Runner
	fmt.Printf("%T, %#v \n", runner, runner)

	if runner == nil {
		fmt.Println("Runner is null")
	}

	var unnanedRunner *Human
	fmt.Printf("Type: %T, Value %s\n", unnanedRunner, unnanedRunner)

	runner = unnanedRunner
	fmt.Printf("Type: %T, Value %s\n", runner, runner)

	namedRunner := &Human{Name: "Педро"}
	fmt.Printf("Type: %T, Value %s\n", namedRunner, namedRunner)

}
