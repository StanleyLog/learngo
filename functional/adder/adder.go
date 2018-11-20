package adder

func Adder() (func(int) int)  {

	iv := 0

	return func(v int) int {
		iv += v
		return iv
	}
}

func test(v *int)  {
	*v += 1
}
