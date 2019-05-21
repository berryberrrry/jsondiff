/*
 * @Author: berryberry
 * @since: 2019-05-21 14:42:12
 * @lastTime: 2019-05-21 19:23:06
 * @LastAuthor: Do not edit
 */
package jsondiff

import (
	"fmt"
	"reflect"
	"strings"
)

type Config struct {
	MaxDiff          int
	MaxDeep          int
	SortedSlice      bool
	HasExceptedField bool
	ExceptedField    []Field
}

type Field struct {
	Key  string
	Deep int
}

type Differ struct {
	Conf Config
	diff []string
	buff []string
}

func New() *Differ {
	return &Differ{
		Conf: Config{
			MaxDiff:          10,
			MaxDeep:          10,
			SortedSlice:      true,
			HasExceptedField: false,
			ExceptedField:    nil,
		},
		diff: []string{},
		buff: []string{},
	}
}

func Diff(expected, actual interface{}) []string {

	differ := &Differ{
		Conf: Config{
			MaxDiff:          10,
			MaxDeep:          10,
			SortedSlice:      true,
			HasExceptedField: false,
			ExceptedField:    nil,
		},
		diff: []string{},
		buff: []string{},
	}

	if expected == nil && actual == nil {
		return nil
	} else if expected == nil && actual != nil {
		differ.saveDiff("<nil>", actual)
		return differ.diff
	} else if expected != nil && actual == nil {
		differ.saveDiff(expected, "<nil>")
		return differ.diff
	}

	expectedVal := reflect.ValueOf(expected)
	actualVal := reflect.ValueOf(actual)

	differ.compare(expectedVal, actualVal, 0)
	if len(differ.diff) > 0 {
		return differ.diff
	}

	return nil
}

func (d *Differ) Compare(expected, actual interface{}, deep int) {
	// Check if one value is nil, e.g. T{x: *X} and T.x is nil
	if !expected.IsValid() || !actual.IsValid() {
		if expected.IsValid() && !actual.IsValid() {
			d.saveDiff(expected.Type(), "<nil>")
		} else if !expected.IsValid() && acutal.IsValid() {
			d.saveDiff("<nil>", actual.Type())
		}
		return
	}
}

func (d *Differ) saveDiff(expectedVal, actualVal interface{}) {
	if len(d.buff) > 0 {
		path := strings.Join(d.buff, ".")
		d.diff = append(d.diff, fmt.Sprintf("%s: %v != %v", path, expectedVal, actualVal))
	} else {
		d.diff = append(d.diff, fmt.Sprintf("%v != %v", expectedVal, actualVal))
	}
}
