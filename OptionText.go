package main

import (
	"encoding/json"
)

type OptionText struct {
	A string `json:"A"`
	B string `json:"B"`
	C string `json:"C"`
	D string `json:"D"`
}

func NewOptionText() *OptionText {
	return &OptionText{
		A: "Let's go",
		B: "Let's run",
		C: "Let's try",
		D: "Let's love"}
}

func (this *OptionText) GetJsonString() (string, error) {
	jb, err := json.Marshal(this)
	if err != nil {
		return "", err
	} else {
		return string(jb[:]), nil
	}
}
