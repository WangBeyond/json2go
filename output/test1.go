package person

type Person struct {
	Name string `json:"name"`
	Married string `json:"married"`
	Company Company `json:"company"`
	Friends []Friend `json:"Friends"`
}

type Company struct {
	Name string `json:"Name"`
	Location string `json:"location"`
	NumEmployees float64 `json:"num_employees"`
	Departments []Department `json:"departments"`
}

type Friend struct {
	Name string `json:"name"`
	Age float64 `json:"age"`
}

type Department struct {
	Name string `json:"name"`
	NumEmployees float64 `json:"num_employees"`
}