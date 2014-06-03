package main

import (
    "fmt"
)

type Animal interface {
    GetName() string
    GetKind() string
    Hello()
}

// dog
type Dog struct {
    name string
}

func (dog Dog) GetName() string {
    return dog.name
}

func (dog Dog) GetKind() string {
    return "dog"
}

func (dog Dog) Hello() {
    fmt.Println("bow-wow!")
}

// cat
type Cat struct {
    name string
}

func (cat Cat) GetName() string {
    return cat.name
}

func (cat Cat) GetKind() string {
    return "cat"
}

func (cat Cat) Hello() {
    fmt.Println("meow")
}

func ShowAnimal(animal Animal) {
    fmt.Printf("animal's name: %s\n", animal.GetName())
    fmt.Printf("animal's kind: %s\n", animal.GetKind())
    animal.Hello()
}

func main() {
    var animal Animal

    animal = Dog{"hachi"}
    ShowAnimal(animal);

    animal = Cat{"tama"}
    ShowAnimal(animal);

    animal2 := new(Dog)
    animal2.name = "pochi"
    ShowAnimal(animal2)
}