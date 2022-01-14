package Utilitys

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Description struct {
	TimeStamp  string `json:"TimeStamp"`
	Action     string `json:"Action"`
	MethodName string `json:"MethodName"`
	LineNumber int    `json:"LineNumber"`
	Priority   int    `json:"Priority"`
}
type Message struct {
	GoMessage  string `json:"GoMessage"`
	AppMessage string `json:"AppMessage"`
	Suggestion string `json:"Suggestion"`
}
type Value struct {
	Description `json:"Description"`
	Message     `json:"Message"`
}

type Exceptions struct {
	Key   int `json:"Key"`
	Value `json:"Value"`
}

func RaiseError() *[]Exceptions {
	byte, err := ioutil.ReadFile("D://Projects//Go-lang//CurrencyServices//ApplicationFiles//Errors.json")
	if err != nil {
		fmt.Println("Can't open Errors file!", err)
		os.Exit(0)
	}

	var jsonMap []Exceptions
	err = json.Unmarshal(byte, &jsonMap)
	if err != nil {
		fmt.Println("Can't Unmarshal Errors file!", err)
		os.Exit(0)
	}
	return &jsonMap
}

func SelectException(Code int, Array *[]Exceptions) *Exceptions {
	for _, v := range *Array {
		if Code == v.Key {
			return &v
		}
	}

	return &Exceptions{
		0,
		Value{
			Description{
				Priority: 100, LineNumber: -1, MethodName: "main", TimeStamp: "", Action: "Bepar rosh"},
			Message{
				GoMessage: "Redam", AppMessage: "Ridi", Suggestion: "Find some water or tissue. "},
		},
	}
}
