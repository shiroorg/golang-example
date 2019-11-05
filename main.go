package main

import (
    "fmt"
)

func main()  {

    fmt.Print("Enter first number: ")
    var inputFirstNumber float32
    fmt.Scanf("%f", &inputFirstNumber)

    fmt.Print("Enter last number: ")
    var inputLastNumber float32
    fmt.Scanf("%f", &inputLastNumber)

    fmt.Print("Enter operation [+ - * /]: ")
    var inputOperation string
    fmt.Scanf("%s", &inputOperation)

    calculate(inputOperation, inputFirstNumber, inputLastNumber)
}

func calculate(operation string, FirstNumber float32, LastNumber float32)  {

    var result float32

    switch operation {
    case "+":
        result = FirstNumber + LastNumber
    case "-":
        result = FirstNumber - LastNumber
    case "*":
        result = FirstNumber * LastNumber
    case "/":
        result = FirstNumber / LastNumber
    default:
        fmt.Println("Unknown Operation")
    }

    fmt.Println(result)
}
