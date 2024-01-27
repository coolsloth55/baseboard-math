package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//boards := []string{"2:12", "1:10", "1:7", "1:5", "2:4", "1:3", "1:8", "1:6", "1:4", "1:3", "1:13", "1:10", "2:8", "1:4", "2:2.5", "1:12", "2:8", "1:7", "2:2.5", "1:10", "1:8", "1:6", "2:5", "1:4"}
	boards := []string{"2:12", "1:10", "1:7", "1:5", "2:4", "2:8", "1:8", "1:6", "1:4", "1:13", "1:10", "2:8", "1:4", "2:2.5", "1:12", "2:8", "1:7", "2:2.5", "1:10", "1:8", "1:6", "1:4"}

	var m = makeMap(boards)
	var f = make(map[string]int)

	fmt.Println(sumMap(m), m)

	combine("4", "7", &f, &m)

	for k := range m {
		if k == "8" || k == "12" || k == "10" {
			addToMap(&f, k, m[k])
			delete(m, k)
			continue
		}

		b := checkMult(m[k], k, 8.0)
		if b {
			fmt.Println(fmt.Sprintf("reduce %d %s to 8", m[k], k))
			incrementMap(&f, "8")
			delete(m, k)
			continue
		}

		b = checkMult(m[k], k, 12)
		if b {
			fmt.Println(fmt.Sprintf("reduce %d %s to 12", m[k], k))
			incrementMap(&f, "12")
			delete(m, k)
			continue
		}
	}

	fmt.Println(f, m)
	fmt.Println(sumMap(f), sumMap(m))

}

func incrementMap(m *map[string]int, k string) {
	_, exists := (*m)[k]
	if exists {
		(*m)[k] = (*m)[k] + 1
	} else {
		(*m)[k] = 1
	}
}

func addToMap(m *map[string]int, k string, v int) {
	_, exists := (*m)[k]
	if exists {
		(*m)[k] = (*m)[k] + v
	} else {
		(*m)[k] = v
	}
}

func checkMult(n int, lstr string, max float64) bool {
	l, _ := strconv.ParseFloat(lstr, 64)
	t := float64(n) * l
	if t <= max {
		return true
	}
	return false
}

func combine(k string, k1 string, m *map[string]int, m1 *map[string]int) {
	lowestKey := k
	addCount := 0
	remainingCount := 0

	ik := (*m1)[k]
	ik1 := (*m1)[k1]

	if ik < ik1 {
		lowestKey = k1
		addCount = ik1 - ik
		remainingCount = ik
	} else {
		addCount = ik - ik1
		remainingCount = ik1
	}

	addToMap(m, "12", addCount)
	delete(*m1, lowestKey)
	addToMap(m1, lowestKey, remainingCount)
	delete(*m1, k1)
}

func makeMap(l []string) map[string]int {
	var m = make(map[string]int)
	for _, val := range l {
		arr := strings.Split(val, ":")
		val, exists := m[arr[1]]
		temp, _ := strconv.Atoi(arr[0])
		if exists {
			m[arr[1]] = temp + val
		} else {
			m[arr[1]] = temp
		}
	}
	return m
}

func sumMap(m map[string]int) float64 {
	total := 0.0
	for k := range m {
		v, _ := strconv.ParseFloat(k, 64)
		t := v * float64(m[k])
		total = total + t
	}
	return total
}
