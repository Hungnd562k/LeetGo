package problems

func RomanToInt(s string) int {
	// strRoman := ""
	// quoteChar := "\""
	// firstIndex := strings.Index(s, quoteChar)
	// lastIndex := strings.LastIndex(s, quoteChar)

	// if firstIndex != -1 && lastIndex != -1 && firstIndex < lastIndex {
	// 	strRoman = s[firstIndex+1 : lastIndex]
	// }
	// fmt.Println(strRoman)
	m := MakeMapRunes()
	// fmt.Println(m)

	arrChar := []rune(s)
	var arrVal []int
	if len(arrChar) == 1 {
		arrVal = append(arrVal, m[string(arrChar[0])])
	} else if len(arrChar) == 2 {
		a := m[string(arrChar[0])]
		b := m[string(arrChar[1])]
		if a < b {
			tmp := string(arrChar[0]) + string(arrChar[1])
			v, ok := m[tmp]
			if ok {
				arrVal = append(arrVal, v)
			} else {
				arrVal = append(arrVal, m[string(arrChar[0])])
				arrVal = append(arrVal, m[string(arrChar[1])])
			}
		}
	} else {
		v := m[string(arrChar[0])]
		arrVal = append(arrVal, v)
		for i := 1; i < len(s); i++ {
			if i != len(s)-1 {
				a := m[string(arrChar[i])]
				b := m[string(arrChar[i+1])]
				if a < b {
					tmp := string(arrChar[i]) + string(arrChar[i+1])
					v1, ok := m[tmp]
					if ok {
						arrVal = append(arrVal, v1)
						i++
					} else {
						v = m[string(arrChar[i])]
						arrVal = append(arrVal, v)
					}
				} else {
					v = m[string(arrChar[i])]
					arrVal = append(arrVal, v)
				}
			} else {
				v = m[string(arrChar[i])]
				arrVal = append(arrVal, v)
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
