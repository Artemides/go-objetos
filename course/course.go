package course

import "fmt"

type course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Clases  map[uint]string
}

func New(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 10
	}
	return &course{
		Name:   name,
		Price:  price,
		IsFree: isFree,
	}
}

func (c *course) ChangePrice(price float64) {
	c.Price = price
}
func (c *course) GetCursos() {
	text := "Los cursos son: "
	for _, clases := range c.Clases {
		text += clases + ", "
	}
	fmt.Println(text[:len(text)-2])
}
