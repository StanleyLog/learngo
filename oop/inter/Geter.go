package inter

type Getter interface {
	Get() string
}

type Setter interface {
	Set(arg interface{})
}