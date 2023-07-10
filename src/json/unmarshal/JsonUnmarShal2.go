package main

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
)

var sampleJson = `[{\"matCode\": \"9998\",\"matDescription\": \"CREAM\"},{\"matCode\": \"1234\",\"matDescription\": \"H/S,LA,Sakura max/128GB,CREAM\"}]`

type MyStruct []struct {
	MatCode        string
	MatDescription string
}

func main() {
	var payload MyStruct
	//
	sampleRegexp := regexp.MustCompile(`\\"`)
	result := sampleRegexp.ReplaceAllString(sampleJson, "\"")
	fmt.Println(string(result))
	rawData := []byte(result)
	err := json.Unmarshal(rawData, &payload)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", payload)
	log.Printf("%v", payload[0].MatCode)
}
