#Go2Json
  
Parse JSON file and create Go file declaring corresponding structs 

## Usage
```
go run main.go <input_file> <output_file> <root_struct> <package_name>
```
For example:
```
go run main.go ./input/test1.json ./output/test1.go Person person
```

## Example
Input JSON:
```
{
  "name": "Brand",
  "married": true,
  "company": {
    "Name":"Lexus",
    "location": "Japan",
    "num_employees": 40000,
    "departments": [
      {
        "name": "sales",
        "num_employees":3000
      },
      {
        "name": "tech"
      }
    ]
  },
  "Friends": [{
    "name": "Mike",
    "age": 24
  },
    {
    "name": "Tom"
    },{
      "name": "Johnson",
      "age": 23
    }]
}
```
<br/>

Output Go file:
```
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
```