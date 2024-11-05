package log

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/tabwriter"
)

func Table[T any](data []T) {
	// Ensure data is a slice of structs
	v := reflect.ValueOf(data)
	if v.Len() == 0 || v.Index(0).Kind() != reflect.Struct {
		fmt.Println("Error: Input must be a slice of structs")
		return
	}

	// Get type information from the first element in the slice
	structType := v.Index(0).Type()

	// Extract field names (table headers)
	headers := map[string]int{}
	var headerOrder []string
	for i := 0; i < structType.NumField(); i++ {
		headerName := structType.Field(i).Name
		headers[headerName] = len(headerName)
		headerOrder = append(headerOrder, headerName)
	}

	for _, item := range data {
		itemValue := reflect.ValueOf(item)
		for headerName, length := range headers {
			fieldValue := itemValue.FieldByName(headerName)
			fieldStr := fmt.Sprintf("%v", fieldValue)
			fieldLen := len(fieldStr)
			if fieldLen > length {
				headers[headerName] = fieldLen
			}
		}
	}

	// Create a new tab writer
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	headerFormat := ""
	underlineFormat := ""
	var valuesFormat []any
	for _, header := range headerOrder {
		length := headers[header]
		headerFormat += fmt.Sprintf("%%-%ds\t", length)
		underlineFormat += strings.Repeat("=", length) + "\t"
		valuesFormat = append(valuesFormat, header)
	}

	// Print headers
	fmt.Fprintf(writer, headerFormat+"\n", valuesFormat...)
	fmt.Fprintf(writer, underlineFormat+"\n\n") //nolint:govet

	// Print rows with dynamic widths
	for _, item := range data {
		rowFormat := []any{}
		itemValue := reflect.ValueOf(item)
		for _, header := range headerOrder {
			fieldValue := fmt.Sprintf("%v", itemValue.FieldByName(header))
			rowFormat = append(rowFormat, fieldValue)
		}
		fmt.Fprintf(writer, headerFormat+"\n", rowFormat...)
	}

	// Flush to ensure all data is written
	writer.Flush()
}
