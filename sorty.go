/*
 *
 *  Copyright 2015 Netflix, Inc.
 *
 *     Licensed under the Apache License, Version 2.0 (the "License");
 *     you may not use this file except in compliance with the License.
 *     You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 *     Unless required by applicable law or agreed to in writing, software
 *     distributed under the License is distributed on an "AS IS" BASIS,
 *     WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *     See the License for the specific language governing permissions and
 *     limitations under the License.
 *
 */
package sorty

import (
	"reflect"
	"sort"
)

type sorter struct {
	data  interface{}
	order KeyComps
}

func NewSorter() *sorter {
	return &sorter{}
}

func (s *sorter) ByKeys(order []string) *sorter {
	keyComps := make(KeyComps,0)
	for _, key := range order {
		switch key[0] {
		case '-':
			keyComps = append(keyComps, KeyComp{key[1:],Descending})
		case '+':
			keyComps = append(keyComps, KeyComp{key[1:],Ascending})
		default:
			keyComps = append(keyComps, KeyComp{key,Ascending})
		}
	}
	return s.ByKeyComps(keyComps)
}

type KeyComp struct {
	Name string
	Comp func(interface{}, interface{}) CompareResult
}

type KeyComps []KeyComp

func (s *sorter) ByKeyComps(keyComps KeyComps) *sorter {
	s.order = keyComps
	return s
}

func (s *sorter) Sort(data interface{}) {
	s.data = data
	sort.Sort(s)
}

func (s *sorter) Len() int {
	return reflect.ValueOf(s.data).Len()
}

func (s *sorter) Swap(i, j int) {
	if i > j {
		i, j = j, i
	}
	arr := reflect.ValueOf(s.data)
	copy := reflect.MakeSlice(arr.Type(), 2, 2)
	reflect.Copy(copy, arr.Slice(i, j+1))
	arr.Index(i).Set(copy.Index(1))
	arr.Index(j).Set(copy.Index(0))
}

func (s *sorter) Less(i, j int) bool {
	arr := reflect.ValueOf(s.data)
	a := arr.Index(i)
	b := arr.Index(j)
	for i := 0; i < len(s.order); i += 1 {
		keyComp := s.order[i]
		af := a.MapIndex(reflect.ValueOf(keyComp.Name)).Interface()
		bf := b.MapIndex(reflect.ValueOf(keyComp.Name)).Interface()

		switch keyComp.Comp(af, bf) {
		case LESSER:
			return true
		case GREATER:
			return false
		}
	}
	return true
}

type CompareResult int8

const (
	LESSER CompareResult = -1 + iota
	EQUAL
	GREATER
)

func Ascending(a, b interface{}) CompareResult {
	switch Descending(a,b) {
	case LESSER:
		return GREATER
	case GREATER:
		return LESSER
	default:
		return EQUAL
	}
}

func Descending(a, b interface{}) CompareResult {
	if a == b {
		return EQUAL
	}
	switch a.(type) {
	case string:
		return lg(a.(string) > b.(string))
	case int:
		return lg(a.(int) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(int))
	case int8:
		return lg(a.(int8) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(int8))
	case int32:
		return lg(a.(int32) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(int32))
	case int64:
		return lg(a.(int64) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(int64))
	case uint:
		return lg(a.(uint) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(uint))
	case uint8:
		return lg(a.(uint8) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(uint8))
	case uint32:
		return lg(a.(uint32) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(uint32))
	case uint64:
		return lg(a.(uint64) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(uint64))
	case float32:
		return lg(a.(float32) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(float32))
	case float64:
		return lg(a.(float64) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(float64))
	}

	return GREATER
}

func lg(b bool) CompareResult {
	if b {
		return LESSER
	}
	return GREATER
}
