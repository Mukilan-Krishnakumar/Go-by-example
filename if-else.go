package main

import "fmt"

func main(){
  
  if i:= 9; i < 10{
    fmt.Println("Greatest Pirate ever")
  }

  if num:=8; num%2!=0{
    fmt.Println("How do you spell hoola hoops?")
  } else{
    fmt.Println(num,"Gopher or Goblin")
  }
  // You don't need paranthesis around expressions but need braces
  // In Go, there are no ternary statements

}
