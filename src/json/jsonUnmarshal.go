package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Todo []struct {
	UserId    int    `json:"userId"` //for beauty marshal
	Id        int    `json:"id"`
	Title     string `json:"title,omitempty"`
	Completed *bool  `json:"completed,omitempty"` //ถ้า api ไม่ส่ง completed มา ให้ใช้ pointer มันจะไม่ขึ้นตาม api
}

type User []struct {
	ID       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}

// var data = `[{
// 	"userId":678	,
// 	"id":555,
// 	"title":"ASD",
// 	"completed":false
// },
// {
// 	"userId":987	,
// 	"id":666,
// 	"title":"DSA",
// 	"completed":false
// }]`

func main() {
	// format1()
	// format2()
	testUnmarshalArrayObject()
}

func format1() {
	dataStruct := User{}
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}
	v := &dataStruct
	json.Unmarshal(body, v)

	data, err := json.Marshal(dataStruct)
	// data, err := json.MarshalIndent(dataStruct,"","    ") // for beauty fmt
	if err != nil {
		return
	}
	fmt.Println(string(data))
	// dataStruct[0].Completed = true
}

func format2() {
	dataStruct := User{}
	dataStruct2 := User{}
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		return
	}
	jsonDecoder := json.NewDecoder(resp.Body)
	jsonDecoder.Decode(&dataStruct)
	jsonDecoder.Decode(&dataStruct2)
	fmt.Println("----------------------END DECODE")
	resp.Body.Close()
	fmt.Println()
	dataStruct[0].Name = "kanta"

	jsonEncoder := json.NewEncoder(os.Stdout)
	jsonEncoder.Encode(dataStruct) //print out
	fmt.Println("111111111111111111111")
	jsonEncoder.Encode(dataStruct2) //empty array

}
