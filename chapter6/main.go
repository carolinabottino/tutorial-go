package main
import "fmt"

func average(xs []float64) float64 {
  total := 0.0

  for _, v := range xs {
    total += v
  }

  return total / float64(len(xs))
}

func main() {
    xs := []float64{98,93,77,82,83}
    total := 0.0

    for _, v := range xs { //traducción: para cada valor (v) dentro del xs...
      total += v
    }

    fmt.Println(total / float64(len(xs)))

    fmt.Println("- otra forma de hacer lo mismo: ", average(xs))
}
