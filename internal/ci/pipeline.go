package ci

type Pipeline struct {
	Name  string `yaml:"name"`
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Name     string   `yaml:"name"`
	Commands []string `yaml:"commands"`
}
