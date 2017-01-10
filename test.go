package main

import (
	"encoding/json"
	"fmt"
)

type Question struct {
	Type   int
	Desc   string
	Opt    IOption
	Answer string
}

func Test() {
	fmt.Println("Hello world.")
	opT := NewOptionText()
	jb, err := json.Marshal(&opT)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(jb[:]))
	}
	fmt.Println("End.")

	ques := []Question{}
	q := Question{Type: 1, Desc: "做什么好呢？", Opt: NewOptionText(), Answer: "A"}
	ques = append(ques, q)
	q = Question{Type: 2, Desc: "我在哪儿？", Opt: NewOptionDiff(), Answer: "2,3"}
	ques = append(ques, q)
	q = Question{Type: 3, Desc: "哈哈哈哈？", Opt: NewOptionImage(), Answer: "B"}
	ques = append(ques, q)
	for _, qq := range ques {
		opt, _ := qq.Opt.GetJsonString()
		fmt.Printf("Ques1: [%d][%s][%s][%s]\n", qq.Type, qq.Desc, opt, qq.Answer)
	}
}

func main1() {
	//Test()
	inFile := "questions"
	access := NewAccess()
	err := access.Init(inFile)
	if err != nil {
		fmt.Printf("Failed to open file. %s\n", err)
		return
	}
	err = access.Load()
	if err != nil {
		fmt.Printf("Failed to read from file. %s\n", err)
		return
	}
	fmt.Println("Processing file succeed.")
	access.Dump()
	fmt.Println("Dump result succeed.")
}
