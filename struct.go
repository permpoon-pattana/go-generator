package generator

import (
	"fmt"
	"reflect"
)

var StructDataFormat = map[string]string{
	"int":     "0",
	"float":   "0.0",
	"float32": "0.0",
	"float64": "0.0",
	"Time":    "time.Now()",
	"string":  "\"\"",
	"":        "\"\"",
}

func NewObjectGenerator(data interface{}, prefix string) *ObjectGenerator {
	generator := &ObjectGenerator{
		Prefix: prefix,
		Data:   data,
	}
	return generator.preGenerate()
}

type ObjectGenerator struct {
	Prefix     string
	Data       interface{}
	StructName string
	Types      reflect.Type
	TotalField int
	Result     string
}

func (g *ObjectGenerator) preGenerate() *ObjectGenerator {
	values := reflect.ValueOf(g.Data)
	g.StructName = reflect.TypeOf(g.Data).String()
	g.Types = values.Type()
	g.TotalField = values.NumField()
	return g
}

func (g *ObjectGenerator) GenerateObject() *ObjectGenerator {
	fmt.Println(">>> start output")
	fmt.Printf("example := %v{\n", g.StructName)
	for i := 0; i < g.TotalField; i++ {
		if g.Prefix == "" {
			fmt.Printf("\t%v: %v,\n", g.Types.Field(i).Name, StructDataFormat[g.Types.Field(i).Type.Name()])
			continue
		}

		fmt.Printf("\t%v: %v.,\n", g.Types.Field(i).Name, g.Prefix)
	}
	fmt.Println("}")
	fmt.Println(">>> end output")
	return g
}
