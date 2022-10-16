package main

func printTest(ti TestInterface) {
	ti.create()
}

func main() {
	t1 := getTest1()
	t1.create()
	t1.get()
	printTest(t1)

	t2 := getTest2()
	t2.create()
	t2.get()
	printTest(t2)
}
