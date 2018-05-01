package GoHTML

import (
	"fmt"
	"strconv"
	"strings"
)

func RelPath(sPath string, tPath string) string {
	ret := ""

	sP := strings.Split(sPath, "/")
	tP := strings.Split(tPath, "/")

	fmt.Println(sP)
	fmt.Println(tP)

	lens := len(sP)
	lentP := len(tP)

	far := lens
	if lentP < lens {
		far = lentP
	}

	fmt.Println(strconv.Itoa(lens) + " " + strconv.Itoa(lentP))

	dist := 0

	for j := 0; j < far-1 && sP[j] == tP[j]; j++ {
		dist++
	}

	fmt.Println(dist)

	for j := lens - dist - 1; j > 0; j-- {
		ret += "../"
	}

	for ; dist < lentP-1; dist++ {
		ret += tP[dist] + "/"
	}
	ret += tP[dist]

	fmt.Println("Result: " + ret)
	return ret
}
