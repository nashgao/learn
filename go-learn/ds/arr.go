package ds


// the size of the array is fixed, it can not be modified once it's defined (the length of the array will always equals to the size of the array)
// golang would only consider multiple arrays as identical when: 1. size of the arrays are the same. 2. the type of the storage are the same

// array initialization: 
//   1. arr := [3]int{1,2,3} explicitly identifies the size of the array
//   2. arr := [...]int{1,2,3} implicity indentifies identifies the size of the array, golang compiler will derive the size

// if (elements of array are literals)
// if (size of array => 4):
//     array => stack
// else :
//     array => static area (will be extracted during the run time)
func TestArr()  {
	// explicitly declare array
	arrEx := [5]int{1,2,3,4,5}

	// implicitly declare array
	arrIm := [...]int{1,2,3}
}


// the size of the slice is dynamic, it will expand when size is not enough for storing
// the difference of array and slice from the grammar perspective is that, you don't define the size when initializing the slice and vice versa

// slice definition 
//   1. []int
//   2. []interface{}
func TestSlice() {

}