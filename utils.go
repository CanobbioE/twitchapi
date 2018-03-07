package gwat

// TODO: use net/url library
import (
	"reflect"
)

// parseInput transform an input struct into a map
func parseInput(input interface{}) map[string]interface{} {
	s := reflect.ValueOf(input)
	params := make(map[string]interface{})

	for i := 0; i < s.NumField(); i++ {
		curr := s.Field(i).Interface()
		tag := getFieldTag(reflect.TypeOf(input).Elem().FieldByIndex([]int{i}))
		if !isNil(curr) {
			params[tag] = curr
		}
	}
	return params
}

// getFieldTag gets the tag of the given field (f)
func getFieldTag(f reflect.StructField) string {
	return string(f.Tag)
}

// isNil returns true if the given value (val) is equal to the zero of its type
func isNil(val interface{}) bool {
	return val == reflect.Zero(reflect.TypeOf(val)).Interface()
}

// addParameters adds to uri multiple parameters with the same name
func addParameters(uri *string, paramName string, values []string) {
	for _, val := range values {
		*uri += val
		*uri += "&" + paramName + "="
	}
}
