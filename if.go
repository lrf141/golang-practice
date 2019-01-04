package main;

import "fmt";
import "math";


func sqrt(x float64) string {
    if x < 0 {
        return sqrt(-x) + "i";
    }
    return fmt.Sprint(math.Sqrt(x));
}

func pow(x float64, n float64, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v;
    }
    return lim;
}

func main() {
    fmt.Println(sqrt(2), sqrt(-4));
    fmt.Println(pow(3, 2, 10), pow(3, 3, 20));
}
