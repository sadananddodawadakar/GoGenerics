# GenericQueue: A Type-Safe Queue Implementation in Go

The **GenericQueue** package provides a simple yet powerful implementation of a queue data structure in Go, utilizing generics for type safety and flexibility. This implementation supports any type of data while maintaining clean and efficient code.

---

## 🧩 Features

- **Type Safety**: Works with any data type while ensuring compile-time type safety.
- **Ease of Use**: Simple APIs for enqueueing and dequeueing items.
- **Reusability**: One implementation works for all types, eliminating the need for redundant code.

---

## 📦 Installation

No special installation is required. Simply copy the code into your project or import the package directly if available.

---

## 🚀 Usage

### **Example Code**
Here’s how you can use the `GenericQueue` to work with different data types:

```go
package main

import "fmt"

func main() {
    // Integer Queue
    intQueue := Queue[int]{}
    intQueue.Enqueue(10)
    intQueue.Enqueue(20)
    fmt.Println(intQueue.Dequeue()) // Output: 10
    fmt.Println(intQueue.Dequeue()) // Output: 20

    // String Queue
    stringQueue := Queue[string]{}
    stringQueue.Enqueue("Go")
    stringQueue.Enqueue("Generics")
    fmt.Println(stringQueue.Dequeue()) // Output: Go
    fmt.Println(stringQueue.Dequeue()) // Output: Generics
}
