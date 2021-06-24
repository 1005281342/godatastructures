package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	pkgName      = "package heap"
	baseFilePath = "" // you base file
)

func main() {
	f, err := os.Open(baseFilePath)
	if err != nil {
		panic(err)
	}
	fileData, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	w := new(bytes.Buffer)

	startPos := strings.Index(string(fileData), pkgName)
	w.WriteString(string(fileData)[startPos : startPos+len(pkgName)])
	ts := []string{"Float32", "Float64", "Int32", "Int16", "Uint64", "Uint32", "Uint16",
		"Uint", "Int"} // all types need to be converted

	for _, upper := range ts {
		lower := strings.ToLower(upper)
		data := string(fileData)
		// Remove header.
		data = data[startPos+len(pkgName):]
		// Remove the special case.
		//data = strings.Replace(data, lengthFunction, "", -1)
		// Common cases.
		data = strings.Replace(data, "int64", lower, -1)
		data = strings.Replace(data, "Int64", upper, -1)
		//if inSlice(lowerSlice(ts), lower) {
		//	data = strings.Replace(data, "length "+lower, "length int64", 1)
		//}
		// Add the special case.
		//data = data + strings.Replace(lengthFunction, "Int64Set", upper+"Set", 1)
		w.WriteString(data)
		w.WriteString("\r\n")
	}
	var t = w.String()
	fmt.Println(t)
	out, err := format.Source(w.Bytes())
	if err != nil {
		log.Print(err)
		//panic(err)
	}
	if err := ioutil.WriteFile("types.go", out, 0660); err != nil {
		//panic(err)
		log.Print(err)
	}
	log.Printf("只用于辅助生成代码，需要手动调整一些错误语法")
}

func lowerSlice(s []string) []string {
	n := make([]string, len(s))
	for i, v := range s {
		n[i] = strings.ToLower(v)
	}
	return n
}

func inSlice(s []string, val string) bool {
	for _, v := range s {
		if v == val {
			return true
		}
	}
	return false
}
