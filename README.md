# Go2Json
  
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
    "Name": "Lexus",
    "location": "Japan",
    "num_employees": 40000,
    "departments": [
      {
        "name": "sales",
        "num_employees": 3000
      },
      {
        "name": "tech"
      }
    ]
  },
  "Friends": [
    {
      "name": "Mike",
      "age": 24,
      "company": {
        "name": "Toyota",
        "profitable": true
      }
    },
    {
      "name": "Tom",
      "hobby": "cycling",
      "company": {
        "name": "Dyson",
        "Ceo": "Somebody"
      }
    },
    {
      "name": "Johnson",
      "age": 23,
      "hobby": "drawing"
    }
  ]
}
```
<br/>

Output Go file:
```
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
```