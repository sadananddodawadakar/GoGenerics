# GenericLinkedList: A Type-Safe Linked List in Go

The **GenericLinkedList** is a versatile, type-safe linked list implementation in Go. By leveraging Go's generics, this linked list supports any data type while maintaining clean and reusable code.

---

## 🧩 Features

- **Type Safety**: Handles any type with compile-time type checking.
- **Flexibility**: Suitable for both primitive and custom data types.
- **Ease of Use**: Provides intuitive APIs for adding and traversing elements.

---

## 📦 Installation

No additional dependencies are required. Copy the code into your project, or refactor it into a reusable package as needed.

---

## 🚀 Usage

### **Example Code**
Below is an example of how to use the `GenericLinkedList` to create and traverse a linked list of integers:

```go
package main

import "fmt"

func main() {
    // Create a linked list for integers
    intList := LinkedList[int]{}
    intList.Add(10)
    intList.Add(20)
    intList.Add(30)
    intList.Print() 
    // Output:
    // 10
    // 20
    // 30

    // Create a linked list for strings
    stringList := LinkedList[string]{}
    stringList.Add("Go")
    stringList.Add("Generics")
    stringList.Add("LinkedList")
    stringList.Print()
    // Output:
    // Go
    // Generics
    // LinkedList
}
