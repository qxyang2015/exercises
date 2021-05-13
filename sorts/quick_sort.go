package Sort

/*
算法步骤
	从数列中挑出一个元素，称为 "基准"（pivot）;
	重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
	递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序；
*/

//快速排序
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	quickSortMethod(arr, 0, len(arr)-1)
	return arr
}

func quickSortMethod(arr []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(arr, l, r)
	quickSortMethod(arr, l, p-1)
	quickSortMethod(arr, p+1, r)
}

func partition(arr []int, l, r int) int {
	counter, p := l, r
	for i := l; i < r; i++ {
		if arr[i] < arr[p] {
			arr[counter], arr[i] = arr[i], arr[counter]
			counter++
		}
	}
	arr[p], arr[counter] = arr[counter], arr[p]
	return counter
}
