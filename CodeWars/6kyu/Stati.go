package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func secondFormat(num float64) string {
	num = num / 3600
	h, m := math.Modf(num)
	m *= 60
	m, s := math.Modf(m)
	s *= 60
	h_str, m_str, s_str := fmt.Sprintf("%02d", int(h)), fmt.Sprintf("%02d", int(m)), fmt.Sprintf("%02d", int(s))

	return h_str + "|" + m_str + "|" + s_str
}

func Stati(strg string) string {
	strg = strings.ReplaceAll(strg, " ", "")
	strsplit := strings.Split(strg, ",")
	if len(strsplit) <= 1 {
		return strsplit[0]
	}
	A := 0
	sec := 0
	min, max := 100000, 0
	median := []int{}
	for _, i := range strsplit {
		nStr := strings.Split(i, "|")
		h, _ := strconv.Atoi(nStr[0])
		m, _ := strconv.Atoi(nStr[1])
		s, _ := strconv.Atoi(nStr[2])
		sec = (h * 3600) + (m * 60) + s
		A += sec
		median = append(median, sec)
		if sec > max {
			max = sec
		}
		if sec < min {
			min = sec
		}
	}
	sort.Ints(median)
	l := len(median)
	med := 0
	if l%2 == 0 {
		med = (median[l/2-1] + median[l/2]) / 2
	} else {
		med = median[l/2]
	}

	R := float64(max - min)
	Average := float64(A) / float64(len(strsplit))
	return "Range: " + secondFormat(R) + " Average: " + secondFormat(Average) + " Median: " + secondFormat(float64(med))
}

func main() {
	fmt.Println(Stati("01|15|59, 1|47|16, 01|17|20, 1|32|34, 2|17|17"))
}
