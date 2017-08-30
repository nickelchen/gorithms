package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func initNumbers() []int {
	return []int{1, 100, -1, -20, 8, 11, 9, 10, -2}
}

func sortedNumbers() []int {
	return []int{-20, -2, -1, 1, 8, 9, 10, 11, 100}
}

func verify(t *testing.T, n []int) {
	s := sortedNumbers()
	if !reflect.DeepEqual(n, s) {
		t.Fatal(fmt.Sprintf("not equal. want %v, get %v", s, n))
	}
}

func TestQuickSort(t *testing.T) {
	numbers := initNumbers()
	QuickSort(numbers)
	verify(t, numbers)
}

func TestInsertSort(t *testing.T) {
	numbers := initNumbers()
	InsertSort(numbers)
	verify(t, numbers)
}

func TestShellSort(t *testing.T) {
	numbers := initNumbers()
	ShellSort(numbers)
	verify(t, numbers)
}

func TestSelectSort(t *testing.T) {
	numbers := initNumbers()
	SelectSort(numbers)
	verify(t, numbers)
}

func TestBubbleSort(t *testing.T) {
	numbers := initNumbers()
	BubbleSort(numbers)
	verify(t, numbers)
}

func TestMergeSort(t *testing.T) {
	numbers := initNumbers()
	numbers = MergeSort(numbers)
	verify(t, numbers)
}
