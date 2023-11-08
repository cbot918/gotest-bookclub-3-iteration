package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

//
//
// Array & Hashing
//
//
/* 217. Contains Duplicate */
func containsDuplicate(nums []int) bool {
	s := make(map[int]bool)
	for _, num := range nums {
		s[num] = true
	}
	return len(s) < len(nums)
}

/* 242. Valid Anagram */
func isAnagram(s string, t string) bool {
	count := make(map[rune]int)
	for _, c := range s {
		count[c] += 1
	}
	for _, c := range t {
		count[c] -= 1
	}
	for _, c := range count {
		if c != 0 {
			return false
		}
	}
	return true
}

/* 1. Two Sum */
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, num := range nums {
		val, ok := m[target-num]
		if ok {
			return []int{val, i}
		}
		m[num] = i
	}
	return nil
}

/* 49. Group Anagrams */
func groupAnagrams(strs []string) [][]string {
	var sort = func(str string) string {
		s := strings.Split(str, "")
		sort.Strings(s)
		return strings.Join(s, "")
	}

	t := make(map[string][]string)
	for _, str := range strs {
		key := sort(str)
		if _, ok := t[key]; !ok {
			t[key] = []string{}
		}
		t[key] = append(t[key], str)
	}

	res := make([][]string, len(t))
	i := 0
	for _, item := range t {
		res[i] = item
		i += 1
	}
	return res
}

/* 347. Top K Frequent Elements */
func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num] += 1
	}

	keys := []int{}
	for key := range count {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i int, j int) bool {
		return count[keys[i]] > count[keys[j]]
	})

	i := 0
	res := []int{}
	for i < k {
		res = append(res, keys[i])
		i += 1
	}

	return res
}

/* 271. Encode and Decode Strings */
type Codec struct{}

func (codec *Codec) Encode(strs []string) string {
	_str := make([]string, len(strs))
	copy(_str, strs)
	res := ""
	for _, str := range _str {
		res += strconv.Itoa(len(str)) + "#" + str
	}
	return res
}
func (codec *Codec) Decode(strs string) []string {

	index := 0
	res := []string{}
	for index < len(strs) {
		start := index
		end := index
		num := []byte("")
		for strs[index] != '#' {
			num = append(num, strs[index])
			start += 1
			end += 1
			index += 1
		}
		numInt, _ := strconv.Atoi(string(num))

		start += 1
		end += numInt + 1

		res = append(res, strs[start:end])

		index = end

	}
	return res
}

/* 238. Product of Array Except Self */
func lc238_productExceptSelf(nums []int) []int {

	ans := make([]int, len(nums))

	prefix := 1
	for i, num := range nums {
		ans[i] = prefix
		prefix *= num
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] *= postfix
		postfix *= nums[i]
	}

	return ans
}

/* 128. Longest Consecutive Sequence */
func longestConsecutive(nums []int) int {

	set := map[int]bool{}

	for _, num := range nums {
		set[num] = true
	}

	res := 1

	for _, num := range nums {
		// 會每一個跑過, 所以看右邊就好
		if set[num-1] {
			continue
		}

		current := 1
		num := num + 1

		for set[num] {
			current += 1
			num += 1
		}

		if current > res {
			res = current
		}

	}

	return res
}

//
//
// Two Pointers
//
//
/* 125. Valid Palindrome */
func isPalindrome(s string) bool {

	ans := true
	bstr := []byte(s)

	str := []byte{}
	for _, ch := range bstr {
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9' {
			str = append(str, ch)
		}
	}
	str = []byte(strings.ToLower(string(str)))

	i := 0
	j := len(str) - 1

	for i <= j {
		if str[i] != str[j] {
			ans = false
		}
		i++
		j--
	}

	return ans
}

/* 15. 3Sum */
func threeSum(nums []int) [][]int {
	fmt.Println(nums)
	// newNums := nums[1:]
	res := [][]int{}

	for _, num := range nums {

		// newNums := nums[i+1:]
		target := -(num)
		fmt.Println("num: ", num)
		// fmt.Println("newNums: ", newNums)
		m := map[int]int{}
		for idx, number := range nums {

			// fmt.Println(number)
			_, ok := m[target-number]
			if ok {
				res = append(res, []int{num, target - number, number})
				break
			}
			m[number] = idx
		}

	}

	fmt.Println(res)

	return [][]int{}

}

func main() {

	{ //Array & Hashing

		// { /* 217. Contains Duplicate */
		// 	nums := []int{1, 2, 3, 1}
		// 	ans := true
		// 	res := containsDuplicate(nums)
		// 	fmt.Println(reflect.DeepEqual(res, ans))
		// }

		// { /* 242. Valid Anagram */
		// 	s1 := "anagram"
		// 	s2 := "nagaram"
		// 	ans := true
		// 	res := isAnagram(s1, s2)
		// 	fmt.Println(ans == res)
		// }

		// { /* 1. Two Sum */
		// 	nums := []int{2, 7, 11, 15}
		// 	target := 9
		// 	ans := []int{0, 1}
		// 	// nums := []int{3, 2, 4}
		// 	// target := 6
		// 	// ans := []int{1, 2}
		// 	// nums := []int{3, 3}
		// 	// target := 6
		// 	// ans := []int{0, 1}
		// 	ret := twoSum(nums, target)
		// 	fmt.Println(ret)
		// 	fmt.Println(reflect.DeepEqual(ret, ans))
		// }
		// 	nums := []int{2, 7, 11, 15}
		// 	target := 9
		// 	ans := []int{0, 1}
		// 	res := twoSum(nums, target)
		// 	fmt.Println(reflect.DeepEqual(res, ans))
		// }

		// { /* 49. Group Anagrams */
		// 	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
		// 	ans := [][]string{[]string{"bat"}, []string{"nat", "tan"}, []string{"ate", "eat", "tea"}}
		// 	res := groupAnagrams(strs)
		// 	fmt.Println(reflect.DeepEqual(res, ans))
		// }

		// { /* 347. Top K Frequent Elements */
		// 	nums := []int{1, 1, 1, 2, 2, 3}
		// 	k := 2
		// 	ans := []int{1, 2}
		// 	res := topKFrequent(nums, k)
		// 	fmt.Println(reflect.DeepEqual(res, ans))
		// }

		// { /* 271. Encode and Decode Strings */
		// 	source := []string{"Helloasdfasdfasdfasdf", "World"}
		// 	c := Codec{}
		// 	// fmt.Println(reflect.DeepEqual(input, c.Decode(c.Encode(input))))
		// 	fmt.Println("source", source)
		// 	encoded := c.Encode(source)
		// 	fmt.Println("encoded: ", encoded)
		// 	decoded := c.Decode(encoded)
		// 	fmt.Println("decoded: ", decoded)
		// 	fmt.Println(reflect.DeepEqual(source, decoded))
		// }

		// { /* 238. Product of Array Except Self */
		// 	nums := []int{1, 2, 3, 4}
		// 	ans := []int{24, 12, 8, 6}
		// 	res := lc238_productExceptSelf(nums)
		// 	fmt.Println(reflect.DeepEqual(res, ans))
		// }

		// { /* 128. Longest Consecutive Sequence */
		// 	nums := []int{100, 4, 200, 1, 3, 2}
		// 	ans := 4
		// 	res := longestConsecutive(nums)
		// 	fmt.Println(reflect.DeepEqual(res, ans))
		// }

	}

	{ // Two Pointers

		// { /* 125. Valid Palindrome */
		// 	s := "A man, a plan, a canal: Panama"
		// 	// s := "0P"
		// 	ans := true
		// 	res := isPalindrome(s)
		// 	fmt.Println(ans == res)
		// }

		{ /* 15. 3Sum */
			nums := []int{-1, 0, 1, 2, -1, -4}
			ans := [][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}}
			ret := threeSum(nums)
			fmt.Println(reflect.DeepEqual(ret, ans))
		}

	}

}
