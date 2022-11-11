package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to CRM Backend Project")
}

func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)

}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	Id := mux.Vars(r)["id"]

	for _, customer := range customers {

		if customer.Id == Id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(customer)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "CUSTOMER NOT FOUND")
}

func createCustomer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var newCustomer Customer
	reqBody, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(reqBody, &newCustomer)

	//Customers in the database have unique ID values
	newCustomer.Id = uuid.New().String()

	customers = append(customers, newCustomer)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(customers)

}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	Id := mux.Vars(r)["id"]

	for idx, customer := range customers {

		if customer.Id == Id {
			w.WriteHeader(http.StatusOK)
			customers = append(customers[:idx], customers[idx+1:]...)
			json.NewEncoder(w).Encode(customer)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "CUSTOMER NOT FOUND")
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	Id := mux.Vars(r)["id"]

	var newCustomer Customer
	reqBody, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(reqBody, &newCustomer)

	for idx, customer := range customers {
		if customer.Id == Id {
			c := &customers[idx]

			c.Id = customer.Id

			if len(newCustomer.Name) != 0 {
				c.Name = newCustomer.Name
			} else {
				c.Name = customer.Name
			}
			if len(newCustomer.Email) != 0 {
				c.Email = newCustomer.Email
			} else {
				c.Email = customer.Email
			}
			if len(newCustomer.Role) != 0 {
				c.Role = newCustomer.Role
			} else {
				c.Role = customer.Role
			}
			if newCustomer.Phone != 0 {
				c.Phone = newCustomer.Phone
			} else {
				c.Phone = customer.Phone

			}

			c.Contacted = newCustomer.Contacted
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(customers)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "CUSTOMER NOT FOUND")
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", index).Methods("GET")

	// Getting a list of all customers
	router.HandleFunc("/customers", getAllCustomer).Methods("GET")

	//Getting data for a single customer
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")

	//Adding a customer
	router.HandleFunc("/customers", createCustomer).Methods("POST")

	//Removing a customer
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	//Updating a customer's information
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PATCH")

	fmt.Println("Server is starting...")
	http.ListenAndServe(":3000", router)
}
