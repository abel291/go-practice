package countries

import (
	"errors"
	"fmt"
)

var country string
var collection []string
var myCountry string
var myCountryMsg string
var errorNoFound = errors.New("pais no encontrado")

// Add agrega el valor de un pais a la collecion de paises.
func Add(country string) {
	collection = append(collection, country)
}

func isInColletion(country string) bool {
	for _, c := range collection {
		if c == country {
			return true
		}
	}
	return false
}

// SetMyCountry seleciona un pais como pais favorito
func SetMyCountry(country string) error {
	is_collection := isInColletion(country)
	if !is_collection {
		return errorNoFound
	}
	myCountry = country
	return nil
}

// List lista los paises
func List() {
	for i, c := range collection {
		myCountryMsg := ""
		if c == myCountry {
			myCountryMsg = "my COUNTRY"
		}
		fmt.Printf("%d. %s %s \n", i+1, c, myCountryMsg)
	}
}
