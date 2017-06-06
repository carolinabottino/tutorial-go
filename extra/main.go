package main

import (
  "fmt"
  "os"
  "math"
)

const c = "C"
var v int = 5
var Pi float64
var a string

type T struct{}

func init() {
  // initialization of package
  Pi = 4 * math.Atan(1)
}

func main() {
  fmt.Println("--- Primera parte ---")
  var goos string = os.Getenv("GOOS")
  fmt.Println("The operating system is: %s\n", goos)
  path := os.Getenv("PATH")
  fmt.Println("Path is %s\n", path)

  fmt.Println("--- Segunda parte ---")
  a = "G" //asignación local al valor
  fmt.Println(a)
  f1()
}

func (t T) Method1() {
  //...
}

func f1() {
  a := "O" //asignación por referencia
  fmt.Println(a)
  f2()
}

func f2() {
  fmt.Println(a)
}
