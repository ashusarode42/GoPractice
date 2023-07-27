package main

import (
	"fmt"
	"reflect"
)

// Define a sortable interface
type Sortable interface {
	Len() int
	Less(int, int) bool
	Swap(int, int)
}

//func BubbleSort(items []type, desc int) (sortedItems []type, err error)

func BubbleSort(items Sortable) {
	length := items.Len()
	for i := 0; i < length; i++ {
		for j := i; j < length-i-1; j++ {
			if items.Less(j, j+1) {
				items.Swap(j, j+1)
			}
		}
	}
}

// Define the IntArr type
type IntArray []int

// Provide Len method to items
func (items IntArray) Len() int {
	return len(items)
}

// Provide the Less method to items
func (items IntArray) Less(i int, j int) bool {
	return items[i] < items[j]
}

// Provide Swap method to items
func (items IntArray) Swap(i int, j int) {
	items[i], items[j] = items[j], items[i]
}
func main() {
	itemsTosort := IntArray{2, 3, 1, -9, 0}
	//Call the sort method
	BubbleSort(itemsTosort)
	fmt.Printf("Sorted int array is: %v\n", itemsTosort)
}
func TestBubblesort() {

	testCases := []struct {
		name            string
		array, expected []int
		sortable        []Sortable
	}{
		{
			name:  "already Sorted",
			array: []int{1, 2, 4, 5, 8},

			expected: []int{1, 2, 4, 5, 8},
		},
		{
			name:  "not sorted",
			array: []int{5, 1, 4, 2, 8},

			expected: []int{1, 2, 4, 5, 8},
		},
	}
	for _, v := range testCases {
		sortedArray := BubbleSort(v.array, Sortable)
		if !reflect.DeepEqual(sortedArray, v.expected) {
			fmt.Println("Not getting expected result!")
		}
		fmt.Println("Test case passed")
	}

}
