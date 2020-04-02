package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const headerTemplate = "package %s\n\n"

type Class struct {
	Name string
	Data map[string]interface{}
}

func main() {
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	rootStructName := os.Args[3]
	packageName := os.Args[4]

	jsonBytes, err := readFile(inputFile)
	if err != nil {
		panic(fmt.Errorf("reading file got error: %s", err))
	}

	code := jsonToGoStructs(jsonBytes, rootStructName, packageName)

	if err = writeToFile(outputFile, code.String()); err != nil {
		panic(fmt.Errorf("writing file got error: %s", err))
	}
}

// jsonToGoStructs parses json byte array and try to deduce corresponding Go struct declaration
func jsonToGoStructs(jsonBytes []byte, rootStructName string, packageName string) *Code {

	var parsed map[string]interface{}
	json.Unmarshal(jsonBytes, &parsed)

	code := &Code{
		Header: fmt.Sprintf(headerTemplate, packageName),
		Body:   StructDefMap{},
	}

	// traverse all classes/structs and fields in BFS
	classes := []Class{{Name: rootStructName, Data: parsed}}
	for i := 0; i < len(classes); i++ {
		class := classes[i]
		def := StructDef{}
		if v, ok := code.Body[underscoreToCamel(class.Name)]; ok {
			def = v
		} else {
			code.Body[underscoreToCamel(class.Name)] = def
		}

		for field, val := range class.Data {
			// if the val is a array or nested array, find the parsed type of its elements
			subClasses, arrayDimensions := getNestedSlice(field, val)
			for _, c := range subClasses {
				classes = append(classes, c)
			}

			// generate array annotation like [][] based on dimensions of array
			arrayAnnotation := ""
			for j := 0; j < arrayDimensions; j++ {
				arrayAnnotation += "[]"
			}

			// generate field definition based on value type
			fieldName := underscoreToCamel(field)
			jsonTag := field
			fieldType := ""
			switch v := val.(type) {
			case map[string]interface{}:
				if len(arrayAnnotation) > 0 {
					// change field type word from plural to singular if it's a slice field
					field = strings.TrimRight(field, "s")
				}
				classes = append(classes, Class{Name: field, Data: v})
				fieldType = underscoreToCamel(field)
			case int, int32, int64:
				fieldType = "int64"
			case float64: // actually all json numbers will be parsed as float64
				fieldType = "float64"
			case bool: // actually all json numbers will be parsed as float64
				fieldType = "bool"
			default:
				fieldType = "string"
			}

			fieldType = arrayAnnotation + fieldType
			def[fieldName] = Field{Name: fieldName, Type: fieldType, JsonTag: jsonTag}
		}

	}
	return code
}

func getNestedSlice(name string, input interface{}) ([]Class, int) {
	slice, ok := input.([]interface{})
	if ok {
		res := []Class{}
		resLvl := 0
		for _, ele := range slice {
			children, lvl := getNestedSlice(name, ele)
			resLvl = lvl + 1
			res = append(res, children...)
		}
		return res, resLvl
	}
	dict, ok := input.(map[string]interface{})
	if ok {
		return []Class{{Name: name, Data: dict}}, 0
	}
	return nil, 0
}

// StructDefMap is a slice of declaration of structure type
type StructDefMap map[string]StructDef

type StructDef map[string]Field

type Field struct {
	Name    string
	Type    string
	JsonTag string
}

type Code struct {
	Header string
	Body   StructDefMap
}

func (s StructDefMap) String() string {
	res := []string{}
	for class, fieldMap := range s {
		structDef := fmt.Sprintf("type %s struct {\n", underscoreToCamel(class))
		for _, def := range fieldMap {
			structDef += fmt.Sprintf("\t%s %s `json:\"%s\"`\n", def.Name, def.Type, def.JsonTag)
		}
		structDef += "}"
		res = append(res, structDef)
	}

	return strings.Join(res, "\n\n")
}

func (c *Code) String() string {
	return c.Header + c.Body.String()
}

func underscoreToCamel(input string) string {
	tokens := strings.Split(input, "_")
	for i, token := range tokens {
		tokens[i] = strings.Title(token)
	}
	return strings.Join(tokens, "")
}

func readFile(filePath string) ([]byte, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	bytes, _ := ioutil.ReadAll(jsonFile)
	return bytes, nil
}

func writeToFile(filePath string, content string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
