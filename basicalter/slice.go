package basicalter

import "sort"

// UniqueStrings remove duplicate string in slice of string.
func UniqueStrings(list []string) []string {
	k := make(map[string]struct{}, len(list))
	r := []string{}
	for _, v := range list {
		if _, ok := k[v]; !ok {
			k[v] = struct{}{}
			r = append(r, v)
		}
	}

	return r
}

// DelEmptyStrings remove empty elements in slice of string.
func DelEmptyStrings(list []string) []string {
	return DelStringInSlice("", list)
}

// DelStringInSlice remove all occurrence of a string in slice of string.
func DelStringInSlice(str string, list []string) []string {
	return FilterStringsWith(list, func(s string) bool {
		return s != str
	})
}

// FilterStringsWith generate a new slice of string
// by applying on 'list' a function 'filter' to determine the inclusion of each element.
func FilterStringsWith(list []string, filter func(string) bool) []string {
	r := make([]string, 0)
	for _, v := range list {
		if filter(v) {
			r = append(r, v)
		}
	}

	return r
}

// ReverseStrings reverse order of string slice last to first, before last to second, etc.
func ReverseStrings(list []string) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}

// SortStringsByLengthInc sort slice of string by length increasing before lexicographic order.
func SortStringsByLengthInc(list []string) {
	sort.Sort(sortStringsLengthInc(list))
}

type sortStringsLengthInc []string

func (s sortStringsLengthInc) Len() int {
	return len(s)
}

func (s sortStringsLengthInc) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortStringsLengthInc) Less(i, j int) bool {
	if len(s[i]) < len(s[j]) {
		return true
	}
	if len(s[j]) < len(s[i]) {
		return false
	}

	return s[i] < s[j]
}

// SortStringsByLengthDec sort slice of string by length decreasing before lexicographic order.
func SortStringsByLengthDec(list []string) {
	sort.Sort(sortStringsLengthDec(list))
}

type sortStringsLengthDec []string

func (s sortStringsLengthDec) Len() int {
	return len(s)
}

func (s sortStringsLengthDec) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortStringsLengthDec) Less(i, j int) bool {
	if len(s[i]) > len(s[j]) {
		return true
	}
	if len(s[j]) > len(s[i]) {
		return false
	}

	return s[i] < s[j]
}

// ReplaceStringsWith replace each string of a slice
// with the result of the function 'replace' passed in arguments.
func ReplaceStringsWith(list []string, replace func(string) string) {
	for i := 0; i < len(list); i++ {
		list[i] = replace(list[i])
	}
}
