package examples

type Foo struct {
}

func (foo Foo) Print() int {
	println("I am a Foo Command")
	return 0
}
