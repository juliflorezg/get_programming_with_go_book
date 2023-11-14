package main

type Planets []string

func (p Planets) terraform() {
	for i := range p {
		p[i] = "New " + p[i]
	}
}

func terraform(worlds []string) {
	for i := range worlds {
		worlds[i] = "New " + worlds[i]
	}
}
