## Custom Constraints
This example demonstrates how to implement custom constraints in Go using generics. It shows how to define an interface to constrain the types that can be used with a generic function.

## Implementation
The implementation consists of the following:

Number interface: A custom constraint that allows only int and float64 types.
Sum[T Number]: A generic function that sums two values of type T, where T is constrained by the Number interface.
Usage
The Sum function can be used to sum two values of either int or float64 types.

## Example
Here is an example of how to use the custom constraints:

Go
package main

import (
    "fmt"
)

type Number interface {
    int | float64
}

func Sum[T Number](a, b T) T {
    return a + b
}

func main() {
    fmt.Println(Sum(1, 2))     // 3
    fmt.Println(Sum(1.1, 2.2)) // 3.3
}
## Running the Example
To run the example, simply execute the following command:

go run main.go
This will output:

Code
3
3.3
## License
This project is licensed under the MIT License.
