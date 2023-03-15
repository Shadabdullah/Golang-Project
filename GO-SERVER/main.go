package main

import(
	"fmt"
	"log"
	"net/http"
)

func formHandler( w http.ResponseWriter,r *http.Request){

	if err := r.ParseForm(); err != nil{

		fmt.Fprint(w,"ParseForm() err:",err)
		return
	}
	fmt.Fprint(w,"Post request succesfull \n")
	name:=r.FormValue("name")
	address:=r.FormValue("address")
	fmt.Fprint(w,"Name : \n",name)
	fmt.Fprint(w,"Address:",address)

}
func helloHandler( w http.ResponseWriter,r *http.Request){
if r.URL.Path != "/hello"{

	http.Error(w,"404 not found",http.StatusNotFound)
	return
}
if r.Method != "GET"{
	http.Error(w,"Method not supported",http.StatusNotFound)
	return
}
fmt.Fprint(w,"hello!")
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
