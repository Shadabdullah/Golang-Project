package main

import(
	"fmt"
	"log"
	"net/http"
)

func helloHandler(){


}
func formHandler (){


	
}

func main(){
fileServer:= http.FileServer(http.Dir("./static"))
http.Handle("/",fileServer)
http.HandleFunc("/form",formHandler)
http.HandleFunc("/hello",helloHandler)
fmt.Printf("server started at port 8000\n")
if err:= http.ListenAndServe(":8000",nil); err != nil{
	log.Fatal(err)
}

}