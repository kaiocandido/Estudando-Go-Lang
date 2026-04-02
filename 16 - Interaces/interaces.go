package main

type retangulo struct {
	largura float64
	altura  float64
}

type circulo struct {
	raio float64
}

// interface forma geometrica que tem o metodo area() que retorna um float64
// interace serve para definir um contrato, ou seja, qualquer tipo que implemente o metodo area() pode ser considerado uma forma geometrica
type formaGeometrica interface {
	area() float64
}

func area(r formaGeometrica) {
	area := r.area()
	println(area)
}

func (r retangulo) area() float64 {
	return r.largura * r.altura
}

func (c circulo) area() float64 {
	return 3.14 * c.raio * c.raio
}

func main() {
	r := retangulo{10, 5}
	c := circulo{5}
	area(r)
	area(c)
}
