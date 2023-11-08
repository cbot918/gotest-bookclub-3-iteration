package main

import (
	"fmt"
	"reflect"
	"testing"
)

/* 217. Contains Duplicate */
func TestContainsDuplicate(t *testing.T) {

	nums := []int{1, 2, 3, 1}
	got := containsDuplicate(nums)
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	} else {
		fmt.Println("217. Contains Duplicate")
		fmt.Println("data: ", nums)
		fmt.Println("result: ", want)
	}
}

/* 242. Valid Anagram */
func TestIsAnagram(t *testing.T) {

	s1 := "anagram"
	s2 := "nagaram"
	got := isAnagram(s1, s2)
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	} else {
		fmt.Println("242. Valid Anagram")
		fmt.Println("s1: ", s1)
		fmt.Println("s2: ", s2)
		fmt.Println("result: ", want)
	}
}

/* 1. Two Sum */
func TestTwoSum(t *testing.T) {

	testcases := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "case1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "case2",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "case3",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}

	for _, cas := range testcases {
		t.Run(cas.name, func(t *testing.T) {
			got := twoSum(cas.nums, cas.target)
			want := cas.want

			if !reflect.DeepEqual(got, want) {
				t.Errorf("want %v got %v", want, got)
			}

		})
	}

}
