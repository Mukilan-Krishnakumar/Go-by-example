package main

import "fmt"

func factorial(num int) int{
  if num == 0{
    return 1
  }
  return num * factorial(num - 1)
}

func main(){
  fmt.Println("Factorial for 5:", factorial(5))
  fmt.Println("Factorial for 10:", factorial(10))
  // Recursive closures
  // Should be a variable and should be declared explicitly before defined
  var fibonaci func(num int) int

  fibonaci = func(num int) int{
    if num < 2{
      return num
    }
    return fibonaci(num-1) + fibonaci(num-2)
  }

  fmt.Println("Recursive Fibonaci of 11: ", fibonaci(11))
}
