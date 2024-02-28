package example

type Options struct {
	SomeOption string
}

func NewOptions() *Options {
	return &Options{
		SomeOption: "default",
	}
}
