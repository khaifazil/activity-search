package main

import "fmt"

const size int = 1000

func sequentialSearchUnsorted(n int, arr [size]int, target int) (int, int) { // unsorted

	var count int = 0
	for i := 0; i < n; i++ {
		count++
		if arr[i] == target {
			return i, count
		}
	}
	return -1, count
}

//Iterative Sequential search on sorted array
func sequentialSearchSorted(n int, arr [size]int, target int) (int, int) { // sorted

	var count int = 0
	for i := 0; i < n; i++ {
		count++
		if arr[i] == target {
			return i, count
		} else if arr[i] > target {
			return -1, count
		}
	}
	return -1, count
}

//Recursive sequential search on unsorted array
func recursiveSeqSearch(size int, start int, arr [size]int, target int,) int{
	if start > size-1{// search not found
		return -1
	}else {

		if target == arr[start]{
			return start
		}else {
			return recursiveSeqSearch(size, start +1, arr, target)
		}
	}
}

//Recursive sequential search on sorted array
func recursiveSeqSearchSorted(size int, start int, arr [size]int, target int, ) int{
	if start > size-1{// search not found
		return -1
	}else {

		if target == arr[start]{
			return start
		}else {
			if target > arr[start] { //if ascending order
				return recursiveSeqSearch(size, start +1, arr, target)
			}else {
				return -1
			}
		}
	}
}

func recursiveSeq(arr [size]int, target int)int{
	return recursiveSeqSearchSorted(len(arr), 0, arr, target)
}

func binarySearch(size int, arr [size]int, target int) (int, int) {
	first:=0
	last:= size-1
	count := 0

	for first <= last{
		count++
		mid := (first + last)/2
		if target == arr[mid]{
			return mid, count
		}else if target < arr[mid]{
			last = mid - 1
		}else{
			first = mid + 1
		}
	}
	return -1, count
}

func recursiveBinarySearch(size int, first int, last int, arr[size]int, target int)int{
	
	if first > last{
		return -1
	}else{
		mid := (first + last)/2

		if arr[mid] == target{
			return mid
		}else{
			if target < arr[mid]{
				return recursiveBinarySearch(size, first, mid-1, arr, target)
			}else{
				return recursiveBinarySearch(size, mid+1, last, arr, target)
			}
		}
	}
}

func recursiveBinary(arr [size]int, target int)int{
	return recursiveBinarySearch(len(arr), 0, len(arr)-1, arr, target)
}

func main() {

	arr := [size]int{}
	for i:= 0; i<len(arr); i++{
		arr[i] = (i+1)*2
	}

fmt.Println(arr)
	var target int = 20
	result, count := sequentialSearchUnsorted(size, arr, target)
	// result2, count2 := binarySearch(size, arr, target)

	if result == -1 {
		fmt.Printf("Target %+v is not found", target)
	} else {
		fmt.Printf("Target %+v is found at position %+v.", target, result+1)
		fmt.Printf("\n%v number of comparisons done", count)
	}
}
