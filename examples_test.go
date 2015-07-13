package sorty

import (
	"encoding/json"
	"fmt"
)

func toJson(d interface{}) []byte {
	js, _ := json.Marshal(d)
	return js
}

func ExampleByKeys_Strings() {
	s := NewSorter().ByKeys([]string{
		"+foo",
		"-bar",
	})

	data := []map[string]string{
		{"foo": "abc", "bar": "xyz"},
		{"foo": "xyz", "bar": "abc"},
		{"foo": "def", "bar": "jhi"},
		{"foo": "def", "bar": "def"},
		{"foo": "mno", "bar": "jkl"},
	}

	s.Sort(data)

	fmt.Printf("%s\n", toJson(data))

	// Output:
	// [{"bar":"xyz","foo":"abc"},{"bar":"jhi","foo":"def"},{"bar":"def","foo":"def"},{"bar":"jkl","foo":"mno"},{"bar":"abc","foo":"xyz"}]
}

func ExampleByKeys_interface() {
	s := NewSorter().ByKeys([]string{
		"foo",
		"-bar",
	})

	data := []map[string]interface{}{
		{"foo": "abc", "bar": 890},
		{"foo": "xyz", "bar": 123},
		{"foo": "def", "bar": 456},
		{"foo": "mno", "bar": 789},
		{"foo": "def", "bar": 789},
	}

	s.Sort(data)

	fmt.Printf("%s\n", toJson(data))

	// Output:
	// [{"bar":890,"foo":"abc"},{"bar":789,"foo":"def"},{"bar":456,"foo":"def"},{"bar":789,"foo":"mno"},{"bar":123,"foo":"xyz"}]
}

func ExampleByKeyComps() {
	s := NewSorter().ByKeyComps(KeyComps{
		KeyComp{"foo", Ascending},
		KeyComp{"bar", Descending},
	})

	data := []map[string]string{
		{"foo": "abc", "bar": "xyz"},
		{"foo": "xyz", "bar": "abc"},
		{"foo": "def", "bar": "jhi"},
		{"foo": "def", "bar": "def"},
		{"foo": "mno", "bar": "jkl"},
	}

	s.Sort(data)

	fmt.Printf("%s\n", toJson(data))

	// Output:
	// [{"bar":"xyz","foo":"abc"},{"bar":"jhi","foo":"def"},{"bar":"def","foo":"def"},{"bar":"jkl","foo":"mno"},{"bar":"abc","foo":"xyz"}]
}
