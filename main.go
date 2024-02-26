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

// MOvies represents a movie with its details.
type Movie struct {
	Id       string `json:"id"`       // Unique identifier of the movie.
	Name     string `json:"name"`     // Name of the movie.
	Title     string `json:"title"`     // Year of release of the movie.
	Director *Director `json:"director"` // Director of the movie.
}

//  the field tag is used to tell Go 
// how to convert the struct field to JSON and vice versa

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func main(){
	

	movies= append(movies, Movie{Id: "1", Name: "The Shawshank Redemption", Title: "1994", Director:  &Director{Firstname: "Frank",Lastname: "Darabont"}})
	movies = append(movies, Movie{Id: "2", Name: "The Godfather", Title: "1972", Director: &Director{Firstname: "Francis",Lastname: "Ford Coppola"}})
	movies = append(movies, Movie{Id: "3", Name: "The Dark Knight", Title: "2008", Director: &Director{Firstname: "Christopher",Lastname: "Nolan"}})
	movies = append(movies, Movie{Id: "4", Name: "The Godfather: Part II", Title: "1974", Director: &Director{Firstname: "Francis",Lastname: "Ford Coppola"}})
	



	r:=mux.NewRouter()
	r.HandleFunc("/movies",getmovies).Methods("GET")
	r.HandleFunc("/movie/{id}",getmovie).Methods("GET")
	r.HandleFunc("/movies",setmovies).Methods("POST")
	r.HandleFunc("/movies/{id}",updatemovies).Methods("PUT")
	r.HandleFunc("/movies/{id}",delmovies).Methods("DELETE")	


	fmt.Println("server has been initiated successfully")
	http.Handle("/", r)
log.Fatal(http.ListenAndServe(":8080", nil))

}

func getmovies(w http.ResponseWriter, r *http.Request){

	// 𝘁𝗵𝗶𝘀 𝗰𝗼𝗱𝗲 𝗰𝗵𝗮𝘀𝗻𝗴𝗲 𝘁𝗵𝗲 𝗵𝗲𝗮𝗱𝗲𝗿 𝗼𝗳 𝘁𝗵𝗲 𝗿𝗲𝘀𝗽𝗼𝗻𝘀𝗲 𝘁𝗼 𝗷𝘀𝗼𝗻
	w.Header().Set("Content-Type", "application/json")

	//𝘀𝗲𝗻𝗱𝗶𝗻𝗴 𝘁𝗵𝗲 𝗝𝗦𝗢𝗡 𝗿𝗲𝘀𝗽𝗼𝗻𝘀𝗲
	json.NewEncoder(w).Encode(movies)
}
func getmovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _,item := range movies{
		if item.Id == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Movie{})
}
func setmovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var temp Movie
	// .Decode(&temp) is a method call on the JSON decoder. 
	// It reads the next JSON-encoded value from its input 
	// (which is r.Body in this case) and stores it in the temp variable.
	_=json.NewDecoder(r.Body).Decode(&temp)
	// 𝗮𝗽𝗽𝗲𝗻𝗱𝗶𝗻𝗴 𝗮 𝗻𝗲𝘄 𝗶𝗱 𝘁𝗼 𝘁𝗵𝗲 𝗺𝗼𝘃𝗶𝗲
	temp.Id = strconv.Itoa(rand.Intn(100000000))
	movies=append(movies, temp)

	// 𝗿𝗲𝘁𝘂𝗿𝗻𝗶𝗻𝗴 𝘁𝗵𝗲 𝗷𝘀𝗼𝗻 𝗿𝗲𝘀𝗽𝗼𝗻𝘀𝗲
	json.NewEncoder(w).Encode(temp)
}
// Handler for the update endpoint
func updatemovies(w http.ResponseWriter, r *http.Request) {
    // Extract the ID from the URL
    vars := mux.Vars(r)
    movieID := vars["id"]

    // Check if the movie exists
    var movieIndex int
    var movieExists bool
    for i, movie := range movies {
        if movie.Id == movieID {
            movieIndex = i
            movieExists = true
            break
        }
    }
    if !movieExists {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w, "Movie with ID %s not found", movieID)
        return
    }

    // Decode the JSON payload into a Movie struct
    var updatedMovie Movie
    err := json.NewDecoder(r.Body).Decode(&updatedMovie)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Error decoding JSON: %v", err)
        return
    }

    // Update the movie with the new data
    movies[movieIndex] = updatedMovie

    // Respond with the updated movie
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(updatedMovie)
}



func delmovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for i,item :=range movies {
		if item.Id == params["id"]{

			//  slice of all movies before the movie to be appended (movies[:i]),
			// and (movies[i+1:])all movies after the movie to be appended
			movies = append(movies[:i], movies[i+1:]...)
			break
		}


	}

}
