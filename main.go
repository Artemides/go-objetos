package main

import (
	"fmt"

	"github.com/Artemides/go-objetos/course"
)

func main() {
	Go := course.New("Go desde 0", 25, false)

	Go.UserIDs = []uint{22545, 14587, 45687, 485456}
	Go.Clases = map[uint]string{
		1: "Introducción",
		2: "Estructuras",
		3: "Ciclos",
		4: "Funciones",
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
