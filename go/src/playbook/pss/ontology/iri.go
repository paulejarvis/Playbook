package ontology

import (
	"fmt"
	"github.com/knakk/rdf"
)

func mustIRI(s string) rdf.IRI {
	iri, err := rdf.NewIRI(s)
	if err != nil {
		panic(fmt.Sprintf("failed to create IRI! %q: %s", s, err))
	}
	return iri
}

func mustLiteral(v interface{}) rdf.Literal {
	lit, err := rdf.NewLiteral(v)
	if err != nil {
		panic(fmt.Sprintf("failed to create literal! %v: %s", v, err))
	}
	return lit
}

// pred holds all the possible predicates
// that link between nodes
type pred string

const (
	predNext       pred = "pred:next"
	predRequires   pred = "pred:requires"
	predRequiredBy pred = "pred:requiredBy"
	predCauses     pred = "pred:causes"
	predCausedBy     pred = "pred:causedBy"
	predHasSubject pred = "pred:hasSubject"
	predOwns       pred = "pred:owns"
	predCanBeDoneBy pred = "pred:canBeDoneBy"
)

var validPredicates = map[pred]bool{
	predNext:       true,
	predRequires:   true,
	predCauses:     true,
	predHasSubject: true,
	predOwns:       true,
	predRequiredBy: true,
	predCausedBy: true,
}

// prop holds all the possible predicates that point to literal
// props
type prop string

const (
	propClass       prop = "prop:class"
	propDescription prop = "prop:description"
	propLabel       prop = "prop:label"
	propMediaUrl    prop = "prop:mediaUrl"
)

var validProps = map[prop]bool{
	propClass:       true,
	propDescription: true,
	propLabel:       true,
	propMediaUrl:    true,
}

type class string

const (
	classProcess class = "class:process"
	classStep          = "class:step"
	classEffect        = "class:effect"
	classMedia = "class:media"
)

var validClasses = map[class]bool{
	classProcess: true,
	classStep:    true,
	classEffect:  true,
	classMedia: true,
}
