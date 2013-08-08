package main

import (
	"bufio"
	"os"
	"runtime"
	"strings"
  "fmt"
  "math"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	scanner := bufio.NewScanner(os.Stdin)

	s := make(map[string]int)
	v := make(map[string]map[string]int)

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

	for key, value := range s {
		keys := strings.Split(key, "##")

    if v[keys[0]] == nil {
		  v[keys[0]] = make(map[string]int)

      // think about this one
		  // v[keys[0]][keys[0]] = value
    } else {
		  keys := strings.Split(key, "##")

		  v[keys[0]][keys[1]] = value
    }
	}

	for key, value := range v {
    dotproduct := 0
    magnitude := 0

    for subkey, subvalue := range value {
      // fmt.Println(key, subkey)
      magnitude += subvalue * subvalue
      dotproduct *= subvalue
    }
    fmt.Println(dotproduct, math.Sqrt(float64(magnitude)))
  }
  fmt.Println(v)
}
