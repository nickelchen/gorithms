package sort

func QuickSort(numbers []int) {
	if len(numbers) <= 1 {
		return
	}
	pivot := numbers[0]
	head, tail := 0, len(numbers)-1
	i := 1
	for i <= tail {
		if numbers[i] > pivot {
			numbers[i], numbers[tail] = numbers[tail], numbers[i]
			tail--
		} else {
			numbers[i], numbers[head] = numbers[head], numbers[i]
			head++
			i++
		}
	}
	QuickSort(numbers[:head])
	QuickSort(numbers[head+1:])
}

func InsertSort(numbers []int) {
	tail := len(numbers) - 1

	for i := 0; i < tail; i++ {
		for j := i + 1; j > 0; j-- {
			if numbers[j-1] < numbers[j] {
				break
			} else {
				numbers[j-1], numbers[j] = numbers[j], numbers[j-1]
			}
		}
	}
}

func BubbleSort(numbers []int) {
	tail := len(numbers) - 1

	for i := tail; i >= 0; i-- {
		swapped := false
		for j := 0; j < i; j++ {
			if numbers[j] >= numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func MergeSort(numbers []int) []int {
	if len(numbers) <= 1 { // only 1 element.
		return numbers
	}

	mid := len(numbers) / 2
	s1 := MergeSort(numbers[:mid]) // now s1 is sorted
	s2 := MergeSort(numbers[mid:]) // now s2 is sorted
	s3 := merge(s1, s2)
	return s3
}

func merge(s1, s2 []int) []int {
	var l int = len(s1) + len(s2) // total length.
	var res []int = make([]int, l, l)
	i := 0
	j := 0
	k := 0
	for k < l {
		if i >= len(s1) { // s1 is exsausted, pick from s2
			res[k] = s2[j]
			j++
		} else if j >= len(s2) { // s2 is exsausted, pick from s1
			res[k] = s1[i]
			i++
		} else if s1[i] >= s2[j] {
			res[k] = s2[j]
			j++
		} else {
			res[k] = s1[i]
			i++
		}
		k++
	}
	return res
}
