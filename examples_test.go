package sorty

import (
	"fmt"
)

func ExampleByKeys() {
	s := NewSorter().ByKeys([]string{
		"+foo",
		"-bar",
	})

	data := []map[string]string{
		{
			"foo": "abc",
			"bar": "xyz",
		},
		{
			"foo": "xyz",
			"bar": "abc",
		},
		{
			"foo": "def",
			"bar": "jhi",
		},
		{
			"foo": "def",
			"bar": "def",
		},
		{
			"foo": "mno",
			"bar": "jkl",
		},
	}

	s.Sort(data)

	foos := make([]string, 0)
	bars := make([]string, 0)
	for _, d := range data {
		foos = append(foos, d["foo"])
		bars = append(bars, d["bar"])
	}
	fmt.Printf("foos: %v\n", foos)
	fmt.Printf("bars: %v\n", bars)

	// Output:
	// foos: [abc def def mno xyz]
	// bars: [xyz jhi def jkl abc]
}

func ExampleByKeyComps() {
	s := NewSorter().ByKeyComps(KeyComps{
		KeyComp{"foo", Ascending},
		KeyComp{"bar", Descending},
	})

	data := []map[string]string{
		{
			"foo": "abc",
			"bar": "xyz",
		},
		{
			"foo": "xyz",
			"bar": "abc",
		},
		{
			"foo": "def",
			"bar": "jhi",
		},
		{
			"foo": "def",
			"bar": "def",
		},
		{
			"foo": "mno",
			"bar": "jkl",
		},
	}

	s.Sort(data)

	foos := make([]string, 0)
	bars := make([]string, 0)
	for _, d := range data {
		foos = append(foos, d["foo"])
		bars = append(bars, d["bar"])
	}
	fmt.Printf("foos: %v\n", foos)
	fmt.Printf("bars: %v\n", bars)

	// Output:
	// foos: [abc def def mno xyz]
	// bars: [xyz jhi def jkl abc]
}
