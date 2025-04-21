package main

import "fmt"

func main() {
	zeekr := Car{Brand: "Zeekr", Year: 2024,
		Country: Country{
			Name:    "China",
			Address: "China",
		},
	}

	updateCarYear(&zeekr)

	fmt.Println(zeekr)
	fmt.Println(zeekr.Year)

	x := 10
	p := &x // Получаем адрес переменной x

	fmt.Println("Значение x:", x)
	fmt.Println("Адрес x:", p)
	fmt.Println("Значение через указатель:", *p)

	*p = 20 // Меняем значение через указатель
	fmt.Println("Новое значение x:", x)
}

type Car struct {
	Brand string
	Year  int

	Country Country
}

type Country struct {
	Name    string
	Address string
}

func updateCarYear(c *Car) {
	c.Year++
}
