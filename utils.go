package twitchapi

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

// parseInput transform an input struct into a map
func parseInput(input interface{}) map[string]interface{} {
	s := reflect.ValueOf(input)
	params := make(map[string]interface{})

	for i := 0; i < s.NumField(); i++ {
		curr := s.Field(i).Interface()
		tag := fieldTag(reflect.TypeOf(input).FieldByIndex([]int{i}))
		if !isEmpty(curr) {
			params[tag] = curr
		}
	}
	return params
}

// getFieldTag gets the tag of the given field (f)
func fieldTag(f reflect.StructField) string {
	return string(f.Tag)
}

// isEmpty returns true if the given value (val) is equal to the zero of its type
func isEmpty(val interface{}) bool {
	if val == nil {
		return true
	}
	// Slice supports
	switch t := val.(type) {
	case []int:
		return len(t) == 0
	case []bool:
		return len(t) == 0
	case []string:
		return len(t) == 0
	case []float64:
		return len(t) == 0
	}
	return val == reflect.Zero(reflect.TypeOf(val)).Interface()
}

// isValid checks if a parameter has a valid value
func isValid(paramName, param interface{}, shouldBe ...interface{}) error {
	for _, val := range shouldBe {
		if param == val {
			return nil
		}
	}
	s := fmt.Sprintf("Invalid \"%s\" parameter. Valid values are: %v.", paramName, shouldBe)
	return errors.New(s)
}

// makeUri creates a uri and returns it as string
func makeUri(location string, qp interface{}, notEmptyParams ...interface{}) string {

	uri := &url.URL{}

	uri, err := url.Parse(location)
	if err != nil {
		panic(err)
	}

	params := parseInput(qp)
	values := url.Values{}
	for k, v := range params {
		switch t := v.(type) {
		case []string:
			for i := range t {
				values.Add(k, t[i])
			}
		case []int:
			for i := range t {
				values.Add(k, strconv.Itoa(t[i]))
			}
		case int:
			values.Add(k, strconv.Itoa(t))
		case string:
			values.Add(k, t)
		}
	}

	uri.RawQuery = values.Encode()

	return uri.String()
}

// checkRequiredFields verifies that all the required fields for the apiCallName
// are not empty
func checkRequiredFields(apiCallName string, params ...interface{}) error {
	for _, p := range params {
		if isEmpty(p) {
			return fmt.Errorf("%s: a required parameter for the request is missing", apiCallName)
		}
	}
	return nil
}

func setDefaultValueIf(condition bool, param, defaultVal interface{}) interface{} {
	if condition {
		return defaultVal
	}
	return param
}
