package main

import "fmt"

func add(a int, b int) int{
  return a + b
}

func subtract(a, b int)int{
  return b - a
}

// In go, we have to set return type, it won't accept without it 
func main(){
  res := add(1,2)
  fmt.Println("Addition result: ", res)
  sub := subtract(1,res)
  fmt.Println("Subtraction result: ", sub)
}
