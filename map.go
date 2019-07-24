package main
import "fmt"

func main(){

testMap := make(map[int] string)

testMap[141]="Mark"

fmt.Println(testMap[141])

fmt.Println(len(testMap))

testMap[142]="Holly"
fmt.Println(len(testMap))

for k,v := range testMap{
	fmt.Println(k, v)
}

delete(testMap,142)
fmt.Println(len(testMap))
 
}