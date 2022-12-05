package main

import (
  "fmt"
  //"core"
  //"./core"
  //"vita/core"
  //"./vita/core"
  //"core.go"
  "./core/core.go"
)

func main() {
  fmt.Println("Hello, world?")

  myfirstpixel := core.IntPixel[1,2,3]

  myfirstpixel.print()
}
