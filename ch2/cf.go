package main

import (
	"Go-by-myself/ch2/lengthconv"
	"Go-by-myself/ch2/tempconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.K(t)
		fmt.Printf("%s = %s, %s = %s\n%s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c), k, tempconv.KToF(k))
		fmt.Println("=========================")

		km := lengthconv.KM(t)
		m := lengthconv.M(t)
		dm := lengthconv.DM(t)
		fmt.Printf("%s = %s, %s = %s\n%s = %s\n",
			km, lengthconv.KMToM(km), m, lengthconv.MToDM(m), dm, lengthconv.DMToKM(dm))
	}
}
