# GO DAY 4 

### function in more advance
function declaration goes like
```
func function_Name(parameterName, parmaterType) returnType {
// function code
}

func add(a int, b int) int {
    return a + b
}
```

but the function can return multiple types
```
func functionName(paramName, paramType) (returnType1, returnType2) {
// function code
}

func divide(a, b float32) (float32, error) {
    if b == 0 {
        return 0, fmt.Error("cannot divide by zero")
    }
    return a / b, nil;
}
```

### error handling
for error handling in go is simple, first we need to do an operation & then if it cause any error than captues it! and as per the error we take corresponding takes regarding this situation.

## comman ways to do so
```
result, err := someFunction()
if err != nil {
// take some action
}

func checkAge(age int) error {
    if age < 18 {
        return fmt.Errorf("age %d is below the legal age",age)
    }
    return nil
}
```

we use Errorf for formating the error message.
