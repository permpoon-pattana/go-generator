package generator

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
)

func NewJSONGenerator(jsonString string, structName string) *JSONGenerator {
	generator := &JSONGenerator{
		StructName: structName,
		Data:       jsonString,
	}
	return generator.preGenerate()
}

type JSONGenerator struct {
	StructName string
	Data       string
	Fields     map[string]interface{}
}

func (g *JSONGenerator) preGenerate() *JSONGenerator {
	fields := map[string]interface{}{}

	err := json.Unmarshal([]byte(g.Data), &fields)
	if err != nil {
		log.Panicln(err)
	}
	g.Fields = fields
	return g
}

func (g *JSONGenerator) GenerateStruct() *JSONGenerator {
	fmt.Println(">>> start output")
	fmt.Printf("type %v struct {\n", g.StructName)
	for k, v := range g.Fields {
		fmt.Printf("\t%v %v `json:\"%v\"`\n", strings.Title(k), reflect.TypeOf(v).Name(), k)
	}
	fmt.Println("}")
	fmt.Println(">>> end output")
	return g
}
