package course

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
func (c *Course) GetCursos() {
	text := "Los cursos son: "

	for _, clases := range c.Clases {
		text += clases + ", "
	}
	fmt.Println(text[:len(text)-2])
}
