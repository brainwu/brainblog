package main

import "encoding/json"
import "strings"
import "os"
import "io"
import "fmt"
import (
	"unicode"
)

func main() {
	encoder := json.NewEncoder(os.Stdout)
	m := make(map[string]string)
	m["wu"] = "junbin"
	m["brain"]="wu"
	encoder.Encode(m)
	var str string = `
	{"Brain":12322312,"Wu":"junbin"}
	{"Brain":6667787,"Wu":"junbin"}
	`
	decoder := json.NewDecoder(strings.NewReader(str))
	type Person struct {
		Brain int
		Wu string
	}
	decoder.UseNumber()
	for {
		var p interface {}
		if err := decoder.Decode(&p); err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Printf("decoded to %#v\n", p)
		m := p.(map[string]interface{})
		fmt.Println(m["Brain"])
//		fmt.Printf("%s : %s", p.Brain, p.Wu)
		fmt.Println()
	}
}
