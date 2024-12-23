package main

import (
	"fmt"
	"os"
	"strconv"
	
	"bookProjects/ch2/cf/cfmfkf"
)

func main() {
	for _, arg := range os.Args[1:] {
		val, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cfmfkf: %v\n", err)
			os.Exit(1)
		}
		f := cfmfkf.Fahrenheit(val)
		c := cfmfkf.Celsius(val)
		m := cfmfkf.Meter(val)
		ft := cfmfkf.Foot(val)
		kg := cfmfkf.Kilogram(val)
		lb := cfmfkf.Lb(val)
		fmt.Printf("%s = %s, %s = %s\n", f, cfmfkf.FToC(f), c, cfmfkf.CToF(c))
		fmt.Printf("%s = %s, %s = %s\n", m, cfmfkf.MToFt(m), ft, cfmfkf.FtToM(ft))
		fmt.Printf("%s = %s, %s = %s\n", kg, cfmfkf.KgToLb(kg), lb, cfmfkf.LbToKg(lb))

	}

}
