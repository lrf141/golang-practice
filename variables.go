package main;

import "fmt";

var c, python, java bool = true, false, true;

func main() {
    var i = 10;
    fmt.Println(i, c, python, java);

    scala, php := true, false;
    fmt.Println(scala, php);
}
