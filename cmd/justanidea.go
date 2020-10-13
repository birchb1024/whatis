package main

//
// Classify some atoms - return interfaces. despatch on type of interface.
//
import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type url string
type number string

func (u url) String() string    { return "<a href=\"" + string(u) + "\"/>" }
func (n number) String() string { return "#" + string(n) }

func whatIs(s string) interface{} {
	if strings.HasPrefix(s, "http://") {
		return url(s)
	}
	if _, err := strconv.Atoi(s); err == nil {
		return number(s)
	}
	return s
}

type bigint big.Int
type prime bigint

func (i bigint) whatIs() interface{} {
	bi := big.Int(i)
	if bi.ProbablyPrime(0) {
		return prime(i)
	}
	return i
}

func main() {
	s := whatIs(`http://www.genyris.org/`)
	u, ok := s.(url)
	fmt.Printf("%T %+v %v\n", u, u, ok) //main.url <a href="http://www.genyris.org/"/> true

	s = whatIs(`34`)
	n, ok := s.(number)
	fmt.Printf("%T %+v %v\n", n, n, ok) // main.number #34 true

	x := big.NewInt(23)
	i := bigint(*x)
	b := i.whatIs()
	fmt.Printf("%T %+v %v\n", b, b, ok) // main.prime {neg:false abs:[23]} true

}
