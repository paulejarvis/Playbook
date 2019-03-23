// Package ensemble provides apis for interactive with useful groupings
// of ontological types (aka subgraphs).
// Sometimes we want to interact with a process, but sometimes we want
// to interact with all the states of a thing. SubGraphs have cycles etc,
// so we need to think carefully about how to compose them.
package ensemble

import "dl/pss/ontology"

type Ensemble interface {
	UniqueElements() []ontology.Resource
}

type Process struct {
	elements map[string]ontology.Resource
	p *ontology.Process
}

func (p *Process) UniqueElements() []ontology.Resource {
	var ee []ontology.Resource
	for k := range p.elements {
		ee = append(ee, p.elements[k])
	}
	return ee
}

func NewProcess(elements []ontology.Resource) (*Process, error) {
	e := map[string]ontology.Resource{}

	for i := range elements {
		// Link elements into lookup
		e[elements[i].GetIri()] = elements[i]
	}

	for k := range e {
		r := e[k]

	}



	return &Process{}, nil
}
