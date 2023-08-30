package main

import (
  "fmt"
  "math"
)

const s string = "Golang :="

func main(){
  fmt.Println(s)

  const n = 500000000 / 5
  fmt.Println(n)
  // A numeric constant has no type until it is given one, such as explicit conversion 
  fmt.Println(int64(n))

  fmt.Println(math.Sin(n))

}
