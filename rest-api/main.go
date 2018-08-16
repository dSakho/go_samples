package main

/*
 * Example showing how convenient it is to create a REST API with 1 file
 */

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// Person data struct
type Person struct {
	ID        int      `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address data struct
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

// Route handler to return all people
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// Route handler to return a single person
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		idIntVal, _ := strconv.Atoi(params["id"])
		if item.ID == idIntVal {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// Route Handler to create a CSV and store the amount of people in the people array
func PersistPeople(w http.ResponseWriter, r *http.Request) {
	headerRow := []string{"id", "firstName", "lastName", "city", "state"}
	rows := make([][]string, len(people))

	for _, person := range people {
		var row []string
		row = append(row, strconv.Itoa(person.ID))
		row = append(row, person.Firstname)
		row = append(row, person.Lastname)
		row = append(row, person.Address.City)
		row = append(row, person.Address.State)
		rows = append(rows, row)
	}

	csvFile, err := os.Create("people.csv")
	if err != nil {
		log.Fatalln(err)
	}

	csvWriter := csv.NewWriter(csvFile)

	csvWriter.Write(headerRow)
	// Write any buffered data to the underlying writer (standard output).
	csvWriter.Flush()

	for _, record := range rows {
		if err := csvWriter.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		log.Fatal(err)
	}
}

// Route handler to create a person
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var newPerson Person
	_ = json.NewDecoder(r.Body).Decode(&newPerson)
	newPerson.ID = len(people) + 1
	people = append(people, newPerson)
	json.NewEncoder(w).Encode(newPerson)
}

// Route handler to update a person
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		idIntVal, _ := strconv.Atoi(params["id"])
		if item.ID == idIntVal {
			var requestBody Person
			_ = json.NewDecoder(r.Body).Decode(&requestBody)

			people[index].Firstname = requestBody.Firstname
			people[index].Lastname = requestBody.Lastname
			people[index].Address = requestBody.Address
			break
		}
		json.NewEncoder(w).Encode(people[index])
	}

}

// Delete a person
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		idIntVal, _ := strconv.Atoi(params["id"])
		if item.ID == idIntVal {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people", CreatePerson).Methods("POST")
	router.HandleFunc("/people/tocsv", PersistPeople).Methods("POST")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", UpdatePerson).Methods("PUT")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	log.Print("API is ready to use!")

	log.Fatal(http.ListenAndServe(":8000", router))
}
