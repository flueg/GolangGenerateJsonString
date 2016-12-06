package main

import (
	"encoding/json"
)

type OptionDiff struct {
	Img string `json:"img"`
}

func NewOptionDiff() *OptionDiff {
	return &OptionDiff{Img: "http:flueg.liu.com/abc/avatar.jpg"}
}

func (this *OptionDiff) GetJsonString() (string, error) {
	jb, err := json.Marshal(this)
	if err != nil {
		return "", err
	} else {
		return string(jb[:]), nil
	}
}
