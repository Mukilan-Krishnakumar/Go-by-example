package main

import (
  "fmt"
  "slices"
)

func main(){
  // Slices are more powerful than arrays, they are only typed by the elements
  // they contain and not by the number of elements
  var s[] string
  fmt.Println("Uninitialized Slice:", s, s == nil, len(s) == 0)

  // Use make function to initialize slice of non null value
  s = make([]string, 3)
  fmt.Println("Uninit with bigger size:", s, "len: ", len(s), "cap: ", cap(s))

  // Slices can be appended but we have to accept the returned slice
  s[0] = "Luffy"
  s[1] = "Zoro"
  s[2] = "Sanji"
  s = append(s, "Nami")
  s = append(s, "Robin", "Tony")
  fmt.Println("The One Piece crew: ", s)

  // We can also copy the slice
  c := make([]string, len(s))
  copy(c,s)
  fmt.Println("Here is the copy of One Piece crew: ", c)

  // Slices can be sliced, a joke by brook
  // Slice similar to python slice

  l := s[0:3]
  fmt.Println("The Monster Trio: ", l)

  t1 := []string{"Nami", "Usoop", "Tony"}
  fmt.Println("The scared Trio: ", t1)

  t2 := []string{"Nami", "Usoop", "Tony"}
  if slices.Equal(t1, t2){
    fmt.Println("Scared Trio is equally scared")
  }

  // Slices can be multi-dim with each dim having varying size
  twoD := make([][]int, 3)
  for i:= 0; i < 3; i++{
    innerLen := i + 1
    twoD[i] = make([]int, innerLen)
    for j:= 0; j < innerLen; j++{
      twoD[i][j] = i + j
    }
  }

  fmt.Println("Two Dimensional Array: ", twoD)

}
