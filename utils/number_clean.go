package utils

import (
	"fmt"
	"strconv"
	"strings"
)

var ones = map[string]int{"zero": 0, "o": 0, "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

var teens = map[string]int{"eleven": 11, "twelve": 12, "thirteen": 13, "fourteen": 14, "fifteen": 15, "sixteen": 16, "seventeen": 17, "eighteen": 18, "nineteen": 19}

var tens = map[string]int{"ten": 10, "twenty": 20, "thirty": 30, "forty": 40, "fifty": 50, "sixty": 60, "seventy": 70, "eighty": 80, "ninety": 90}

var hundreds = map[string]int{"hundred": 100}

var thusands = map[string]int{"thousand": 1000}

type slicedInt struct {
	start int
	end   int
	num   string
}

func WordToNum(s string) string {
	var num []int
	var start int
	var end int
	var cleanedNums []slicedInt
	start = -1
	sSlice := strings.Split(s, " ")
	for i, word := range sSlice {
		if start == -1 {
			start = i
		}
		if v, found := ones[word]; found {
			num = append(num, v)
			continue
		}
		if v, found := teens[word]; found {
			num = append(num, v)
			continue
		}
		if v, found := tens[word]; found {
			num = append(num, v)
			continue
		}
		if v, found := hundreds[word]; found {
			num = append(num, v)
			continue
		}
		if v, found := thusands[word]; found {
			num = append(num, v)
			continue
		}
		if len(num) > 0 {
			sNum := intSliceToNum(num)
			end = i - 1
			cleanedNums = append(cleanedNums, slicedInt{
				start: start,
				end:   end,
				num:   sNum,
			})
			num = []int{}
		}
		start = -1
	}
	if len(num) != 0 {
		sNum := intSliceToNum(num)
		end = len(sSlice) - 1
		cleanedNums = append(cleanedNums, slicedInt{
			start: start,
			end:   end,
			num:   sNum,
		})
	}
	newS := newSlice(sSlice, cleanedNums)

	return newS
}

func intSliceToNum(slice []int) string {
	var totalNum int
	var thusandsTotal int
	var hundredsTotal int
	for i := 0; i < len(slice); i++ {
		if slice[i] == 100 {
			totalNum = totalNum * 100
			hundredsTotal = totalNum
			totalNum = 0
		} else if slice[i] == 1000 {
			totalNum = totalNum * 1000
			thusandsTotal = totalNum
			totalNum = 0
		} else if slice[i] < 10 && i != len(slice)-1 {
			tempNum := strconv.Itoa(slice[i])
			j := 1
			for {
				if slice[i+j] < 10 {
					if slice[i] > 10 {
						convInt, _ := strconv.Atoi(tempNum)
						add := convInt + slice[i+j]
						tempNum = strconv.Itoa(add)
					} else {
						tempNum = tempNum + strconv.Itoa(slice[i+j])
					}
					i++
					if i+j != len(slice)-1 {
						break
					}
					continue
				}
				if slice[i+j] > 10 && slice[i+j] < 100 {
					convInt, _ := strconv.Atoi(tempNum)
					add := convInt*100 + slice[i+j]
					tempNum = strconv.Itoa(add)
					i++
					if i+j != len(slice)-1 {
						break
					}
					continue
				}
				break
			}
			convInt, _ := strconv.Atoi(tempNum)
			totalNum += convInt
		} else {
			totalNum += slice[i]
		}
	}
	return strconv.Itoa(thusandsTotal + hundredsTotal + totalNum)
}

func newSlice(originalS []string, s []slicedInt) string {
	for _, sliced := range s {
		startFlag := true
		for i := sliced.start; i <= sliced.end; i++ {
			if startFlag {
				originalS[i] = sliced.num
				startFlag = false
			} else {
				originalS[i] = ""
			}
		}
	}
	var trimedSlice []string
	var skipNext bool
	for _, so := range originalS {
		if strings.TrimSpace(so) != "" {
			if so == "point" || so == "decimal" {
				skipNext = true
			} else if skipNext {
				decimalPointNum := trimedSlice[len(trimedSlice)-1]
				trimedSlice[len(trimedSlice)-1] = fmt.Sprintf("%s.%s", decimalPointNum, so)
				skipNext = false
			} else if !skipNext {
				trimedSlice = append(trimedSlice, so)
			}
		}
	}
	return strings.Join(trimedSlice, " ")
}
