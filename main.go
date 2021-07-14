package main

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Clases  map[uint]string
}

func (c *Course) ChangePrice(price float64) {
	c.Price = price
}
func (Course) getName() {
	fmt.Println(" Holi")
}
func (c Course) GetCursos() {
	text := "Los cursos son: "

	for _, clases := range c.Clases {
		text += clases + ", "
	}
	fmt.Println(text[:len(text)-2])
}

func main() {
	Go := Course{
		Name:    "Go desde 0",
		Price:   12.75,
		IsFree:  false,
		UserIDs: []uint{22545, 14587, 45687, 485456},
		Clases: map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Ciclos",
			4: "Funciones",
		},
	}
	/*Css := Course{Name: "CSS", IsFree: false}
	Js := Course{}
	Js.Name = "JS para expertos"
	Js.Price = 59.90
	Js.UserIDs = []uint{544778, 74785, 7458}
	fmt.Printf("%+v\n", Go)
	fmt.Printf("%+v\n", Css)
	fmt.Printf("%+v\n", Js)*/
	Go.GetCursos()
	Go.ChangePrice(44.5)
	fmt.Println(Go.Price)
}
