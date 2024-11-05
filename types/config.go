package types

type Repo struct {
	Name   string `yaml:"name"`
	URL    string `yaml:"url"`
	Branch string `yaml:"branch"`
}

type Config struct {
	Repositories *[]Repo `yaml:"repositories"`
}
