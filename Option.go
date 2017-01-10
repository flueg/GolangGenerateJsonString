package main

import (
	"fmt"
)

type IOption interface {
	GetJsonString() (string, error)
}

func NewOption(t string, fields []string) IOption {
	switch t {
	case "1":
		opt := NewOptionText()
		opt.A = fields[0]
		opt.B = fields[1]
		opt.C = fields[2]
		if len(fields) > 3 {
			opt.D = fields[3]
		}
		return opt
	case "2":
		opt := NewOptionDiff()
		opt.Img1 = fields[0]
		opt.Img2 = fields[1]
		return opt
	case "3":
		opt := NewOptionImage()
		i := 0
		opt.A = fields[i]
		i++
		opt.B = fields[i]
		i++
		opt.C = fields[i]
		i++
		if len(fields) > 6 {
			opt.D = fields[i]
			i++
		}
		opt.ImgA = fields[i]
		i++
		opt.ImgB = fields[i]
		i++
		opt.ImgC = fields[i]
		i++
		if len(fields) > 6 {
			opt.ImgD = fields[i]
			i++
		}
		return opt
	default:
		fmt.Printf("Option type not found.\n")
		return nil
	}

	//	return &opt
}
