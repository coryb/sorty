package sorty

import (
	"math/rand"
	"testing"
	"time"
)

func TestLessError(t *testing.T) {
	// Less requires list of maps
	stuff := make([]interface{}, 0)
	stuff = append(stuff, map[string]string{"key": "val"})
	stuff = append(stuff, []int{42})

	// fail unless we get a panic
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	NewSorter().ByKeys([]string{"key"}).Sort(stuff)

	stuff[0], stuff[1] = stuff[1], stuff[0]
	NewSorter().ByKeys([]string{"key"}).Sort(stuff)
}

func TestLessError_reversedItems(t *testing.T) {
	// Less requires list of maps
	stuff := make([]interface{}, 0)
	stuff = append(stuff, []int{42})
	stuff = append(stuff, map[string]string{"key": "val"})

	// fail unless we get a panic
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	NewSorter().ByKeys([]string{"key"}).Sort(stuff)
}

func TestDups(t *testing.T) {
	// Less requires list of maps
	stuff := make([]interface{}, 0)
	stuff = append(stuff, map[string]string{"key": "val"})
	stuff = append(stuff, map[string]string{"key": "val"})

	NewSorter().ByKeys([]string{"key"}).Sort(stuff)

	if len(stuff) != 2 {
		t.Fail()
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func TestTypes(t *testing.T) {
	rand.Seed(time.Now().Unix())
	stuff := make([]map[string]interface{}, 0)
	for i := 0; i < 1000; i++ {
		stuff = append(stuff, map[string]interface{}{
			"string":  randString(25),
			"int":     rand.Int(),
			"int8":    int8(rand.Int()),
			"int16":   int16(rand.Int()),
			"int32":   int32(rand.Int()),
			"int64":   int64(rand.Int()),
			"uint":    uint(rand.Int()),
			"uint8":   uint8(rand.Int()),
			"uint16":  uint16(rand.Int()),
			"uint32":  uint32(rand.Int()),
			"uint64":  uint64(rand.Int()),
			"float32": rand.Float32(),
			"float64": rand.Float64(),
		})
	}

	NewSorter().ByKeys([]string{
		"string",
		"-int",
		"int8",
		"-int16",
		"int32",
		"-int64",
		"uint",
		"-uint8",
		"uint16",
		"-uint32",
		"uint64",
		"-float32",
		"float64",
	}).Sort(stuff)

	if len(stuff) != 1000 {
		t.Fail()
	}
}

func TestUnknownTypes(t *testing.T) {
	// set up unsortable list, we dont know how to sort bool
	stuff := make([]interface{}, 0)
	stuff = append(stuff, map[string]bool{"key": true})
	stuff = append(stuff, map[string]bool{"key": false})

	// fail unless we get a panic
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	NewSorter().ByKeys([]string{"key"}).Sort(stuff)
}
