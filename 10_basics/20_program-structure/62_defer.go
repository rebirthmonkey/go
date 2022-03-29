package main

// 结果是 6：f4是匿名返回值，匿名返回值是在return执行时被声明，因此defer声明时，还不能访问到匿名返回值，defer的修改不会影响到返回值。
func f4() int { // 匿名返回值
	r := 6
	defer func() {
		r *= 7
	}()
	return r
}

// 结果是 42：f5先给返回值r赋值r=6，再执行defer语句修改r=42，然后函数return。
func f5() (r int) { // 有名返回值
	defer func() {
		r *= 7
	}()
	return 6
}

// 结果为 6：f6是有名返回值，但是因为r是作为defer的传参，在声明defer的时候，就进行参数拷贝传递，所以defer只会对defer函数的局部参数有影响，不会影响到调用函数的返回值。
func f6() (r int) { // 有名返回值
	defer func(r int) {
		r *= 7
	}(r)
	return 6
}

func main(){
	println("f4： ",f4())
	println("f5： ",f5())
	println("f6： ",f6())
}
