sorty
=========

[![Build Status](https://travis-ci.org/coryb/sorty.svg?branch=master)](https://travis-ci.org/coryb/sorty)
[![Coverage Status](https://coveralls.io/repos/coryb/sorty/badge.svg?branch=master)](https://coveralls.io/r/coryb/sorty?branch=master)
[![GoDoc](https://godoc.org/github.com/coryb/sorty?status.png)](https://godoc.org/github.com/coryb/sorty)

**sorty** golang library to allow sorting slices of maps easy.

## Sample

```go
package main

import (
	"fmt"
	"github.com/coryb/sorty"
	"os"
)

func main() {
    // sort maps first in ascending order on the "foo" key
    // then in decending order on the "bar" key
	s := NewSorter().ByKeys([]string{
		"+foo",
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

    // Data will now look like:
    // []map[string]interface{}{
	//	{"foo": "abc", "bar": 890},
	//	{"foo": "def", "bar": 789},
	//	{"foo": "def", "bar": 456},
	//	{"foo": "mno", "bar": 789},
	//	{"foo": "xyz", "bar": 123},
	// }
}
```

## Why?

I found myself writing several tools that needed to pull large json documents from rest apis.  The rest apis did not always return 
data in a sorted order or in a sorted order that was not what I needed to process/display them in.  I found it tedious to write the sort.Interface
for sorting on multiple keys so I created an abstraction.  I dont expect it do be very performant, but I am typically sorting dozens of items instead of 
millions.

## Installation

```go
import "github.com/coryb/sorty"
```

To install sorty according to your `$GOPATH`:

```console
$ go get github.com/coryb/sorty
```

## API Documentation

See the Documentation published on godoc [here](https://godoc.org/github.com/coryb/sorty).
