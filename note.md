# 创建 go 项目

为了创建 go 项目，首先得安装 go，然后 go mod init xxx，新建项目吗。
每个 go 项目都必须要有 main.go 作为入口：

```go
package main

func main()  {
	println("Hello Go!")
}
```

注意 main.go 的 package 名称必须是 main！。

然后是 go.mod，模块化管理工具：

```bash
module hello_go

go 1.21.0
```

### iota 与常亮组

在 go 中，提供了 iota 的关键字以声明常量组：

```go
package main

const (
	status1 = iota // 0
	status2 // 1
	status3 // 2
	status4 // 3
	status5 // 4
	status6 = 8
	// status7默认跟上面的值一样，然后继续计算
	status7 // 8
	status8  // 9
)

```

当第一个 const 常量被赋值 iota 后，后续字段的值将根据上一个字段的值继续加 1。
值得注意的是，iota 的值仍然在增加：

```go
const (
	status1 = iota // 0
	status2        // 1
	status3        // 2
	status4        // 3
	status5        // 4
	status6 = 8    // 8 (显式设置为 8, 此时 iota = 5)
	status7        // 8 (继承 status6 的值, 此时 iota = 6)
	status8        // 9 (继承 status7 的值 + 1, 此时 iota = 7)
	status9 = iota // 8 (使用当前 iota 的值, 此时 iota = 8)
	status10       // 9 (此时 iota = 9)
)
```

==在声明常量组时，尽管字段的值是字符串，但是这一组常量组的 iota 仍将从 0, 1, 2 进行递增==：

```go
const (
	status1 = "ready"    // status1 的值是 "ready", iota = 0
	status2 = "running"  // status2 的值是 "running", iota = 1
	status3 = "stopped"  // status3 的值是 "stopped", iota = 2
	code1   = iota       // code1 的值是 3, 因为 iota 此时 = 3
	code2   = 100        // code2 的值是 100, iota 此时 = 4
	code3   = iota       // code3 的值是 5, 因为 iota 此时 = 5
)
```

### 方法

go 声明方法的语法如下：

```go
func xxx(参数列表) (返回值列表) {

}
```

比如：

```go

func add(x int, y int) (int, error)  {
	return x + y, nil
}

func Func1(a, b, c int, str string) string {
	return ""
}

func Func2(a int, b int) (string, error) {
	return "", nil
}

func Func3(a int, b int) (str string, err error) {
	return "a", nil
}
```

go 中声明入参列表时，参数的类型写在变量后，同时返回值列表支持多个返回值，且返回值可以有名称。
使用：

```go
package main

func add(x int, y int) (int, error)  {
	return x + y, nil
}

func Func1(a, b, c int, str string) string {
	return str
}

func Func2(a int, b int) (string, error) {
	return "Func2", nil
}

func Func3(a int, b int) (str string, err error) {
	return "Func3", nil
}

func main()  {
	res, _ := add(2, 5);
	println(res)

	str := Func1(1, 1, 1, "aaa");
	println(str)

	str2, _ := Func2(2, 2);
	println(str2)

	str3, _ := Func3(5, 5);
	println(str3);
}
```

对于不想要的返回值，可以直接用 \_ 作为变量名。

### go 的函数使用

go 的函数使用有点类似于 js，支持函数式编程：

```go
package main

func add(x int, y int) (int, error)  {
	return x + y, nil
}

func Func1(a, b, c int, str string) string {
	return str
}

func Func2(a int, b int) (string, error) {
	return "Func2", nil
}

func Func3(a int, b int) (str string, err error) {
	return "Func3", nil
}

/**
内部函数
*/
func Func5(name string) string {
	fu := func (name string) string {
		return "hello ," + name;
	}

	return fu(name)
}

/**
函数作为返回值
*/
func Func6() func(name string) string {
	return func(name string) string {
		return "hello ," + name;
	}
}

/**
闭包
*/
func Func7(num int) func(y int) int {
	return func(y int) int {
		return num + y
	}
}

func main()  {
	res, _ := add(2, 5);
	println(res)

	str := Func1(1, 1, 1, "aaa");
	println(str)

	str2, _ := Func2(2, 2);
	println(str2)

	str3, _ := Func3(5, 5);
	println(str3);

	Func4 := Func3
	a, _ := Func4(2, 5);
	println(a)

	println(Func5("爹"))

	Funcx := Func6()
	println(Funcx("毛泽东"));

	add8 := Func7(8)
	println(add8(20)) // 28

	add7 := Func7(7)
	println(add7(20)) // 27
}
```

### 匿名函数与 defer

go 中声明的匿名函数可以立即执行：

```go
package main

func main()  {
	func() {
		println("立即执行的匿名函数")
	}()
}
```

上述立即执行的匿名函数常用于 defer：

```go
package main

func main()  {
	defer func() {
		println("defer1")
	}()

	defer func() {
		println("defer2")
	}()
}
```

defer 可以延迟匿名函数的执行，先后顺序的 defer 相当于先后入栈，因此在后面的 defer 会先执行。

#### defer 与闭包

defer 常与闭包一起使用，匿名函数中的变量的值由调用时使用：

```go
package main

/**
执行时确认
*/
func DeferClose1() {
	i := 0;
	defer func() {
		println((i))
	}()
	i++
}

/**
作为入餐传入，一开始就确定为0
*/
func DeferClose2() {
	i := 0;
	defer func(val int) {
		println((val))
	}(i)
	i++
}

func main()  {
	DeferClose1()
	DeferClose2()
}
```

defer 仅能修改命名返回值中的局部变量，看下面的例子：

```go
func DeferReturn() int {
	a := 0
	defer func() {
		a = 2
	}()
	return a
}
```

让我们深入分析一下这个函数的执行过程：

a := 0：这行代码初始化一个名为 a 的局部变量，并赋值为 0。

defer func() {...}()：这行代码注册一个匿名函数，它将在 DeferReturn 函数返回之前执行。这个匿名函数会修改 a 的值。

return a：这行代码是 DeferReturn 函数的返回语句。在返回语句执行时，它首先计算返回值（此时 a 的值是 0），然后它执行先前注册的 defer 函数。

关键的部分在于，当 return a 语句被执行时，返回值已经被确定为 0（a 在这时的值）。即使 defer 中的函数修改了 a 的值，这个修改对已经确定的返回值没有影响。

而下面的例子：

```go
func DeferReturn() (a int) {
	a = 0
	defer func() {
		a = 2
	}()
	return a
}
```

在这个版本中，a 是一个命名返回值。在 return a 语句执行时，它首先计算返回值（此时 a 的值是 0），然后执行先前注册的 defer 函数。由于 a 是一个命名返回值，defer 函数中对 a 的修改会影响到最终的返回值。

因此，DeferReturn() 函数的返回值在这个版本中将是 2。

#### defer 练习：

练习 1，打印 10 个 10:

```go
func Defer()  {
	for i := 0; i < 10; i++ {
		defer func() {
			println(i)
		}()
	}
}

func main()  {
	Defer()
}
```

练习 2，倒序打印 9 -> 0，注意 defer 是入栈结构：

```go
package main

func Defer()  {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			println(val)
		}(i)
	}
}

func main()  {
	Defer()
}
```

练习 3，每次都是新的局部变量：

```go
package main

func Defer()  {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			println(j)
		}()
	}
}
```

### if 语句

go 的 if 语句与 ts 的区别在于，if 和 if else 后面的条件不用括号。同时在 if 语句的条件前，支持声明一个新的变量：

```go
func IfNewVar(a, b int) string {
	if distance := a - b; distance >= 0 {
		return "a >= b"
	} else {
		return "a < b"
	}
}

```
