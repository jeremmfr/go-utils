# changelog

## v0.7.0

* :warning: **BREAKING CHANGES**: remove deprecated functions in `basicalter`: `UniqueStrings`, `DelStringInSlice`, `FilterStringsWith`, `ReverseStrings`, `ReplaceStringsWith`
* add new `basicnew` package with this functions:
  * `MapKeys`: MapKeys generate a new slice with all keys of a map.
  * `MapValues`: MapValues generate a new slice with all values of a map.
* add `MergeMaps` function in `basicalter` package (copies all key/value pairs in maps to a map)

## v0.6.0

* :warning: **BREAKING CHANGES**: remove deprecated functions in `basiccheck`: `StringInSlice`, `EqualStringSlice`, `IntInSlice`, `Int64InSlice`, `OneOfStringsWith`, `AllStringsWith`
* add `UniqueInSlice` function with generic parameters in `basicalter` package to replace `UniqueStrings` function which is now deprecated
* add `DelInSlice` function with generic parameters in `basicalter` package to replace `DelStringInSlice` function which is now deprecated
* add `FilterInSliceWith` function with generic parameters in `basicalter` package to replace `FilterStringsWith` function which is now deprecated
* add `ReverseSlice` function with generic parameters in `basicalter` package to replace `ReverseStrings` function which is now deprecated
* add `ReplaceInSliceWith` function with generic parameters in `basicalter` package to replace `ReplaceStringsWith` function which is now deprecated

## v0.5.0

* golang 1.18 is now the minimum version
* add `OneInSliceWith` function with generic parameters in `basiccheck` package to replace `OneOfStringsWith` function which is now deprecated
* add `InSlice` function with generic parameters in `basiccheck` package to replace `StringInSlice`, `IntInSlice`, `Int64InSlice` functions which are now deprecated
* add `EqualSlice` function with generic parameters in `basiccheck` package to replace `EqualStringSlice` function which is now deprecated
* add `AllInSliceWith` function with generic parameters in `basiccheck` package to replace `AllStringsWith` function which is now deprecated

## v0.4.1

* optimize `UniqueStrings` (reduce memory usage)

## v0.4.0

* add `OneOfStringsWith` and `AllStringsWith` functions in `basiccheck` package
* add `FilterStringsWith` and `ReplaceStringsWith` functions in `basicalter` package

## v0.3.1

* optimize `AbsoluteInt` (with Two's complement implementation)
* panic if call `AbsoluteInt` with the minimum value of integers (the absolute of the minimum value of integers is impossible to return)

## v0.3.0

* simplify `ReverseStrings` to reverse original slice

## v0.2.0

* add `ReverseStrings` function
* rename `CheckStringHasOneOfPrefixes` to `StringHasOneOfPrefixes`

## v0.1.1

* fix useless return slice for SortStringsByLength functions
* fix missing function comments and variable name

## v0.1.0

First release
