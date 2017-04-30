package main

import "fmt"

func main() {
  count := 0

  for i := 2; i <= 10000; i += 2 {
    if collatz(i, i, true) {
      count++
    }
  }

  fmt.Println(count)
}

func collatz(origin, n int, init bool) bool {
  if init || n % 2 == 1 {
    n = n * 3 + 1
  } else if n % 2 == 0 {
    n /= 2
  }

  if n == origin {
    return true
  }

  if n == 1 {
    return false
  }

  return collatz(origin, n, false)
}
