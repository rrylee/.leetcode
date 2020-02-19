package main

func main() {

}

func MergeSort(nums []int) {
	sort(nums, 0, len(nums)-1)
}

func sort(nums []int, l, r int) {
	if l == r {
		return
	}
	mid := l + (r-l)>>1
	sort(nums, l, mid)
	sort(nums, mid+1, r)
	merge(arr, l, mid, r)
}

func merge(nums []int, l, mid, r int) {
	temp := make([]int, r-l+1, r-l+1)
}
