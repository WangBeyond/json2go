package person

type Person struct {
	Company Company `json:"company"`
	Friends []string `json:"Friends"`
	Name string `json:"name"`
	Married bool `json:"married"`
}

type Company struct {
	Profitable bool `json:"profitable"`
	Ceo string `json:"Ceo"`
	Name string `json:"name"`
	Location string `json:"location"`
	NumEmployees float64 `json:"num_employees"`
	Departments []string `json:"departments"`
}

type Friends struct {
	Age float64 `json:"age"`
	Company Company `json:"company"`
	Hobby string `json:"hobby"`
	Name string `json:"name"`
}

type Departments struct {
	Name string `json:"name"`
	NumEmployees float64 `json:"num_employees"`
}