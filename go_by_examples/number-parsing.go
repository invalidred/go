package main

import (
	"fmt"
	"strconv"
)

func main() {
	p := fmt.Println
	f, _ := strconv.ParseFloat("1.2345", 64)
	p(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	p(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	p(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	p(u)

	k, _ := strconv.Atoi("135")
	p(k)

	_, e := strconv.Atoi("wat")
	p(e)
}
