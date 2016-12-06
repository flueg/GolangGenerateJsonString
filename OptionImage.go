package main

import (
	"encoding/json"
)

type OptionImage struct {
	A    string `json:"A"`
	B    string `json:"B"`
	C    string `json:"C"`
	D    string `json:"D"`
	ImgA string `json:"ImgA"`
	ImgB string `json:"ImgB"`
	ImgC string `json:"ImgC"`
	ImgD string `json:"ImgD"`
}

func NewOptionImage() *OptionImage {
	return &OptionImage{
		A:    "Let's go",
		B:    "Let's run",
		C:    "Let's try",
		D:    "",
		ImgA: "http:flueg.liu.com/abc/oa.jpg",
		ImgB: "http:flueg.liu.com/abc/ob.jpg",
		ImgC: "http:flueg.liu.com/abc/oc.jpg",
		ImgD: ""}
}

func (this *OptionImage) GetJsonString() (string, error) {
	jb, err := json.Marshal(this)
	if err != nil {
		return "", err
	} else {
		return string(jb[:]), nil
	}
}
