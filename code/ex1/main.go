package main

import "fmt"

// Основная структура отчета. Применяем встраивание
type report struct {
	temperature
	location
	timeInterval
}

// Структура координат
type location struct {
	lat  float64
	long float64
}

// Структура температур
type temperature struct {
	high float64
	low  float64
}

// Структура проведенного времени
type timeInterval struct {
	days int
}

// Метод для температуры
func (t temperature) average() float64 {
	return (t.high + t.low) / 2
}

// Методы для координат
func (l location) getLat() float64 {
	return l.lat
}

func (l location) getLong() float64 {
	return l.long
}

// Структура для времени
func (t timeInterval) getTime() int {
	return t.days
}

func main() {
	report := report{
		temperature{
			high: 1.1,
			low:  -4.8,
		},
		location{
			lat:  55.7522200,
			long: 37.6155600,
		},
		timeInterval{
			5,
		},
	}

	// Проверка работы вызова методов при встраивании
	fmt.Println(report.average())
	fmt.Println(report.temperature.average())

	fmt.Println(report.getTime())
	fmt.Println(report.timeInterval.getTime())

	fmt.Println(report.getLat())
	fmt.Println(report.location.getLat())
}
