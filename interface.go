package main
import "fmt"
import "math"

type Shape interface{
	area() float64
}

type Rectangle struct{
	height float64
	width float64
}

type Circle struct{
	radius float64
}

func (r Rectangle) area() float64{
	return r.height * r.width
}

func (c Circle) area() float64{
	return math.Pi * math.Pow(c.radius,2)
}

func getArea(shape Shape) float64{
	return shape.area()
}

func main(){
	rect := Rectangle{20,10}
	circ := Circle{6}

	fmt.Println("The area of Rectangle is",getArea(rect))
	fmt.Println("The area of the Circle is",getArea(circ))
}