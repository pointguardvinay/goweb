package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/poinguardvinay/goweb/service"
)

type customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	customeres, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {

		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customeres)

	} else {
		w.Header().Add("Content-Type", "application/xml")
		json.NewEncoder(w).Encode(customeres)
	}

}

/*
func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers := []customer{
		{Name: "Vinay", City: "Solapur", Zipcode: "41302"},
		{Name: "Shailu", City: "Pune", Zipcode: "411033"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {

		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}

/*
func getCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	fmt.Fprintf(w, vars["customer_id"])

}

func catHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "This is coming from the function")
}

func createCustomer(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Flat Post Request received")

}
*/
