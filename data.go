package main

type Customer struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     int    `json:"phone"`
	Contacted bool   `json:"contacted"`
}

var customers = []Customer{
	{
		Id:        "1",
		Name:      "Tom",
		Role:      "Software Engineer",
		Email:     "tom@gmail.com",
		Phone:     123456892,
		Contacted: true,
	},
	{
		Id:        "2",
		Name:      "Tim",
		Role:      "Senior Software Engineer",
		Email:     "tim@gmail.com",
		Phone:     213456892,
		Contacted: false,
	},
	{
		Id:        "3",
		Name:      "Max",
		Role:      "Product Manager",
		Email:     "max@gmail.com",
		Phone:     231564892,
		Contacted: true,
	},
	{
		Id:        "4",
		Name:      "Joe",
		Role:      "Software Engineer",
		Email:     "joe@gmail.com",
		Phone:     123456892,
		Contacted: true,
	},
	{
		Id:        "5",
		Name:      "Ross",
		Role:      "Civil Engineer",
		Email:     "ross@gmail.com",
		Phone:     123456892,
		Contacted: true,
	},
	{
		Id:        "6",
		Name:      "Adam",
		Role:      "IT Support Engineer",
		Email:     "adam@gmail.com",
		Phone:     123456892,
		Contacted: true,
	},
	{
		Id:        "7",
		Name:      "Jermy",
		Role:      "Software Engineer",
		Email:     "jermy@gmail.com",
		Phone:     123456892,
		Contacted: true,
	},
	{
		Id:        "8",
		Name:      "Sam",
		Role:      "Principal Engineer",
		Email:     "sam@gmail.com",
		Phone:     123456892,
		Contacted: true,
	},
	{
		Id:        "9",
		Name:      "Eric",
		Role:      "DevOps Engineer",
		Email:     "eric@gmail.com",
		Phone:     123456892,
		Contacted: true,
	},
	{
		Id:        "10",
		Name:      "David",
		Role:      "Junior Software Engineer",
		Email:     "david@gmail.com",
		Phone:     123456892,
		Contacted: true,
	},
}
