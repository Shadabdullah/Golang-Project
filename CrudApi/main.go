package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct{ 
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}
type Director struct{

	Firstname string `json:"firstname"`
	Lastname string `json:"lasttname"`
}



var movies []Movie

// function

func getMovies(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	json.NewEncoder(w).Encode(movies)


}
func deletMovie(w http.ResponseWriter ,r *http.Request){

	w.Header().Set("Content-Type","application/json")
	param := mux.Vars(r)

	for index , item:= range movies{


		if item.ID== param["id"]{
			movies = append(movies[:index],movies[:index+1]...)
			break
		}
	}
json.NewEncoder(w).Encode(movies)

}
func getMovie(w http.ResponseWriter ,r *http.Request){

	w.Header().Set("Content-Type","application/json")
	param := mux.Vars(r)

	for _, item:= range movies{


		if item.ID== param["id"]{

			json.NewEncoder(w).Encode(item)
			return
		}
	}


}

func creatMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies,movie)
	json.NewEncoder(w).Encode(movie)


}




func updateMovie(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type","application/json")

	param := mux.Vars(r)

	for index ,item:= range movies{

		if param["id"] == item.ID{

			movies = append(movies[:index],movies[:index+1]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = param["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}


}


func main(){

	r:= mux.NewRouter()

	m1 := Movie{ID: "1", Isbn: "1234567890", Title: "The Godfather", Director: &Director{Firstname: "Francis", Lastname: "Ford Coppola"}}
	m2 := Movie{ID: "2", Isbn: "0987654321", Title: "The Shawshank Redemption", Director: &Director{Firstname: "Frank", Lastname: "Darabont"}}
	movies = append(movies, m1, m2)
	//Function to perform operation
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",creatMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deletMovie).Methods("DELETE")

	fmt.Print("server started at 8080\n")
	log.Fatal(http.ListenAndServe(":8080",r))

}


