package main

import (
	"fmt"
	"strconv"
	"strings"
	"math"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	temp := countAndSay(n - 1)
	result := ""
	for i := 0; i < 10; i++ {
		stemp := strconv.Itoa(i)
		if strings.Count(temp, stemp) > 0 {
			result = strconv.Itoa(strings.Count(temp, stemp)) + stemp + result
		}
	}
	return result
}

func convert(s string, numRows int) string {
	stringSli := make([]string, numRows)
	temp := ""
	result := ""
	for i := 0; i < int(math.Ceil(float64(len(s)) / float64(numRows+1))); i++ {
		if (i + 1) * (numRows+1)>len(s){
			temp = s[i * (numRows+1):]
		}else{
			temp = s[i * (numRows+1):(i + 1) * (numRows+1) ]
		}

		for j, x := range (temp) {
			if j  != numRows {
				stringSli[j] = stringSli[j] + string(x)
			}
			stringSli[(numRows - 1) / 2] = stringSli[(numRows - 1) / 2] + string(x)
		}
	}
	for _, x := range (stringSli) {
		result = result + x;
	}
	return result
}

func main() {

	fmt.Println(countAndSay(5))

	b := "abc"
	for _, x := range (b) {
		fmt.Println(string(x))

	}


	BB := "PAYPALISHIRING"
	result := convert(BB, 3)
	fmt.Println(result)






}