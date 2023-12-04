package main

import "fmt"

// Структура собаки
type Dog struct{}

// Метод для собаки
func (dog *Dog) WoofWoof() {
	fmt.Println("Гав-Гав!")
}

// Структура кошки
type Cat struct{}

// Метод кошки
func (cat *Cat) MeowMeow(isCall bool) {
	if isCall {
		fmt.Println("Ты меня позвал, мяу-мяу!")
	}
}

// Адаптер для наших структур
type AnimalAdapter interface {
	Reaction()
}

// Адаптер для собаки
type DogAdapter struct {
	*Dog
}

// Метод для адаптера собаки
func (adapter *DogAdapter) Reaction() {
	adapter.WoofWoof()
}

// Конструктор адаптера для собаки
func NewDogAdapter(dog *Dog) AnimalAdapter {
	return &DogAdapter{dog}
}

// Адаптер для кошки
type CatAdapter struct {
	*Cat
}

// Метод для адаптера кошки
func (adapter *CatAdapter) Reaction() {
	// адаптер автоматически зовет кошку isCall = true
	adapter.MeowMeow(true)
}

// Конструктор для кошки
func NewCatAdapter(cat *Cat) AnimalAdapter {
	return &CatAdapter{cat}
}

// Структура жены
type Wife struct {
}

// Метод жены
func (w *Wife) Reaction() {
	fmt.Println("Привет, Дорогой")
}

func main() {
	myFamily := [3]AnimalAdapter{NewDogAdapter(&Dog{}), NewCatAdapter(&Cat{}), &Wife{}}
	for _, member := range myFamily {
		member.Reaction()
	}
}
