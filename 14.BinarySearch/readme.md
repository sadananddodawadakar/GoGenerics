## GoGenerics
This repository contains examples of using generics in Go. The specific example provided here is an implementation of a generic binary search function.

## Binary Search
This example demonstrates how to implement a binary search using generics in Go. The BinarySearch function can operate on any ordered type.

## Implementation
The implementation consists of the following function:

BinarySearch[T constraints.Ordered]: A generic function that performs a binary search on a sorted slice of type T, where T is constrained to ordered types.
Usage
The BinarySearch function can be used to find the index of a target value in a sorted slice. If the target value is not found, the function returns -1.

## Example
Here is an example of how to use the binary search function:

Go
package main

import (
    "fmt"

    "golang.org/x/exp/constraints"
)

func BinarySearch[T constraints.Ordered](data []T, target T) int {
    low, high := 0, len(data)-1
    for low <= high {
        mid := (low + high) / 2
        if data[mid] < target {
            low = mid + 1
        } else if data[mid] > target {
            high = mid - 1
        } else {
            return mid
        }
    }
    return -1
}

func main() {
    nums := []int{1, 3, 5, 7, 9}
    fmt.Println(BinarySearch(nums, 5)) // 2
}
## Running the Example
To run the example, simply execute the following command:

go run main.go
This will output:

2
## License
This project is licensed under the MIT License.
