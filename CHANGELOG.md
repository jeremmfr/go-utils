# changelog

* golang 1.18 is now the minimum version
* add function `OneOfSliceWith` with generic parameters in package `basiccheck` to replace `OneOfStringsWith` function which is now deprecated
* add function `InSlice` with generic parameters in package `basiccheck` to replace `StringInSlice`, `IntInSlice`, `Int64InSlice` functions which are now deprecated
* add function `EqualSlice` with generic parameters in package `basiccheck` to replace `EqualStringSlice` function which is now deprecated
* add function `AllInSliceWith` with generic parameters in package `basiccheck` to replace `AllStringsWith` function which is now deprecated

## v0.4.1

* optimize `UniqueStrings` (reduce memory usage)

## v0.4.0

* add functions `OneOfStringsWith` and `AllStringsWith` in package `basiccheck`
* add functions `FilterStringsWith` and `ReplaceStringsWith` in package `basicalter`

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
