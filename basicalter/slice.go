package basicalter

import "sort"

// UniqueInSlice remove duplicate elements in slice
// and return the result with a new slice.
func UniqueInSlice[T comparable, S ~[]T](list S) S {
	switch {
	case list == nil:
		return nil
	case len(list) == 0:
		return make(S, 0, cap(list))
	case len(list) == 1:
		r := make(S, 1, cap(list))
		r[0] = list[0]

		return r
	default:
		k := make(map[T]struct{}, len(list))
		r := make(S, 0, cap(list))

		for _, v := range list {
			if _, ok := k[v]; !ok {
				k[v] = struct{}{}
				r = append(r, v)
			}
		}

		return r
	}
}

// DelEmptyStrings remove empty elements in slice of string
// and return the result with a new slice.
func DelEmptyStrings(list []string) []string {
	return DelInSlice("", list)
}

// DelInSlice remove all occurrence of an element in a slice
// and return the result with a new slice.
func DelInSlice[T comparable, S ~[]T](elem T, list S) S {
	return FilterInSliceWith(list, func(e T) bool {
		return e != elem
	})
}

// FilterInSliceWith generate a new slice
// by applying on an input slice 'list' a function 'filter'
// to determine the inclusion of each element.
func FilterInSliceWith[T any, S ~[]T](list S, filter func(T) bool) S {
	if list == nil {
		return nil
	}
	r := make(S, 0, cap(list))
	for _, v := range list {
		if filter(v) {
			r = append(r, v)
		}
	}

	return r
}

// ReverseSlice reverse order of a slice last to first, before last to second, etc.
func ReverseSlice[T any](list []T) {
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

// ReplaceInSliceWith replace each element of a slice
// with the result of the function 'replace' passed in arguments.
func ReplaceInSliceWith[T any](list []T, replace func(T) T) {
	for i := 0; i < len(list); i++ {
		list[i] = replace(list[i])
	}
}
