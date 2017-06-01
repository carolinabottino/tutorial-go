package main
import "fmt"

var a string = "cinco"

func main() {
    var x string = "Hello, Caro! Te mandé "
    y := " mensajes ayer." //lo mismo que antes, Go hace inferencia de tipo
    z := x + a + y //concatenación

    fmt.Println(z)

    fmt.Print("Segunda parte. \n Enter a number: ")
    var input float64
    fmt.Scanf("%f", &input)

    output := input * 2
    fmt.Println("- output: ", output)
}
