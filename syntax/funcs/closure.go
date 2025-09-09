package main

import "fmt"

var i = 0

func Closure(name string) func() string {
	// 闭包
	// name 变量
	// 方法本身
	return func() string {
		i++
		world := "world"
		return "hello, " + name + world
	}
}

func ClosureInvoke() {
	c := Closure("大明")
	println(c())
}

func DeferClosureLoopV1() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Printf("i 的地址是 %p，值是 %d\n", &i, i)
		}()
	}
}

func DeferClosureLoopV2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			fmt.Printf("i 的地址是 %p，值是 %d \n", &val, val)
		}(i)
	}
}

func DeferClosureLoopV3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			fmt.Printf("j 的地址是 %p，值是 %d\n", &j, j)
		}()
	}
}
