package main
import ("fmt")

func add(a,b int) int{
	return a+b
}

func sub(a,b int) int{
	return a-b
}

func mul (a,b int) int{
	return a*b
}

func div(a,b float64) float64{
	return a/b
}





func main(){

	var a int
	var b int
	fmt.Println("Enter inputs")
	 fmt.Scanln(&a)
	 fmt.Scanln(&b)
	var c string
	fmt.Println("Enter Function to be performed")
	fmt.Scanln(&c)
	switch {
	case c=="add":fmt.Println("Output:",add(a,b))
	case c=="sub":fmt.Println("Output:",sub(a,b))
	case c=="mul":fmt.Println("Output:",mul(a,b))
	case c=="div":fmt.Println("Output:",div(float64(a),float64(b)))
	default:fmt.Println("Please enter valid Function")
	}}

