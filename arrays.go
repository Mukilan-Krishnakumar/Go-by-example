package main

import (
  "fmt"
)

func main(){
  // Declaring an integer array, which can only hold 5 values
  var a [5]int
  fmt.Println("empty array:", a)

  //Setting values
  a[4] = 420
  fmt.Println("After adding", a)

  // Declare and initialize in a single line
  b := [5]int{1, 3, 5, 7, 9}
  fmt.Println("Declaring and Initializing", b)
  
  // Two dimensional array
  var bigarr [2][3]int
  for i:=0;i<2;i++{
    for j:= 0;j<3;j++{
      bigarr[i][j] += i + j
    }
  }

  fmt.Println("2 Dimensional Array", bigarr)
}
