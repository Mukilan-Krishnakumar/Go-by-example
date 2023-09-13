package main

import "fmt"

// Variadic function can take in any number of trailing arguments, example fmt.Println
func sum(nums ...int)int{
  fmt.Println(nums, " ")
  total:=0
  for index, num:= range nums{
    fmt.Println("Index", index, "Value", num)
    total += num
  }
  return total
}

func main(){
  response := sum(1,2,3,4)
  fmt.Println("Total", response)

  check := []int{1,2,3,4}
  // You have to use spread it
  res := sum(check...)
  fmt.Println("After spreading", res)
}
