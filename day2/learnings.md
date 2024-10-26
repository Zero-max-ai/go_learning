# Go Learning DAY 2

### data types & variables
data types mean what type of data we have.
suppose we have a name so that will be string which can be wrapped around " quotes ", and if the data in a number we can use integer or floating point numbers, we also have boolean which mean the data either be true or false

- int , float , string , boolean

variables are the actual piece which hold the data itself. the above data types we have need to be stored somewhere, so that's how the born of variable happen!!

syntax for this in go going to be

```
var name string = "Zero-max-ai"

: In this example to initaite the variable we use the `var` keyword and then the name of the variable which we are going to called and after the we need to mention the state of data or the type of data, which is in our case going to be a string and the init of the actual data.
```

in var we have const also which means if we initalize the variable once then we can't change there values through out the scope.

#### inner dtypes
- int further divides into (int, int8, int16, int32, int64) - (uint, uint8, uint16).
- float can either be (float32 or float64)


### user\_inputs ???
- for user inputs we are requried to use Scan(&var\_name) or Scanln(&var\_name), it can also take multiple values together.

in int we can do printf likes %d to normal number, %b for binary, %o base 8(ig its hexa) nad %x & %X for base 16 lower and uppercase respectively.

in float we have %f for direct float number, %.2f for 2 digits are decimal and %8.2f for legit demm users. it just as 8 tab spaces initially before pass number! so again who the fucks uses it??? it used for the formatting the sheets and beyond our mind thinking cap.

in string we can do %s for string print, %q for quoted string print, %10s for space then fucking word again same as floating one!!, %-10s ahhhhhhhhhhh fuck (word      ) end result.

operation are same not rep! I know them

### functions & scope!!
a function is a block of instruction which can be run individually when they got called, which can take some parameters from us, and may or may not return some value as per the function ------ data type huh!!

soo what's scope? a scope i nothing but a vision how far we can see so suppose!

we have { saa    } in this braces saa can see within those braces so that's its scope and if we have
```
package main
import "fmt"

func main() {
    name := "zero"
    {
        hello := "yoah"
    }
}
```

in this case name is having global scope and hello is having limited. :)


### how to handle errors!!
but why to handle errors? as errors cause unexpected result of the program and we don't want to show such things to the users.. it basically crash the system man! and stop the program. soo handling errors are imp.

to do se we return nil values to the users check err\_hand.go file
