package main

func printTest1(ti TestInterface1) {
	ti.create()
}

func printTest2(ti TestInterface2) {
	ti.delete()
}

func main() {
	var test = testImpl{"test"}

	printTest1(test)
	printTest2(test)
}
