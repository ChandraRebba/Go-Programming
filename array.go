package main
import "fmt"

func main(){

	var arrayTest[5] float64
	arrayTest[0]=3.14
	arrayTest[1]=4.6
	arrayTest[2]=247728
	arrayTest[3]=36.33
	arrayTest[4]=499.21
	//using basic recursion
	for i := 0; i < len(arrayTest); i++ {
		fmt.Println("using recursion",arrayTest[i])
	}
	//using range
	for i,value := range arrayTest {

		fmt.Println(value,i)
	}
	//Slices
	sliceTest :=[]int{2,4,6,8,10}
	sliceTest2 := sliceTest[2:5]
	fmt.Println(sliceTest2[0:2])
	sliceTest3:=make([]int,5,10)
	copy(sliceTest3,sliceTest)
	sliceTest3=append(sliceTest3,0,-1)
	fmt.Println(sliceTest3[0:])
}