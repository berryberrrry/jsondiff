/*
 * @Author: berryberry
 * @since: 2019-05-21 20:53:59
 * @lastTime: 2019-05-21 21:25:44
 * @LastAuthor: Do not edit
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/berryberrrry/jsondiff"
)

func main() {
	expectedData, err := ioutil.ReadFile("/home/berryberry/workspaces/git_workspaces/goProjects/src/shannonpdf/experiments/test_aws_pdf/招股书P020190312667782509416.pdf-1.json")
	if err != nil {
		panic(err)
	}
	actualData, err := ioutil.ReadFile("/home/berryberry/workspaces/git_workspaces/goProjects/src/shannonpdf/experiments/test_aws_pdf/招股书P020190312667782509416.pdf-2.json")
	if err != nil {
		panic(err)
	}

	var expected, actual map[string]interface{}
	err = json.Unmarshal(expectedData, &expected)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(actualData, &actual)
	if err != nil {
		panic(err)
	}

	differ := jsondiff.New()
	differ.AddExpectedField("b", 1)
	differ.AddExpectedField("textLen", -1)
	differ.AddExpectedField("textOffset", -1)
	differ.Conf.MaxDiff = 10000
	differ.Conf.MaxDeep = 15
	// differ.AddExpectedField("c", 1)
	diffs := differ.Compare(expected, actual)
	fmt.Println(strings.Join(diffs, "\n"))
}
