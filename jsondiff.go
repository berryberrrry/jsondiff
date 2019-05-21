/*
 * @Author: berryberry
 * @since: 2019-05-21 14:42:12
 * @lastTime: 2019-05-21 21:46:57
 * @LastAuthor: Do not edit
 */
package jsondiff

import (
	"fmt"
	"reflect"
	"strings"
)

type Config struct {
	MaxDiff     int
	MaxDeep     int
	SortedSlice bool
	// HasExceptedField bool
	ExceptedFields map[int]map[string]int
}

type Differ struct {
	Conf Config
	diff []string
	buff []string
}

func New() *Differ {
	exceptedFields := make(map[int]map[string]int, 0)
	return &Differ{
		Conf: Config{
			MaxDiff:     10,
			MaxDeep:     10,
			SortedSlice: true,
			// HasExceptedField: false,
			ExceptedFields: exceptedFields,
		},
		diff: []string{},
		buff: []string{},
	}
}

func (d *Differ) Compare(expected, actual map[string]interface{}) []string {
	d.compareMap(expected, actual, 1)
	return d.diff
}

func (d *Differ) compareMap(expected, actual map[string]interface{}, deep int) {
	exceptedKeysMap := d.Conf.ExceptedFields[deep]
	globalExceptedKeysMap := d.Conf.ExceptedFields[-1]
	for expectedKey, expectedVal := range expected {
		if _, ok := globalExceptedKeysMap[expectedKey]; ok {
			continue
		}
		if _, ok := exceptedKeysMap[expectedKey]; ok {
			continue
		}
		actualVal := actual[expectedKey]
		if actualVal != nil {
			d.push(fmt.Sprintf("map[%s]", expectedKey))
			d.compareVal(expectedVal, actualVal, deep+1)
			d.pop()
		} else {
			d.saveDiff(expectedVal, "<nil>")
		}
	}
	if len(expected) == len(actual) {
		return
	}

	for actualKey, actualVal := range actual {
		if _, ok := globalExceptedKeysMap[actualKey]; ok {
			continue
		}
		if _, ok := exceptedKeysMap[actualKey]; ok {
			continue
		}

		expectedVal := expected[actualKey]
		if expectedVal == nil {
			d.saveDiff("<nil>", actualVal)
		}
	}
}

func (d *Differ) compareArray(expected, actual []interface{}, deep int) {
	expectedLen := len(expected)
	actualLen := len(actual)
	maxLen := expectedLen
	if actualLen > maxLen {
		maxLen = actualLen
	}

	for i := 0; i < maxLen; i++ {
		d.push(fmt.Sprintf("array[%d]", i))
		if i < expectedLen && i < actualLen {
			d.compareVal(expected[i], actual[i], deep+1)

		} else if i < expectedLen {
			d.saveDiff(expected[i], "<nil>")
		} else {
			d.saveDiff("<nil>", actual[i])
		}
		d.pop()
	}
}

func (d *Differ) compareVal(expectedVal, actualVal interface{}, deep int) {
	if deep > d.Conf.MaxDeep {
		return
	}
	expectedType := reflect.TypeOf(expectedVal)
	actualType := reflect.TypeOf(actualVal)
	if expectedType != actualType {
		d.saveDiff(expectedType, actualType)
		return
	}

	switch expectedVal.(type) {
	case map[string]interface{}:
		d.compareMap(expectedVal.(map[string]interface{}), actualVal.(map[string]interface{}), deep)
	case []interface{}:
		d.compareArray(expectedVal.([]interface{}), actualVal.([]interface{}), deep)
	default:
		if !reflect.DeepEqual(expectedVal, actualVal) {
			d.saveDiff(expectedVal, actualVal)
		}
	}

}

func (d *Differ) saveDiff(expectedVal, actualVal interface{}) {
	if len(d.diff) >= d.Conf.MaxDiff {
		return
	}
	if len(d.buff) > 0 {
		path := strings.Join(d.buff, ".")
		d.diff = append(d.diff, fmt.Sprintf("%s: %v != %v", path, expectedVal, actualVal))
	} else {
		d.diff = append(d.diff, fmt.Sprintf("%v != %v", expectedVal, actualVal))
	}
}

func (d *Differ) push(str string) {
	d.buff = append(d.buff, str)
}

func (d *Differ) pop() {
	if len(d.buff) > 0 {
		d.buff = d.buff[0 : len(d.buff)-1]
	}
}

func (d *Differ) AddExpectedField(key string, deep int) {
	keysMap := d.Conf.ExceptedFields[deep]
	if keysMap == nil {
		newMap := make(map[string]int, 0)
		newMap[key] = 1
		d.Conf.ExceptedFields[deep] = newMap
	} else {
		keysMap[key] = 1
		d.Conf.ExceptedFields[deep] = keysMap
	}
}

// func (d *Differ) RemoveExectedField(key string, deep int) {
// 	//TODO
// }
