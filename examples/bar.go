package examples

type Bar struct {
}

func (bar Bar) Print() int {
	println("I am a Bar Command")
	return 0
}
