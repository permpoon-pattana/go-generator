# Generator

## Object generator
### From struct

Example to use
```golang
package main

type Example struct {
	A string
	B string
}

func main() {
	example := Example{}
	generator.NewObjectGenerator(example, "").GenerateObject()
}
```

```bash
$ go run main.go

>>> start output
example := main.Example{
        A: "",
        B: "",
}
>>> end output
```


```golang
package main

func main() {
	jsonString := `
	{
        "name": "xxx",
        "age": 5
    }
	`
	generator.NewJSONGenerator(jsonString, "Test").GenerateStruct()
}

```

```bash
$ go run main.go

>>> start output
type Test struct {
	Name string  `json:"name"`
	Age  float64 `json:"age"`
}
>>> end output
```