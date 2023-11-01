# Go Basic

- https://man7.org/linux/man-pages/man1/find.1.html
- https://man7.org/linux/man-pages/man1/grep.1.html


```
go run <path.go>
go build <path.go>
./<path>
```

## 2023/10/18
- https://go.dev/tour/welcome/1
- https://go-tour-jp.appspot.com/welcome/1

- https://go-tour-jp.appspot.com/basics/1
- exported name: https://go-tour-jp.appspot.com/basics/3

```go
variable_name variable_type
// not variable_name: variable_type
// 
```
- function: https://go-tour-jp.appspot.com/basics/4
- naked return: https://go-tour-jp.appspot.com/basics/7
- implicity type: `:=`: https://go-tour-jp.appspot.com/basics/10

```go
fmt.Printf("%T %v", variable_name, variable_name)
```
- initial value: https://go-tour-jp.appspot.com/basics/12

- not while but for: https://go-tour-jp.appspot.com/flowcontrol/4

- switch: ./go_tour/content/flowcontrol/switch.go
    - it can make if-sentence simpler

- defer: https://go-tour-jp.appspot.com/flowcontrol/13

- pointer: https://go-tour-jp.appspot.com/moretypes/1
    - ./go_tour/content/moretypes/pointers.go

- struct pointer: https://go-tour-jp.appspot.com/moretypes/4
    - Is it ok to use `V.x`

- array: https://go-tour-jp.appspot.com/moretypes/6
- slice: https://go-tour-jp.appspot.com/moretypes/7
    - which are like references to array: https://go-tour-jp.appspot.com/moretypes/8
    - unnamed struct: https://go-tour-jp.appspot.com/moretypes/9
    - length and capacity: https://go-tour-jp.appspot.com/moretypes/11
        - panic: runtime error: slice bounds out of range [:10] with capacity 4
        - If you use [hoge:], capacity tends to decrease
    - make: https://go-tour-jp.appspot.com/moretypes/13
        - https://pkg.go.dev/golang.org/x/tour/pic#Show
    - nil: https://go-tour-jp.appspot.com/moretypes/12
    - append: https://go-tour-jp.appspot.com/moretypes/15

- for and range: https://go-tour-jp.appspot.com/moretypes/16

- map: 
    - how-to-create
        - make: https://go-tour-jp.appspot.com/moretypes/19
        - Don't use make: https://go-tour-jp.appspot.com/moretypes/20
    - how-to-get
        - https://go-tour-jp.appspot.com/moretypes/22
        - 危ない匂い: https://go-tour-jp.appspot.com/moretypes/22

- (other) godoc
    - `GOPATH=$(pwd)`
    - おそらく　moduleを作っていないため何も見れない

- function values:
    - https://go-tour-jp.appspot.com/moretypes/24

- method for struct
    - https://go-tour-jp.appspot.com/methods/1
- method for type
    - https://go-tour-jp.appspot.com/methods/3
- pointer to change field values
    - https://go-tour-jp.appspot.com/methods/4

- 感想: https://go-tour-jp.appspot.com/methods/8
    - 自分たちで気をつける必要があるため面倒そう

- interface
    - https://go-tour-jp.appspot.com/methods/9
    - implemented implicitly: https://go-tour-jp.appspot.com/methods/10
    - not substituted yet -> nil: https://go-tour-jp.appspot.com/methods/12

- anytime
    - empty interface: https://go-tour-jp.appspot.com/methods/14

- type switch
    - https://go-tour-jp.appspot.com/methods/16

- errors
    - https://go-tour-jp.appspot.com/methods/19

- io
    - https://go-tour-jp.appspot.com/methods/21

# NEXT
- https://go-tour-jp.appspot.com/methods/23
