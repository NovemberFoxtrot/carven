package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	scanner := bufio.NewScanner(os.Stdin)

	s := make(map[string]int, 0)

	for scanner.Scan() {
		somestring := scanner.Text()

		t := strings.Split(somestring, " ")

		for i := 0; i < (len(t) - 1); i++ {
			k1 := t[i] + "##" + t[i+1]
			k2 := t[i+1] + "##" + t[i]

			s[k1] += 1
			s[k2] += 1
		}
	}

	v := make(map[string]map[string]int)

	for key, value := range s {
		keys := strings.Split(key, "##")

		v[keys[0]] = make(map[string]int)

		v[keys[0]][keys[0]] = value
	}

	for key, value := range s {
		keys := strings.Split(key, "##")

		v[keys[0]][keys[1]] = value
	}

	for key, value := range v {
		if len(value) > 2000 {
			fmt.Println(key, value)
		}
	}
}
