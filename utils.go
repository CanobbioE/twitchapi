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
func makeUri(location string, qp interface{}) string {

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

// checkRequiredFields verifies that all/any of the required fields
// (based on the specified logic) for the apiCallName are not empty
// logic = "all" means ALL parameters must be NOT empty
// logic = "any" means AT LEAST ONE parameter must be NOT empty
func checkRequiredFields(apiCallName, logic string, params ...interface{}) error {
	for _, p := range params {
		if isEmpty(p) && logic == "all" {
			return fmt.Errorf("%s: a required parameter for the request is missing", apiCallName)
		} else if !isEmpty(p) && logic == "any" {
			return nil
		}
	}
	return nil
}

// setDefaultValueIf returns the default value for a param if a condition is met.
func setDefaultValueIf(condition bool, param, defaultVal interface{}) interface{} {
	if condition {
		return defaultVal
	}
	return param
}

// makeRange creates an int slice with all the integers from min to max (inclusive)
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := 0; min <= max; min++ {
		a[i] = min
	}
	return a
}
