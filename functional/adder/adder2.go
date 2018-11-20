package adder

type IAdder func(int) (int, IAdder)

func Adder2(base int) IAdder {
	
	return func(v int) (int, IAdder) {
		return base + v, Adder2(base + v)
	}

}