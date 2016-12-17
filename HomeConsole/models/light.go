package models

type Config struct {
	Lights []Light `xml:"light"`
}

type Light struct {
	Name string `xml:"Name"`
	URL  string `xml:"URL"`
}
