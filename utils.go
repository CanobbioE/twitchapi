package twitchapi

import (
	"errors"
	"fmt"
	"reflect"
)

// parseInput transform an input struct into a map
func parseInput(input interface{}) map[string]interface{} {
	s := reflect.ValueOf(input)
	params := make(map[string]interface{})

	for i := 0; i < s.NumField(); i++ {
		curr := s.Field(i).Interface()
		tag := fieldTag(reflect.TypeOf(input).FieldByIndex([]int{i}))
		if !isNil(curr) {
			params[tag] = curr
		}
	}
	return params
}

// getFieldTag gets the tag of the given field (f)
func fieldTag(f reflect.StructField) string {
	return string(f.Tag)
}

// isNil returns true if the given value (val) is equal to the zero of its type
func isNil(val interface{}) bool {
	if val == nil {
		return true
	}
	return val == reflect.Zero(reflect.TypeOf(val)).Interface()
}

// addParameters adds to uri multiple parameters with the same name
func addParameters(uri *string, paramName string, values []string) {
	for _, val := range values {
		*uri += paramName + "=" + val + "&"
	}
}

// isValid checks if a parameter has a valid value
func isValid(paramName, param string, shouldBe []string) error {
	for _, val := range shouldBe {
		if param == val {
			return nil
		}
	}
	s := fmt.Sprintf("Invalid \"%s\" parameter. Valid values are: %v.", paramName, shouldBe)
	return errors.New(s)
}
