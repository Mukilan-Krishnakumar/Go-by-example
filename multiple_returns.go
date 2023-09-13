package main

import "fmt"

//Multiple returns typed

func greatness()(string,string){
  return "J Cole", "Kendrick Lamar"
}
// Use blank identifier to only receive subset values
func main(){
  a, b:= greatness()
  fmt.Println(a, b)
  j_cole, _ := greatness()
  fmt.Println(j_cole)
}
