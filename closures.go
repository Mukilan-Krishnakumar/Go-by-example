package main

import "fmt"

func onePiece(index int) func() string{
  onePiece_list := []string{"Grand Line", "East Blue", "West Blue", "All Blue"}
  // This return func closes over the variable onePiece_list to form a closure
  return func() string{
    response := ""
    if index < len(onePiece_list){
      response = onePiece_list[index]
    }
    return response
  }
}

func main(){
  getChar := onePiece(1)
  // Have to call the function returned
  fmt.Println(getChar())

}
