package domain

// creating respository

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}

type Customer struct {
	Id          string
	Name        string
	City        string
	ZipCode     string
	DateofBirth string
	Status      string
}
