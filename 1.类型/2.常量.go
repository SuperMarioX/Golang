常量值必须是编译期可确定的数字、字符串、布尔值。

const x, y int = 1, 2   // 多常量初始化
const s = "Hello, World!"   // 类型推断

const (     //常量组
    a, b = 10, 100
    c bool = false
)

func main() {
    const x = "xxx" // 未使用的局部常量不会引发编译错误。
}

不支持 1UL、2LL 这样的类型后缀。 

在常量组中，如不提供类型和初始化值，那么视作与上一个常量相同。

const (
    s   = "abc"
    x   // x = "abc"
)

常量值还可以是 len、cap、unsafe.Sizeof 等编译期可确定结果的函数返回值。

const (
    a = "abc"
    b = len(a)
    c = unsafe.Sizeof(b)
)

如果常量类型足以存储初始化值，那么不会引发溢出错误。

const (
    a   byte = 100       // int to byte
    b   int  = 1e20      // float64 to int, overflows
)


枚举
关键字 iota 定义常量组中从 0 开始按行计数的自增枚举值。

const (
    Sunday = iota   // 0
    Monday          // 1  通常省略后续行表达式
    Tuesday         // 2
    Wednesday       // 3
    Thursday        // 4
    Friday          // 5
    Saturday        // 6

)

const (
    _       = iota              // iota = 0
    KB int64=1<<(10*iota) MB    // iota=1
    GB                          // 与 KB 表达式相同，但 iota = 2
    TB
)

在同一常量组中，可以提供多个 iota，它们各自增长。

const (
    A,B = iota,iota << 10     //0,0<<10 
    C,D                       // 1, 1 << 10
)

如果 iota 自增被打断，须显式恢复。
 
const (
    A =iota     //0
    B           // 1
    C="c"       //c
    D           // c，与上  相同。
    E = iota    // 4，显式恢复。注意计数包含了 C、D 两 。
    F           // 5
)

可通过自定义类型来实现枚举类型限制。

type Color int

const (
    Black Color = iota
    Red
    Blue
)

func test(c Color) {}

func main() {
    c := Black
    test(c)
    
    x := 1
    test(x) // Error: cannot use x (type int) as type Color in function argument

    test(1) // 常量会被编译器自动转换。 
}
