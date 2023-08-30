package main

import "fmt"

func main(){
  var a = "initial"
  fmt.Println(a)

  var b, c int = 8, 10
  fmt.Println(b, c)

  // Variables initialized without values are zero-valued
  var d int
  fmt.Println(d)

  var e = true
  fmt.Println(e)
  // := is shorthand for declaring and initializing a variable
  f:= "apple"
  fmt.Println(f)
}
