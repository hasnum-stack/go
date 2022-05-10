package main

import "fmt"

type Person struct {
	name string
}

type Person1 struct {
	name string
	age  int
}

func test5() {
	a := true

	b_a := &a
	b_b := b_a

	*b_a = false

	fmt.Printf("b_a: %v\n", *b_a)
	fmt.Printf("b_b: %v\n", *b_b)

	// b_b:
}

func test4() {
	p1 := Person1{"zzh", 20}
	p2 := p1
	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("p2: %v\n", p2)

	fmt.Printf("p_p1: %p\n", &p1)
	fmt.Printf("p_p2: %p\n", &p2)

	p1.age = 25
	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("p2: %v\n", p2)

	p3 := &Person1{"zzh", 20}
	p4 := p3
	fmt.Printf("p3: %v\n", p3)
	fmt.Printf("p4: %v\n", p4)

	fmt.Printf("p3: %p\n", p3)
	fmt.Printf("p4: %p\n", p4)

	p3.age = 18
	fmt.Printf("p3: %v\n", p3)
	fmt.Printf("p4: %v\n", p4)
}

func test1() {
	person := Person{"zzh"}
	fmt.Printf("person: %v\n", person)
	fmt.Printf("person.name: %v\n", person.name)

	p_p := &person
	fmt.Printf("p: %p\n", p_p)

	p2 := new(Person)
	fmt.Printf("p2: %t\n", p2)
}

func showPerson(person *Person1) {
	person.age = 20
	fmt.Printf("%v", person.age)
}

func showPerson2(person Person1) {
	person.age = 20
	fmt.Printf("%v", person.age)
}

func test3() {
	p := Person1{"zzh", 19}
	showPerson2(p)
	fmt.Printf("p: %v\n", p)
}

func test2() {
	p := &(Person1{"zzh", 19})
	showPerson(p)
	fmt.Printf("p: %v\n", p)
}
func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	test5()
}
