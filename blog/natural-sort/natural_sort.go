package natural_sort

import (
	"math"
	"regexp"
	"strconv"
)

const PrintFormat = "%-25s: %v\n"

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

// NaturalLess02 compare string by separating number and string manually
func NaturalLess02(str1, str2 string) bool {
	pStr1 := processStr(str1)
	pStr2 := processStr(str2)
	minLen := int(math.Min(float64(len(pStr1)), float64(len(pStr2))))
	for i := 0; i < minLen; i++ {
		e1 := pStr1[i]
		e2 := pStr2[i]
		if e1 == e2 {
			continue
		}
		ei1, err1 := strconv.Atoi(e1)
		ei2, err2 := strconv.Atoi(e2)

		switch {
		case err1 == nil && err2 == nil && ei1 != ei2:
			return ei1 < ei2
		case err1 != nil && err2 != nil:
			return e1 < e2
		case err1 != nil:
			return true
		case err2 != nil:
			return false
		}
	}
	return len(str1) < len(str2)
}

func processStr(str string) []string {
	var processedStr []string
	for i := 0; i < len(str); {
		j := i + 1

		for ; j < len(str); j++ {
			if isDigit(str[i]) != isDigit(str[j]) {
				break
			}
		}
		processedStr = append(processedStr, str[i:j])

		i = j
	}
	return processedStr
}

// NaturalLess03 compare string by separating number and string using regexp
func NaturalLess03(str1, str2 string) bool {
	pStr1 := processWithRegexp(str1)
	pStr2 := processWithRegexp(str2)
	minLen := int(math.Min(float64(len(pStr1)), float64(len(pStr2))))
	for i := 0; i < minLen; i++ {
		e1 := pStr1[i]
		e2 := pStr2[i]
		if e1 == e2 {
			continue
		}
		ei1, err1 := strconv.Atoi(e1)
		ei2, err2 := strconv.Atoi(e2)

		switch {
		case err1 == nil && err2 == nil && ei1 != ei2:
			return ei1 < ei2
		case err1 != nil && err2 != nil:
			return e1 < e2
		case err1 != nil:
			return true
		case err2 != nil:
			return false
		}
	}
	return len(str1) < len(str2)
}

func processWithRegexp(str string) []string {
	reg := regexp.MustCompile(`\d+|\D+`)
	return reg.FindAllString(str, -1)
}

// NaturalLess01 compare string byte by byte
func NaturalLess01(str1, str2 string) bool {
	idx1, idx2 := 0, 0
	for idx1 < len(str1) && idx2 < len(str2) {
		c1, c2 := str1[idx1], str2[idx2]
		dig1, dig2 := isDigit(c1), isDigit(c2)

		switch {
		case dig1 != dig2:
			return dig1
		case !dig1:
			if c1 != c2 {
				return c1 < c2
			}
			idx1++
			idx2++
		default:
			for ; idx1 < len(str1) && str1[idx1] == '0'; idx1++ {
			}
			for ; idx2 < len(str2) && str2[idx2] == '0'; idx2++ {
			}
			nonZero1, nonZero2 := idx1, idx2
			for ; idx1 < len(str1) && isDigit(str1[idx1]); idx1++ {
			}
			for ; idx2 < len(str2) && isDigit(str2[idx2]); idx2++ {
			}
			if len1, len2 := idx1-nonZero1, idx2-nonZero2; len1 != len2 {
				return len1 < len2
			}
			if nr1, nr2 := str1[nonZero1:idx1], str2[nonZero2:idx2]; nr1 != nr2 {
				return nr1 < nr2
			}
			if nonZero1 != nonZero2 {
				return nonZero1 < nonZero2
			}
		}
	}
	return len(str1) < len(str2)
}

//func main() {
//	strs := []string{
//		"ab", "abc1",
//		"abc01", "abc2",
//		"abc5", "abc10",
//	}
//	fmt.Printf(PrintFormat, "Origin", strs)
//	var strs00, strs01, strs02, strs03 []string
//	strs00 = make([]string, len(strs))
//	strs01 = make([]string, len(strs))
//	strs02 = make([]string, len(strs))
//	strs03 = make([]string, len(strs))
//
//	copy(strs00, strs)
//	sort.Strings(strs00)
//	fmt.Printf(PrintFormat, "Sort by stdsorter", strs00)
//
//	copy(strs01, strs)
//	sort.Slice(strs01, func(i, j int) bool {
//		return NaturalLess01(strs01[i], strs01[j])
//	})
//	fmt.Printf(PrintFormat, "Sort by NaturalLess01", strs01)
//
//	copy(strs02, strs)
//	sort.Slice(strs02, func(i, j int) bool {
//		return NaturalLess02(strs02[i], strs02[j])
//	})
//	fmt.Printf(PrintFormat, "Sort by NaturalLess02", strs02)
//
//	copy(strs03, strs)
//	sort.Slice(strs03, func(i, j int) bool {
//		return NaturalLess03(strs03[i], strs03[j])
//	})
//	fmt.Printf(PrintFormat, "Sort by NaturalLess03", strs03)
//}
