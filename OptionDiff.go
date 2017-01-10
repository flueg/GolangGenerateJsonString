package main

import (
	"encoding/json"
)

type OptionDiff struct {
	Img1 string `json:"img1"`
	Img2 string `json:"img2"`
}

func NewOptionDiff() *OptionDiff {
	return &OptionDiff{Img1: "http:flueg.liu.com/abc/avatar.jpg"}
}

func (this *OptionDiff) GetJsonString() (string, error) {
	jb, err := json.Marshal(this)
	if err != nil {
		return "", err
	} else {
		return string(jb[:]), nil
	}
}
