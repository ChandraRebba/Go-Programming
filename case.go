package main
import "fmt"
func main(){

var age int
fmt.Println("What is your age")
fmt.Scan(&age)

if age>=15 && age<18{
	fmt.Println("You can drive with a learners permit")
}else if age>=18 {
	fmt.Println("You can get a license to drive")
}else{
	fmt.Println("You cannot get a license to drive")
}

switch {
case age>=16 && age<=18: fmt.Println("Can get learners permit")
case age>=18: fmt.Println("Can vote")
default: fmt.Println("User underage")
}

}
	
