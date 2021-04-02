package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/yuzuy/fizzbuzz"
)

func main() {
	flag.Parse()

	limitStr := flag.Arg(0)
	if limitStr == "" {
		limitStr = "100"
	}
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		fmt.Println("fizzbuzz: Please specify an integer for limit")
		os.Exit(1)
	}
	if limit == 0 {
		limit = math.MaxInt64
	}

	var result string
	var i int64 = 1
	for ; i <= limit; i++ {
		result += fizzbuzz.FizzBuzz(i) + " "
	}
	fmt.Println(result)
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s <limit>\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "  limit is set 100 by default. if set 0, set the max of int32\n")
	}
}
