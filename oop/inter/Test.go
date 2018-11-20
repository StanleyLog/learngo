package inter

type Text struct {
	Context string
}

func (t *Text) Get() string {
	return t.Context
}

func (t *Text) Set(arg interface{}) {
	t.Context = arg.(string)
}
