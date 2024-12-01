package url

/*
	Exemplo de URL válida:
	scheme://hostname[:port]/path[?query][#fragment]

*/

type URL struct {
	scheme   string
	hostname string
	port     string
	path     string
	query    string
}

func New(url string) (*URL, error) {
	return nil, nil
}

func (u *URL) String() string {
	return ""
}

