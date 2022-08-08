package main

import (
	"fmt"
	"log"
)

var HELLO = "Привіт"

const name = "Вася"

func (s Square) Perimeter() {
	fmt.Printf("%T, %#v \n", s, s)
	fmt.Printf("Периметр фігури %d \n", s.Side*4)
}

func (s *Square) Scale(multiplier int) {
	fmt.Printf("\n%T, %#v \n", s, s)
	s.Side *= multiplier
	fmt.Printf("%T, %#v \n", s, s)
}

func (s Square) WrongScale(multiplier int) {
	fmt.Printf("\n%T, %#v \n", s, s)
	s.Side *= multiplier
	fmt.Printf("%T, %#v \n", s, s)
}

func RunLesson12() {
	//	var name = "hello"
	//	key := true
	//	fmt.Printf("Type: %T, Value %s\n", HELLO, HELLO)
	//	fmt.Printf("Type: %T, Value %s \n", name, name)
	//	fmt.Printf("Type: %T, Value %v\n", key, key)
	//	result := fmt.Sprintf("%s довбойоб", name)
	//	fmt.Println(result)
	//	var num1 int64 = 100
	//	var num2 int32 = 200
	//	fmt.Println(num1 + int64(num2))
	//	fmt.Println(unsafe.Sizeof(num1))
	//	fmt.Println(unsafe.Sizeof(num2))
	//Greet()
	//GreetWithName("Сергій", "Москаленко")
	//
	//first, second := 1, 2
	//sum := Sum(first, second)
	//fmt.Println(sum)

	square := Square{Side: 4}
	pSquare := &square

	square2 := Square{Side: 2}

	square.Perimeter()
	square2.Perimeter()

	pSquare.Scale(2)
	pSquare.Perimeter()

	square.Scale(2)
	pSquare.Perimeter()

	square.WrongScale(2)
	square.Perimeter()

}

func Greet() {
	log.Println("Ping")
}

func GreetWithName(name string, surname string) {
	log.Println("Name: " + name + ", surname: " + surname)
}

func Sum(number1 int, number2 int) (sum int) {
	return number1 + number2
}

type Square struct {
	Side int
}
