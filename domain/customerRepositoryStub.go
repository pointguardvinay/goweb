package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {

	customers := []Customer{

		{Id: "1000", Name: "Vinay", City: "Thergaon", ZipCode: "411033", DateofBirth: "2000-01-01", Status: "1"},
		{Id: "1001", Name: "John", City: "Delhi", ZipCode: "452000", DateofBirth: "1988-01-01", Status: "1"},
		{Id: "1002", Name: "HHarry", City: "Bijeing", ZipCode: "1023255", DateofBirth: "2000-08-09", Status: "1"},
	}

	return CustomerRepositoryStub{
		customers,
	}
}
