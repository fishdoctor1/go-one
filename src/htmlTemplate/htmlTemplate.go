package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

type Todo struct {
	UserId    int    `json:"userId"` //for beauty marshal
	Id        int    `json:"id"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed,omitempty"` //ถ้า api ไม่ส่ง completed มา ให้ใช้ pointer มันจะไม่ขึ้นตาม api
}

func main() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		return
	}
	todoDecoder := json.NewDecoder(resp.Body)
	fmt.Println(resp.Body)
	// todoDecoder := json.NewDecoder(bytes.NewReader(resp)) //for from file
	todos := []Todo{}
	todoDecoder.Decode(&todos)
	todoEncoder := json.NewEncoder(os.Stdout)
	todoEncoder.Encode(todos)
	// fmt.Println(todos)
	indexTemplateBytes, err := ioutil.ReadFile("E:/Learning/Golang/go-udemy/src/htmlTemplate/index.html")
	if err != nil {
		return
	}
	indexTemplate := template.Must(template.New("index").Parse(string(indexTemplateBytes)))
	indexTemplate.Execute(os.Stdout, todos)
}

func fromfile() {
	resp, err := ioutil.ReadFile("E:/Learning/Golang/go-udemy/src/htmlTemplate/htmlTemplate.go")
	if err != nil {
		return
	}
	// todoDecoder := json.NewDecoder(resp.Body)
	todoDecoder := json.NewDecoder(bytes.NewReader(resp)) //for from file
	todos := []Todo{}
	todoDecoder.Decode(&todos)

	indexTemplateBytes, err := ioutil.ReadFile("E:/Learning/Golang/go-udemy/src/htmlTemplate/index.html")
	if err != nil {
		return
	}
	indexTemplate := template.Must(template.New("index").Parse(string(indexTemplateBytes)))
	indexTemplate.Execute(os.Stdout, todos)
}
