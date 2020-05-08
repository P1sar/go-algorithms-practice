package recursiveSum

import (
	"fmt"
	"math/big"
	"strconv"
)


func Pow(a, b uint64) uint64 {
	var p uint64 = 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}
func PowTenToBString(b uint64) string {
	r := "1"
	for i:=uint64(0); i< b; i++ {
		r = r + "0"
	}
	return r
}

func addString(a, b string) string {
	var au big.Int
	var bu big.Int
	au.SetString(a, 10)
	bu.SetString(b, 10)

	var res big.Int

	res.Add(&au, &bu)

	return res.String()
}
//11 11
func calibrateLength(x,y string) (string, string) {
	if len(x) > len(y) {
		for len(y) < len(x) {
			y = "0" + y
		}
	} else {
		for len(x) < len(y) {
			x = "0" + x
		}
	}
	if len(x) % 2 != 0 {
		fmt.Println("NOT ODD X")
		x = "0" + x
	}
	if len(y) %2 != 0 {
		fmt.Println("NOT ODD Y")
		y = "0" + y
	}


	return x, y
}

func multiplyOnConstant(c, x string) string {
	for i:=1; i < len(c); i++ {
		x = x + "0"
	}
	return x
}


func MultiplyRecursively(x, y string) string {
	fmt.Printf("Income: %v %v\n", x, y)
	lenX := len(x)
	lenY := len(y)
	if lenX <= 1 || lenY <= 1 {
		xUint, _ := strconv.ParseUint(x, 10, 64)
		yUint, _ := strconv.ParseUint(y, 10, 64)
		r := strconv.FormatUint(xUint * yUint, 10)
		fmt.Printf("result %v\n", r)
		return r
	}
	x, y = calibrateLength(x, y)
	fmt.Printf("Calibrated: %v %v\n", x,y)
	lenX = len(x)
	lenY = len(y)


	var a, b, c, d string
	a = x[0:lenX/2]
	b = x[lenX/2:lenX]
	c = y[0:lenY/2]
	d = y[lenY/2:lenY]
	fmt.Println(a,b,c,d)



	constant := PowTenToBString(uint64((lenX+lenY)/2))
	fmt.Printf("constant %v\n", constant)

	ac := MultiplyRecursively(a,c)
	fmt.Printf("ac %v\n", ac)

	first := multiplyOnConstant(constant, ac)
	fmt.Printf("first %v\n", first)

	ad := MultiplyRecursively(a,d)
	fmt.Printf("ad %v\n", ad)
	bc := MultiplyRecursively(b,c)
	fmt.Printf("bc %v\n", bc)

	second := addString(multiplyOnConstant(PowTenToBString(uint64(lenX/2)), ad),multiplyOnConstant(PowTenToBString(uint64(lenY/2)),bc))
	fmt.Printf("second %v\n", second)


	third := MultiplyRecursively(b,d)
	fmt.Printf("third %v\n", third)

	result := addString(addString(first, second),third)

	fmt.Printf("result %v\n", result)
	return result
}

//
//var a string = "3141592653589793238462643383279502884197169399375105820974944592"
//var b string = "2718281828459045235360287471352662497757247093699959574966967627"