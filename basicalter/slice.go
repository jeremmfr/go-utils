package basicalter

import "sort"

// UniqueStrings remove duplicate string in slice of string.
func UniqueStrings(s []string) []string {
	k := make(map[string]bool)
	r := []string{}
	for _, v := range s {
		if _, vv := k[v]; !vv {
			k[v] = true
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
	r := make([]string, 0)
	for _, v := range list {
		if v != str {
			r = append(r, v)
		}
	}

	return r
}

// SortStringsByLengthInc sort slice of string by length increasing before lexicographic order.
func SortStringsByLengthInc(s []string) []string {
	sort.Sort(sortStringsLengthInc(s))

	return s
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
func SortStringsByLengthDec(s []string) []string {
	sort.Sort(sortStringsLengthDec(s))

	return s
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
