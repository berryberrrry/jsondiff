# JSON Diff in Golang

## How To Use

```
func main() {
	expectedData, err := ioutil.ReadFile("expected.json")
	if err != nil {
		panic(err)
	}
	actualData, err := ioutil.ReadFile("actual.json")
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

```