package main

import "fmt"

type Point struct {
	X int32 `json:"x" yaml:"X-point"`
	Y int32 `json:"y" yaml:"Y-point"`
}

func (point Point) PrettyPrint() string {
	return fmt.Sprintf("(X:%v, Y:%v)", point.X, point.Y)
}
