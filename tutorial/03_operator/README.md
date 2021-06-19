# 연산자

## Arithmetic

| Operator | Name | Types |
|:--------:|------|-------|
|`+`|sum|integers, floats, complex values, strings|
|`-`|difference|integers, floats, complex values|
|`*`|product|integers, floats, complex values|
|`/`|quotient|integers, floats, complex values|
|`%`|remainder|integers|
|`&`|bitwise AND|integers|
|`\|`|bitwise OR|integers|
|`^`|bitwise XOR|integers|
|`&^`|bit clear (AND NOT)|integers|
|`<<`|left shift|`integer << unsigned integer`|
|`>>`|right shift|`integer >> unsigned integer`|

> The shift operators shift the left operand by the shift count specified by the<br>
right operand, which must be non-negative. If the shift count is negative at run<br>
time, a [run-time panic](https://golang.org/ref/spec#Run_time_panics) occurs. The shift operators implement arithmetic shifts<br>
if the left operand is a signed integer and logical shifts if it is an unsigned<br>
integer. There is no upper limit on the shift count. Shifts behave as if the<br>
left operand is shifted n times by 1 for a shift count of n. As a result, `x << 1`<br>
is the same as `x*2` and `x >> 1` is the same as `x/2` but truncated towards<br>
negative infinity.

#### Example of bit clear operator

```go
    a := 0x0A // 1010
    b := 0x07 // 0111
    fmt.Println(a &^ b) // a &^ b
                        // a & ~b (AND NOT)
                        // 1010 & 1000
                        // :0x1000 (8)
```

## Comparison
### Comparison operators compare two operands and yield an untyped boolean value.

| Operator | Name | Types |
|:--------:|------|-------|
|`==`|equal|comparable|
|`!=`|not equal|comparable|
|`<`|less|integers, floats, strings|
|`<=`|less or equal|integers, floats, strings|
|`>`|greater|integers, floats, strings|
|`>=`|greater or equal|integers, floats, strings|

* Boolean, integer, floats, complex values and strings are comparable.
* Strings are ordered lexically byte-wise.
* Two pointers are equal if they point to the same variable or if both are nil.
* Two channel values are equal if they were created by the same call to make or if both are nil.
* Two interface values are equal if they have identical dynamic types and equal concrete values or if both are nil.
* A value x of non-interface type X and a value t of interface type T are equal if t’s dynamic type is identical * to X and t’s concrete value is equal to x.
* Two struct values are equal if their corresponding non-blank fields are equal.
* Two array values are equal if their corresponding elements are equal.

## Logical
### Logical operators apply to boolean values. The right operand is evaluated conditionally.
| Operator | Name | Types |
|:--------:|------|-------|
|`&&`|conditional AND | `p && q` means "if p then q else false"|
|`\|\|`|conditional OR | `p \|\| q` means "if p then true else q"|
|`!`|NOT|`!p` means "not p"|

## Pointers and channels
| Operator | Name | Types |
|:--------:|------|-------|
|`&`|address of|`&x` generates a pointer to `x`|
|`*`|pointer indirection|`*x` denotes the variable pointed to by `x`|
|`<-`|receive|`<-ch` is the value received from channel `ch`|

## Operator precedence
#### Go operator precedence spells `MACAO`
### Unary operators
* Unary operators have the highest priority and bind the strongest.
* `+`(positive) `-`(negative) `!`(not) `^`(bitwise complement)
* `*`(pointer indirection) `&`(address of) `<-`(receive)

### Binary operators (<b>MACAO</b>)
| Priority | Operators | Note |
|:--------:|------|-------|
|1|`* / % << >> & &^`|<b>M</b>ultiplicative|
|2|`+ - \| ^`|<b>A</b>dditive|
|3|`== != < <= > >=`|<b>C</b>omparison|
|4|`&&`|<b>A</b>nd|
|5|`\|\|`|<b>O</b>r|

* Binary operators of the same priority associate from <b>left to right</b>.

### Statement operators
* The `++` and `--` operators <b>form statements</b> and fall outside the operator hierarchy.

##### [Why are ++ and -- statements and not expressions? And why postfix, not prefix?](https://golang.org/doc/faq#inc_dec)
> &nbsp;&nbsp;
Without pointer arithmetic, the convenience value of pre- and postfix increment<br>
operators drops. By removing them from the expression hierarchy altogether,<br>
expression syntax is simplified and the messy issues around order of evaluation<br>
of ++ and -- (consider f(i++) and p[i] = q[++i]) are eliminated as well.<br>
&nbsp;&nbsp;The simplification is significant. As for postfix vs. prefix, either<br>
would work fine but the postfix version is more traditional; insistence on<br>
prefix arose with the STL, a library for a language whose name contains,<br>
ironically, a postfix increment.

### Examples
| Expression | Evaluation order |
|:--------:|:------:|
|`x/y*z`|`(x/y)*z`|
|`*p++`|`(*p)++`|
|`^a>>b`|`(^a)>>b`|
|`1+2*a[i]`|`1+(2*a[i])`|
|`m==n+1 && <-ch>0`|`(m==(n+1)) && (<-ch) > 0)`|
