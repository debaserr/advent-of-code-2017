package main

type Program struct {
	name string
	weight int
	parent *Program
	children []*Program
}

func NewProgram(name string, weight int, children []string) *Program {
	program := &Program{
		name,
		weight,
		nil,
		make([]*Program, 0),
	}

	for _, child := range children {
		program.children = append(program.children, &Program{child, 0, nil, make([]*Program, 0),})
	}

	return program
}

func (p *Program) IsParent() bool {
	return len(p.children) > 0
}
