package main

import "fmt"

func zeroval(ival int){
  ival = 0
  fmt.Println("Inside fn, i value: ", ival)
}

func zeroptr (iptr *int){
  // Deferences the pointer from its memory address to the current value at that address. Assigning value for dereferenced pointer changes the value at referenced address. 
  *iptr = 0
}

func main(){
  i := 10

  fmt.Println("The original i value: ", i)

  fmt.Println("Calling zeroval fn")
  zeroval(i)
  fmt.Println("The original i value: ", i)
  // & for accessing pointers

  fmt.Println("Calling zeroptr fn")
  zeroptr(&i)
  fmt.Println("The origina i value: ", i)
  fmt.Println("The pointer location for i: ", &i)
}
