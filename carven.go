package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"runtime"
	"runtime/pprof"
	"strings"
)

type ParseData map[string]int
type MultiSetData map[string]map[string]int

func CalcDotProduct(v1, v2 []float64) float64 {
	result := 0.0

	for index, _ := range v1 {
		result += v1[index] * v2[index]
	}

	return result
}

func CalcMagnitude(v []float64) float64 {
	if len(v) == 0 {
		return 0.0
	}

	var subtotal float64

	for _, somefloat := range v {
		subtotal += float64(somefloat * somefloat)
	}

	return math.Sqrt(subtotal)
}

func BuildVector(outerkey, innerkey string, outervalue, innervalue map[string]int) ([]float64, []float64) {
	outervector := make([]float64, 0)
	innervector := make([]float64, 0)

	commonkeys := make(map[string]bool)

	for outerstring, _ := range outervalue {
		commonkeys[outerstring] = true
	}

	for innerstring, _ := range innervalue {
		commonkeys[innerstring] = true
	}

	if len(commonkeys) == 0 {
		return outervector, innervector
	}

	for commonstring, commonbool := range commonkeys {
		if commonbool == true {
			outervector = append(outervector, float64(outervalue[commonstring]))
			innervector = append(innervector, float64(innervalue[commonstring]))
		}
	}

	return outervector, innervector
}

func CalcCosim(v1, v2 []float64) float64 {
	dotproduct := CalcDotProduct(v1, v2)

	if dotproduct > 0 {
		v1m := CalcMagnitude(v1)
		v2m := CalcMagnitude(v2)

		return dotproduct / (v1m * v2m)
	} else {
		return 0.0
	}
}

func CleanString(s string) string {
	reg, err := regexp.Compile("[^A-Za-z ]+")

	if err != nil {
		log.Fatal(err)
	}

	safe := reg.ReplaceAllString(s, "")
	safe = strings.ToLower(strings.Trim(safe, ""))
	safe = strings.TrimSpace(safe)

	return safe
}

func Parse(m MultiSetData) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		somestring := CleanString(scanner.Text())

		if len(somestring) == 0 {
			continue
		}

		t := strings.Split(somestring, " ")

		if len(t) <= 1 {
			continue
		}

		for i := 0; i < (len(t) - 1); i++ {
			if m[t[i]] == nil {
				m[t[i]] = make(map[string]int)
			}
			m[t[i]][t[i+1]] += 1

			if m[t[i+1]] == nil {
				m[t[i+1]] = make(map[string]int)
			}
			m[t[i+1]][t[i]] += 1
		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	f, err := os.Create("carvan.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	m := make(MultiSetData)

	fmt.Println("GOing to parse")

	Parse(m)

	var keys []string

	if len(os.Args) > 1 {
		keys = os.Args[1:]
	}

	fmt.Println("GOing to find any matches")

	for _, outerkey := range keys {
		for innerkey, innervalue := range m {
			v1, v2 := BuildVector(outerkey, innerkey, m[outerkey], innervalue)

			if cosign := CalcCosim(v1, v2); cosign > 0.78 {
				fmt.Println(outerkey, innerkey, "\t\t", math.Floor(cosign*100))
			}
		}
	}
}
