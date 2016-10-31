package bench

func hello(name string) string {
	return "Hello " + name + "!"
}

// RO is Reomote object
type RO struct {
	Hello func(string) (string, error)
}

// Args ...
type Args struct {
	Name string
}

// Hello ...
type Hello int

// Hello ...
func (*Hello) Hello(args *Args, result *string) error {
	*result = "Hello " + args.Name + "!"
	return nil
}
