package main

import "encoding/json"
import "strings"
import "fmt"

func main() {
	var str string = `
	{"Brain": 123456789,"Wu":"junbin"}
	{"Brain": 987654321,"Wu":"junbin"}
	`
	decoder := json.NewDecoder(strings.NewReader(str))
	decoder.Decode(nil)
	r := decoder.Buffered()
	var bs []byte = make([]byte, 100)
	r.Read(bs)
	str2 := strings.TrimSpace(string(bs))
	fmt.Println(str2)
	fmt.Println("---")
}
