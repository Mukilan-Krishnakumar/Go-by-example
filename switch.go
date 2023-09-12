package main

import (
  "fmt"
  "time"
)

func main(){
  i := 2
  fmt.Println("Initializing i to 2")
  switch i{
  case 1:
    fmt.Println("Got one")
  case 2:
    fmt.Println("Got two")
  }

  // Using commas to seperate conditions
  switch time.Now().Weekday(){
    case time.Saturday, time.Sunday:
      fmt.Println("King of Pirates")
    default:
      fmt.Println("Noob")
  }

  whatAmI := func(i interface{}){
    switch t := i.(type){
    case bool:
       fmt.Println("Bool")
    case int:
      fmt.Println("Int")
    default:
      fmt.Println("Don't know", t)
    }  
  }
  whatAmI(1)
  whatAmI(true)
  whatAmI("Luffy")
}
