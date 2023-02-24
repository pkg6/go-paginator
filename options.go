package paginator

type Options struct {
	VarPage  string
	Path     string
	Query    map[string]string
	Fragment string
}

func DefaultOptions() Options {
	return Options{
		VarPage:  "page",
		Path:     "/",
		Query:    map[string]string{},
		Fragment: "",
	}
}
