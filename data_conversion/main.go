// ...existing code...
package main

import (
    "fmt"
    "strconv"
)

func main() {
    var number int = 10
    str := strconv.Itoa(number)
    fmt.Printf("int -> string: %d -> %q\n", number, str)

    s := "42"
    n, err := strconv.Atoi(s)
    if err != nil {
        fmt.Println("Atoi error:", err)
    } else {
        fmt.Printf("string -> int: %q -> %d\n", s, n)
    }

    f := 3.1415
    sf := strconv.FormatFloat(f, 'f', 4, 64)
    fmt.Printf("float64 -> string: %v -> %q\n", f, sf)

    s2 := "2.718"
    f2, err := strconv.ParseFloat(s2, 64)
    if err != nil {
        fmt.Println("ParseFloat error:", err)
    } else {
        fmt.Printf("string -> float64: %q -> %v\n", s2, f2)
    }
}