package main

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
Visitor позволяет добавить дополнительный функционал для различных объектов,
при этом сильно не меняя их структуру, по сути реализуется поведение, которое должно быть свойственно
самому объекту, но не может быть ему присвоенно из-за отсутствия возможности рефакторить оригинального код

+ позволяет собственно добовлять новый функционал в сложную систему
- плохо себя показывает в случае частого изменения иерархии элементов

В качестве примера можно привести, ситуация, когда нужно проделать какую-либо операцию над объектами различных классов
*/

// Visitor
type Doctor interface {
	checkForDog()
	checkForHuman()
}

// Наследник Visitor'a
type Vet struct{}

func (this *Vet) checkForDog() {
	fmt.Println("Собака проверена")
}

func (this *Vet) checkForHuman() {
	fmt.Println("Я не специализируюсь на людях")
}

// Наследник Visitor'a
type HumanDoctor struct{}

func (this *HumanDoctor) checkForHuman() {
	fmt.Println("Человек проверен")
}

func (this *HumanDoctor) checkForDog() {
	fmt.Println("Я не спициализируюсь на собаках")
}

type Animal interface {
	accept(d *Doctor)
}

type Dog struct{}

func (this *Dog) accept(d Doctor) {
	d.checkForDog()
}

func main() {
	dog := new(Dog)
	doctor := new(Vet)
	doctor2 := new(HumanDoctor)

	dog.accept(doctor)
	dog.accept(doctor2)
}
