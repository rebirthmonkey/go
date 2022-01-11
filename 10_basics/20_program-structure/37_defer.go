package main

// 结果是 6：f1 是匿名返回值，匿名返回值是在 return 执行时被声明，因此 defer 声明时，还不能访问到匿名返回值，defer 的修改不会影响到返回值。
func f1() int { // 匿名返回值
	var r int = 6
	defer func() {
		r *= 7
	}()
	return r
}

// 结果是 42：f2 先给返回值 r 赋值 r=6，再执行 defer 语句修改 r=42，然后函数 return。
func f2() (r int) { // 有名返回值
	defer func() {
		r *= 7
	}()
	return 6
}

// 结果为 6：f3 是有名返回值，但是因为 r 是作为 defer 的传参，在声明 defer 的时候，就进行参数拷贝传递，所以defer只会对defer函数的局部参数有影响，不会影响到调用函数的返回值。
func f3() (r int) { // 有名返回值
	defer func(r int) {
		r *= 7
	}(r)
	return 6
}

func main(){

	println("f1： ",f1())
	println("f2： ",f2())
	println("f3： ",f3())
}
