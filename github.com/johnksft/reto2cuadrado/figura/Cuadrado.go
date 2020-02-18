package figura

type Cuadrado struct {
	Alto   float32
	Largo   float32
	Area float32
	Perimetro float32
}

func (c *Cuadrado) CalcPerimeter() {
	c.Perimetro = (c.Alto*2)+(c.Largo*2)
}

func (c *Cuadrado) CalcArea() {
	c.Area = (c.Alto * c.Largo)
}
