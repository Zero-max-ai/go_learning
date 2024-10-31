# Day 6 IN GO Programming

### what are pointers?
- pointers store the memory address of any variable rather the variable's value itself. Syntax is
`var p *int`

### but why use pointers?
it reduces memory by avoiding copying data, allow functions to modify the original data, imporove performance when data structure is large!

### types of pointers!
- address operator(&) = return memory address (less used)
- dereference operator(*) = access the value located at memory address (more used)

for exp
```
func main() {
    var x int = 10
    // here p hold the memory address of x
    var p *int = &x

    *p = 20 // now x value is going to be modified by using pointer
}
```

### garbage collection vs manual memory mgmt
- Go garbage collector automatically frees memory that is no longer referenced, it makes safer & easier for developers.
- C++ require the manual allocation and deallocation of the memory. which could lead the memory leaks.

- Go is easier to handle but have less control where as c++ provide full control but risk prone!
