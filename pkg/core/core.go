package core

import "fmt"

type Image struct {
  ( )

type IntPixel struct{
  values []int
  length int
}

type FloatPixel struct{
  values []float32
  length int
}


func (p IntPixel) print() {
  fmt.Println(p.values)
}
