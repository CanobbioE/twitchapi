package twitchapi

import (
	"reflect"
	"testing"
)

type TestStruct struct {
	Field string `field_name`
	Value int    `field_value`
}

func TestFieldTag(t *testing.T) {
	ts := TestStruct{}
	var tests = []struct {
		input reflect.StructField
		want  string
	}{
		{reflect.TypeOf(ts).FieldByIndex([]int{0}), "field_name"},
	}

	for _, test := range tests {
		if got := fieldTag(test.input); got != test.want {
			t.Errorf("fieldTag(%v) = %s (expected %s)", test.input, got, test.want)
		}
	}
}

func TestParseInput(t *testing.T) {

	var tests = []struct {
		input interface{}
		want  map[string]interface{}
	}{
		{
			TestStruct{
				Field: "name",
				Value: 42,
			},
			map[string]interface{}{
				"field_name":  "name",
				"field_value": 42,
			},
		},
	} // tests

	for _, test := range tests {
		got := parseInput(test.input)
		for k, v := range test.want {
			val, ok := got[k]
			if !ok || val != v {
				t.Errorf("parseInput(): Key %s has value %v (expected %v)", k, val, v)
			}
		}
	}
}

func TestIsNil(t *testing.T) {
	var tests = []struct {
		input interface{}
		want  bool
	}{
		{0, true},
		{1, false},
		{"", true},
		{"hello", false},
		{nil, true},
		{false, true},
		{true, false},
		{[]string{}, true},
		{[]string{"hi"}, false},
	}
	for _, test := range tests {
		if got := isNil(test.input); got != test.want {
			t.Errorf("isNil(%v) = %b (expected %b)", test.input, got, test.want)
		}
	}
}

func TestIsValid(t *testing.T) {
	var tests = []struct {
		input1  string
		input2  string
		input3  []string
		wantErr bool
	}{
		{"name", "test", []string{"this", "is", "good", "test"}, false},
		{"name", "test", []string{"not", "good"}, true},
		{"name", "", []string{}, true},
	}

	for _, test := range tests {
		err := isValid(test.input1, test.input2, test.input3)
		if test.wantErr && err == nil {
			t.Errorf("isValid(): Expecting error, got nil")
		}
		if !test.wantErr && err != nil {
			t.Errorf("isValid(): Expecting nil, got err")
		}
	}
}

func TestMakeUri(t *testing.T) {
	var tests = []struct {
		ep   string
		qp   interface{}
		want string
	}{
		{
			BaseURL + "games",
			GameQueryParameters{
				Names: []string{"league of legends"},
			},
			"https://api.twitch.tv/helix/games?name=league+of+legends",
		},
	}

	for _, test := range tests {
		got := makeUri(test.ep, test.qp)
		if got != test.want {
			t.Errorf("makeUri = %s (expected %s)", got, test.want)
		}
	}
}
