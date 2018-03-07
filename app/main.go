package main

import (
  "log"
)

func main() {
  err := CopyDir("/src", "/dest")
  if err != nil {
    log.Fatal(err)
  }
}
