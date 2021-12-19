package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Gena", City: "Kiev", Zipcode: "110001", DateofBirth: "1990-10-21", Status: "1"},
		{Id: "1002", Name: "Tanya", City: "Kiev", Zipcode: "110001", DateofBirth: "1992-10-21", Status: "1"},
	}

	return CustomerRepositoryStub{customers}
}
