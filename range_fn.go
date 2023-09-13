package main

import "fmt"

func main(){
  // Range function is used to iterate over elements in different data structures
  nums := []int{42, 69, 96}
  sum := 0
  for i, num := range nums{
    if num == 69{
      fmt.Println("Caught 69 at index", i)
    }
    sum += num
  }
  fmt.Println("Total Sum: ", sum)
  one_piece := map[string]string{"King of Pirates": "Luffy D. Monkey", "Greatest Swordsman": "Zoro", "DevilChild": "Nico Robin"}
  for k, v := range one_piece{
    fmt.Println(k, v)
  }
  for k := range one_piece{
    fmt.Println(k)
  }

  for i, char:= range "one piece"{
    fmt.Println(i, char)
  }
}
