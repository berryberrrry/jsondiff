# JSON Diff in Golang

## How To Use


```

func main() {
    expectedData := []byte(
        `{
            "a":"a",
            "b":"b",
            "c":["a","b","c"]}`)
    actualData := []byte(
        `{
            "a":"a",
            "b":"a",
            "c":["a","b","d"]}`)

    var expected, actual map[string]interface{}
    err := json.Unmarshal(expectedData, &expected)
    if err != nil {
        panic(err)
    }
    err = json.Unmarshal(actualData, &actual)
    if err != nil {
        panic(err)
    }

    differ := jsondiff.New()
    // differ.AddExpectedField("b", 1)

    diffs := differ.Compare(expected, actual)
    fmt.Println(strings.Join(diffs, "\n"))

    //========================================
    // result:
    // map[b]: b != a
    // map[c].array[2]: c != d

    // if you add 
    // differ.AddExpectedField("b",1)
    // result will be :  map[c].array[2]: c != d

}

```