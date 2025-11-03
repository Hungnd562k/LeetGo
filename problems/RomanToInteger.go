package problems

func RomanToInt(s string) int {
	m := MakeMapRunes()

	arrChar := []rune(s)
	var arrVal []int

	if len(arrChar) == 1 {
		v, ok := m[string(arrChar[0])]
		if ok {
			arrVal = append(arrVal, v)
		}
	} else if len(arrChar) == 2 {
		v, ok := m[string(arrChar[0])+string(arrChar[1])]
		if ok {
			arrVal = append(arrVal, v)
		} else {
			arrVal = append(arrVal, m[string(arrChar[0])])
			arrVal = append(arrVal, m[string(arrChar[1])])
		}
	} else {
		for i := 0; i < len(arrChar); i++ {
			if i == len(arrChar)-1 {
				arrVal = append(arrVal, m[string(arrChar[i])])
				break
			}
			a := m[string(arrChar[i])]
			b := m[string(arrChar[i+1])]
			if a < b {
				v, ok := m[string(arrChar[i])+string(arrChar[i+1])]
				if ok {
					arrVal = append(arrVal, v)
					i++
				} else {
					arrVal = append(arrVal, m[string(arrChar[i])])
				}
			} else {
				arrVal = append(arrVal, m[string(arrChar[i])])
			}
		}
	}

	var sum int
	for _, v := range arrVal {
		sum += v
	}
	return sum
}

func MakeMapRunes() map[string]int {
	m := make(map[string]int)
	m["I"] = 1
	m["IV"] = 4
	m["V"] = 5
	m["IX"] = 9
	m["X"] = 10
	m["XL"] = 40
	m["L"] = 50
	m["XC"] = 90
	m["C"] = 100
	m["CD"] = 400
	m["D"] = 500
	m["CM"] = 900
	m["M"] = 1000

	return m
}
