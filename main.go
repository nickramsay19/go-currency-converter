package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"strconv"
)

type Currency struct {
	name string
	rate float32
}

func main(){

	// get arguments and check if valid
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Failed: Please provide arguments. Use \"help\" to recieve helpful dialogue.")
	} else if args[0] == "help" || args[0] == "h" {
		fmt.Println("Showing help for Nick's Currency Converter\n\nUsage: curr <args>\n  curr help - Shows this help dialogue.\n  curr <VAL> <CURRENCY 1> <CURRENCY 2> - Shows the value of VAL in <URRENCY 1 when converted to CURRENCY 2.\n")
		os.Exit(1)
	}

	// get currency data from curr.csv
	content, err := ioutil.ReadFile("./curr.csv")
	if err != nil{
		log.Fatal(err)
	}

	// create base unit USD
	var usd Currency
	usd.name = "USD"
	usd.rate = 1.0

	// declare currencies slice
	var currs []Currency
	currs = append(currs, usd)

	// loop through each byte in content
	readingName := true
	var nameb []byte
	var rateb []byte
	var curr Currency

	for i := 0; i < len(content); i++ {
		c := string(content[i])

		if c == "\n" {
			// append new currency
			curr.name = string(nameb)
			rate, err := strconv.ParseFloat(string(rateb), 32)
			if err != nil {
				log.Fatal(err)
			}
			curr.rate = float32(rate)
			currs = append(currs, curr)

			// reset line reading
			readingName = true
			nameb = nameb[:0]
			rateb = rateb[:0]

		} else if c == "," {
			readingName = false
		} else if readingName {
			nameb = append(nameb, content[i])
		} else {
			rateb = append(rateb, content[i])
		}
	}

	// get value from args
	val64, err := strconv.ParseFloat(args[0], 32)
	if err != nil {
		fmt.Println("Error: value provided was invalid. Please provide a floating point number.")
		os.Exit(1)
	}
	val := float32(val64)

	// get currencies from args
	var fromCurr Currency
	var toCurr Currency
	for i := 0; i < len(currs); i++ {
		if currs[i].name == args[1] {
			fromCurr = currs[i]
		} else if currs[i].name == args[2] {
			toCurr = currs[i]
		}
	}

	// print the currency conversion
	fmt.Printf("%f\n", (val / fromCurr.rate) * toCurr.rate)
}