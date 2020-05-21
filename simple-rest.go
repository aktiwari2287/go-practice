package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Emp struct{
	Name string `json:"Name"`
	Address string  `json:"Address"`
}

type Foo struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}

func main() {
	e,_:=json.Marshal(Emp{Name:"Anand", Address: "Ranchi"})
	names:=map[string]int{
		"Anand":1,
		"Pooja":2}

	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		io.WriteString(w, string(e))
	})
	http.HandleFunc("/map", func(writer http.ResponseWriter, request *http.Request) {
		jsonNames,_:=json.Marshal(names)
		io.WriteString(writer, string(jsonNames))
	})

	http.HandleFunc("/addNames", func(w http.ResponseWriter, r *http.Request) {
		b,_:=ioutil.ReadAll(r.Body)
		type EmpName struct {
			Name string
			Id int
		}
		var e EmpName
		_=json.Unmarshal(b, &e)
		names[e.Name]=e.Id
	})


	http.ListenAndServe(":9090", nil)
}