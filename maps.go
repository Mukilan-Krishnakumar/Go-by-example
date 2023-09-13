package main

import (
  "fmt"
  "maps"
)

func main(){
  // Maps are dicts of go 
  m := make(map[string]int)

  m["luffy"] = 100
  m["zoro"] = 98

  fmt.Println(m)

  // Trying to print Non existent key
  // If the key doesn't exist, zero value of the value type is returned
  fmt.Println(m["sanji"])
  
  fmt.Println("Length of maps: ", len(m))

  // we can delete a key value pair and can use clear to remove all key value pairs
  delete(m, "zoro")
  fmt.Println("After deletion: ", m)

  clear(m)
  fmt.Println("After clearing: ", m)

  // To differentiate between missing keys and keys with values of zero
  _, prs := m["sanji"]
  fmt.Println("Does sanji key exist: ", prs)

  // Declare and Initialize in the same line
  n := map[string]int{"Luffy":100, "Zoro":98}
  n2 := map[string]int{"Luffy":100, "Zoro":98}
  if maps.Equal(n, n2){
    fmt.Println("These are equal", n)
  }

}
