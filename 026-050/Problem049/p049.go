package main

import "math"
import "fmt"

// IsPrime ..
func IsPrime(n int) bool {
	flag := n > 1
	if n > 2 {
		flag = n%2 != 0
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			flag = false
			break
		}
	}
	return flag
}

func digits(p int) string {
	ds := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	sp := fmt.Sprint(p)
	for _, s := range sp {
		ds[s-'0']++
	}
	return fmt.Sprintf("%x%x%x%x%x%x%x%x%x%x", ds[0], ds[1], ds[2], ds[3], ds[4], ds[5], ds[6], ds[7], ds[8], ds[9])
}

func main() {
	m := make(map[string][]int)
	for p := 1001; p < 10000; p += 2 {
		if IsPrime(p) {
			ds := digits(p)
			// fmt.Println(p, ds)
			if _, ok := m[ds]; !ok {
				m[ds] = make([]int, 0)
			}
			m[ds] = append(m[ds], p)
		}
	}
	for k, v := range m {
		if len(v) < 3 {
			delete(m, k)
			continue
		}
		for i := 0; i < len(v)-2; i++ {
			for j := i + 1; j < len(v)-1; j++ {
				for k := j + 1; k < len(v); k++ {
					if v[j]-v[i] == v[k]-v[j] {
						fmt.Println(v[i], v[j], v[k])
					}
				}
			}
		}
	}

}
