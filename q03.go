package main

import "fmt"

func main() {
  card_length := 100
	cards := make([]bool, card_length)

	for i := 0; i < card_length; i++ {
    cards[i] = true
	}

  for i := 1; i <= card_length; i++ {
    for j := i; j <= card_length; j += i {
      cards[j-1] = !cards[j-1]
    }
  }

  for i := 0; i < card_length; i++ {
    if cards[i] == false {
      fmt.Println(i + 1)
    }
  }
}
