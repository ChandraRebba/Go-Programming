package main
import("fmt"
"net/http")

func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w, "Hello \n")
}

func handler2(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello world")
}

func main(){
	http.HandleFunc("/",handler)
	http.HandleFunc("/world",handler2)
	http.ListenAndServe(":8080",nil)
}