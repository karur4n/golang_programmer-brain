package main

import "fmt"

func main() {
  fmt.Println(cutBar(3, 20, 1))
  fmt.Println(cutBar(5, 100, 1))
}

func cutBar(m, n, current int) int {
  if current >= n {
    return 0
  } else if current < m {
    return 1 + cutBar(m, n, current * 2)
  } else {
    return 1 + cutBar(m, n, current + m)
  }
}
