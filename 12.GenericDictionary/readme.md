## Generic Dictionary
This example demonstrates how to implement a generic dictionary in Go. The Dictionary can store key-value pairs where the key is of a comparable type and the value can be of any type.

## Implementation
The implementation consists of the following types and functions:

Dictionary[K comparable, V any]: Represents a dictionary where K is the type of the key and V is the type of the value.
NewDictionary[K comparable, V any]: A constructor function that initializes and returns a new Dictionary.
Usage
The Dictionary type allows you to store values of different types using keys of a comparable type.

## Example
Here is an example of how to use the Dictionary:

Go
package main

import "fmt"

type Dictionary[K comparable, V any] struct {
    items map[K]V
}

func NewDictionary[K comparable, V any]() *Dictionary[K, V] {
    return &Dictionary[K, V]{items: make(map[K]V)}
}

func main() {
    dict := NewDictionary[string, int]()
    dict.items["age"] = 30
    fmt.Println(dict.items["age"]) // 30

    dict1 := NewDictionary[string, string]()
    dict1.items["Karnataka"] = "Bangalore"
    fmt.Println(dict1.items["Karnataka"])
}
## Running the Example
To run the example, simply execute the following command:

go run main.go
This will output:

Code
30
Bangalore
## License
This project is licensed under the MIT License.
