package main
import "fmt"

func factorial(n int) int{
if (n==0){
	return 1
}

return n*factorial(n-1)

}

func PrintOne(){
	fmt.Println(1)
}
func PrintTwo(){
	fmt.Println(2)
}

func Div(a,b int) int{
	defer func(){
		fmt.Println(recover())//helps the program run after fatal error
	}()
c:= a/b 
return c
}

func demPanic(){
	defer func(){
		fmt.Println(recover())
	}()
	panic("This is a custom error")//Prints out custom errors,use of defer and recover to handle the same error
}

func main(){

	defer PrintTwo()//the defer clause makes the func wait until all other functions are executed
	PrintOne()
	fmt.Println(Div(3,0))
	fmt.Println(Div(6,3))
	demPanic()
}
