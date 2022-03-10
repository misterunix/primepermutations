package main

import (
	"flag"
	"fmt"
	"math/big"
	"strings"

	gd "github.com/misterunix/cgo-gd"
)

type Pattern struct {
	Pattern  string
	Location []int
}

// prec : We'll do computations with 200 bits of precision in the mantissa.
const prec = 10000

var ibuf0 *gd.Image

func main() {
	var p int

	flag.IntVar(&p, "p", 11, "Primenumber")
	flag.Parse()

	var numbers []string // numbers : slice containing the string representation of the permutation

	//var patterns []Pattern // patterns : the patteren for each permutation

	fmt.Println("Program start.")

	ibuf0 = gd.CreateTrueColor(800, 800)

	// Prime : Arbitrary-precision value of the working prime number
	Prime := new(big.Float).SetPrec(prec).SetInt64(int64(p))

	//Prime := new(big.Float)
	//b := Prime.SetPrec(prec)
	//fmt.Println(b)
	//Prime.SetInt64(int64(p))

	for i := 1; i < p; i++ {
		divisor := new(big.Float).SetPrec(prec).SetInt64(int64(i))
		tmp := new(big.Float)
		tmp.Quo(divisor, Prime)
		//tmps1 := fmt.Sprintf("%.360f\n", tmp)
		tmps1 := fmt.Sprintf("%.3000f\n", tmp)
		tmps2 := strings.TrimPrefix(tmps1, "0.")

		numbers = append(numbers, tmps2)
	}

	fmt.Println("number")
	fmt.Println(numbers[0])
	//for i := 0; i < len(numbers); i++ {
	//	fmt.Println(numbers[i])
	//}

	var p1 string
	// find the pattern
	for i := 1; i < len(numbers[0]); i++ {

		p1 = numbers[0][0:i] // numbers 0 - i

		//fmt.Println(":", p1, ":")

		p2 := numbers[0][i : i+len(p1)]

		q := strings.Count(p2, p1)
		//fmt.Println(q, p1)
		if q == 1 {
			break
		}

		/*
			for j := len(p1) + 1; j < len(numbers[0])-len(p1); j++ {
				y := numbers[0][j : j+len(p1)]

				q := strings.Count(p1, y)
				if strings.Compare(p1, y) == 0 {
					fmt.Printf("q: %d p1: %s, y: %s\n", q, p1, y)
				}
			}
		*/
	}

	fmt.Println("Key:", p1)

	/*

		// Loop through the whole numbers to get the prime permutations
		for i := 1; i < p; i++ {
			t1 := new(big.Float).SetPrec(prec).SetInt64(int64(i))
			t := new(big.Float)

			t.Quo(t1, Prime)

			n := fmt.Sprintf("%.360f", t)
			nn := strings.TrimPrefix(n, "0.")
			fmt.Println(i, nn)
			numbers = append(numbers, nn)
			//fmt.Printf("%d\t%s\n", i, nn)
		}

		// Search for patterns
		nop := len(numbers) // nop : Number of items in the slice 'numbers'

		for i := 0; i < nop; i++ {
			strlen := len(numbers[i])
			var tp string

			for j := 0; j < strlen; j++ {
				tp = numbers[i][0:j]

				length := len(tp)
				if length == 0 {
					continue
				}

				ntp := numbers[i][j : j+length]
				fmt.Println(i, j, len(tp), len(ntp), tp, ":", ntp)

				if strings.Compare(tp, ntp) == 0 {
					fmt.Println("Match:", tp)
					pp := Pattern{}
					pp.Pattern = tp
					//pp.Location = append(pp.Location, 1)

					for k := 0; k < nop; k++ {
						tn := strings.Index(numbers[k], tp)
						pp.Location = append(pp.Location, tn)
					}

					patterns = append(patterns, pp)
					break
				}

			}

		}

		fmt.Println("Showing patterns")

		for _, s := range patterns {
			fmt.Printf("%s - ", s.Pattern)
			for _, ii := range s.Location {
				fmt.Printf("%02d ", ii)
			}
			fmt.Println()
		}

	*/

}
