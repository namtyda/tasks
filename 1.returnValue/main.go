package main

import (
	"sort"
	"strconv"
)

func main() {
	GetMapValuesSortedByKey(map[int]string{
		9:  "Сентябрь",
		1:  "Январь",
		2:  "Февраль",
		10: "Октябрь",
		5:  "Май",
		7:  "Июль",
		8:  "Август",
		12: "Декарь",
		3:  "Март",
		6:  "Июнь",
		4:  "Апрель",
		11: "Ноябрь",
	})
}

//ReturnInt func
func ReturnInt() int {
	return 1
}

//ReturnFloat func
func ReturnFloat() float32 {
	return 1.1
}

//ReturnIntArray func
func ReturnIntArray() [3]int {
	return [3]int{1, 3, 4}
}

//ReturnIntSlice func
func ReturnIntSlice() []int {
	return []int{1, 2, 3}
}

//IntSliceToString func
func IntSliceToString(sl []int) (result string) {
	for _, number := range sl {
		str := strconv.FormatInt(int64(number), 10)
		result += str
	}
	return result
}

//MergeSlices func
func MergeSlices(slF []float32, slInt []int32) []int {
	result := make([]int, 0, len(slF)+len(slInt))
	for _, value := range slF {
		result = append(result, int(value))
	}
	for _, value := range slInt {
		result = append(result, int(value))
	}
	return result
}

//GetMapValuesSortedByKey func
func GetMapValuesSortedByKey(mm map[int]string) []string {
	result := make([]string, 0, len(mm))
	keys := make([]int, 0, len(mm))
	for key := range mm {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, value := range keys {
		result = append(result, mm[value])
	}
	return result
}
