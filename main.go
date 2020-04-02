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
		Body:   []string{},
	}

	// traverse all classes/structs and fields in BFS
	classes := []Class{{Name: rootStructName, Data: parsed}}
	for i := 0; i < len(classes); i++ {
		def := ""
		class := classes[i]
		def += fmt.Sprintf("type %s struct {\n", underscoreToCamel(class.Name))
		for field, val := range class.Data {

			// if the val is a array or nested array, find the parsed type of its elements
			arrayDimensions := 0
			for {
				slice, ok := val.([]interface{})
				if !ok {
					break
				}
				arrayDimensions++
				if len(slice) == 0 {
					break
				}
				val = slice[0]
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
			default:
				fieldType = "string"
			}

			fieldType = arrayAnnotation + fieldType
			def += fmt.Sprintf("\t%s %s `json:\"%s\"`\n", fieldName, fieldType, jsonTag)
		}
		def += "}"
		code.Body = append(code.Body, def)
	}
	return code
}

// StructDefs is a slice of declaration of structure type
type StructDefs []string

type Code struct {
	Header string
	Body   StructDefs
}

func (s StructDefs) String() string {
	return strings.Join(s, "\n\n")
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
