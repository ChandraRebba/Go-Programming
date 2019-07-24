package main
import "fmt"

type Library struct{

	bookID int
	bookName string
	Author string
}

type Rectangle struct{

	height float64
	width float64
}

func (rect *Rectangle) area() float64{

return rect.width * rect.height

}

func main(){

	book1 := Library{bookID: 111, bookName:"Harry Potter", Author:"J.K Rowling"}
	fmt.Println("The book available is",book1.bookName)
	rect := Rectangle{height:10,width:20}
	fmt.Println("The area of the Rectangle is",rect.area())

}