type DynamicArray struct {
    capacity int
    size int
    array []int
}

func NewDynamicArray(capacity int) *DynamicArray {
	if capacity <= 0 {
		capacity = 1
	}

	return &DynamicArray{
		capacity: capacity,
		size:     0,
		array:    make([]int, capacity),
	}
}

func (da *DynamicArray) Get(i int) int {
	if i < 0 || i >= da.size {
		panic("index out of bounds")
	}

	return da.array[i]
}

func (da *DynamicArray) Set(i int, n int) {
    da.array[i] = n
     
}

func (da *DynamicArray) Pushback(n int) {
    if da.size == da.capacity {
		da.resize()
	}

	// Insert at the next free index
	da.array[da.size] = n
	da.size++
}

func (da *DynamicArray) Popback() int {
	if da.size == 0 {
		panic("array is empty")
	}

	da.size--
	return da.array[da.size]
}

func (da *DynamicArray) resize() {
	newCapacity := da.capacity * 2
	newArray := make([]int, newCapacity)

	// Copy existing elements
	for i := 0; i < da.size; i++ {
		newArray[i] = da.array[i]
	}

	da.array = newArray
	da.capacity = newCapacity
}

func (da *DynamicArray) GetSize() int {
    return da.size

}

func (da *DynamicArray) GetCapacity() int {
    return da.capacity
}
