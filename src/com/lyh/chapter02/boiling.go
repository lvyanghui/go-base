package chapter02

import (
	"fmt"
	"flag"
	"strings"
)

const boilingF  = 212.0

func WaterBoiling(){
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gF or %gC\n", f, c)
}

func fToC(f float64) float64{
	return (f - 32) * 5 / 9
}

func SwitchResult(){
	const freezingF, boildingF = 32.0, 212.0
	fmt.Printf("%gF = %gC\n", freezingF, fToC(freezingF))
	fmt.Printf("%gF = %gC\n", boildingF, fToC(boildingF))
}

func TestFlag(){
	var n = flag.Bool("n", false, "omit trailing newline")
	var sep = flag.String("s"," ", "separator")
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !* n{
		fmt.Println()
	}
}