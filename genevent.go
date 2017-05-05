package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"text/template"
)

type Def struct {
	Name     string    `xml:"name,attr"`
	Comment  string    `xml:"comment,attr"`
	Value  string `xml:"value,attr"`
}

type Enum struct {
	Name     string    `xml:"name,attr"`
	Comment  string    `xml:"comment,attr"`
	File  string    `xml:"file,attr"`
	Defs  []Def    `xml:"def"`
}

type Result struct {
	XMLName xml.Name `xml:"result"`
	Enums []Enum `xml:"enum"`
}

var result = Result{}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("usage:SrcFileName templateFileName")
		return
	}

	srcFile := os.Args[1]
	tplFile := os.Args[2]

	if !parse(srcFile) {
		return
	}

	for _,e := range result.Enums {
	
		if !output(e, tplFile) {
			return
		}

	}
	
	fmt.Println("OK")
}

func parse(name string) bool {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return false
	}

	var buffer [1024 * 1024]byte
	n, rerr := file.Read(buffer[0:])
	if rerr != nil {
		fmt.Println(rerr)
		return false
	}

	err = xml.Unmarshal(buffer[0:n], &result)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func genlist(n string) []string {
	num, _ := strconv.Atoi(n)
	ret := make([]string, num)
	for i := 0; i < num; i++ {
		ret[i] = strconv.Itoa(i)
	}
	return ret
}

func output(e Enum, src string) bool {

	file, err := os.Create(e.File)
	if err != nil {
		fmt.Println(err)
		return false
	}

	t := template.New("text")
	if err != nil {
		fmt.Println(err)
		return false
	}

	t = t.Funcs(template.FuncMap{"genlist": genlist})

	srcfile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		return false
	}

	var buffer [1024 * 1024]byte
	n, rerr := srcfile.Read(buffer[0:])
	if rerr != nil {
		fmt.Println(rerr)
		return false
	}

	t, err = t.Parse(string(buffer[0:n]))
	if err != nil {
		fmt.Println(err)
		return false
	}

	err = t.Execute(file, e)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
