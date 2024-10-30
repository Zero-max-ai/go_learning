# DAY 5 Go LEARNING

### what does defer / panic & recover does??
- defer - in terms of go programming defer is a block of code which will run after the end of the function..
For ex-
```
func hi () {
    defer func() { fmt.Println("hello") }
    fmt.Println("hi there")
}
```

in this above case the output is going to be like
`hi there` first & then then `hello` will be printed


- panic - in term of go programming panic is a situation where the program started to exit the execution of the function in which the panic function reside.

- recover - in term of go programming recover is a situation where the recover tries to make the function go normal after the panic situtaion. recover always runs in the defer function

```
func safeDivision(a, b int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from Panic: %v\n", r)
        }
    }()

    fmt.Println(a / b)
}
```

- named return values - this means that we are pre defining the function and telling what are the values we are going to return in every sitautation!
for ex-

```
func divide(a, b float32) (result float32, err error) {}
func main() {
    res, err := divide(10, 0)
    if err != nil {
        fmt.Println("error re")
        // return function automatically return the result & the error
        return
    }
    result := a/b
    // return the result & the error
    return
}
```
