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
	predNext        pred = "next"
	predRequires    pred = "requires"
	predRequiredBy  pred = "requiredBy"
	predCauses      pred = "causes"
	predCausedBy    pred = "causedBy"
	predHasSubject  pred = "hasSubject"
	predState       pred = "state"
	predOwns        pred = "owns"
	predCanBeDoneBy pred = "canBeDoneBy"
)

var validPredicates = map[pred]bool{
	predNext:       true,
	predRequires:   true,
	predCauses:     true,
	predHasSubject: true,
	predOwns:       true,
	predRequiredBy: true,
	predCausedBy:   true,
}

// prop holds all the possible predicates that point to literal
// props
type prop string

const (
	propClass       prop = "class"
	propDescription prop = "description"
	propLabel       prop = "label"
	propMediaUrl    prop = "mediaUrl"

)

var validProps = map[prop]bool{
	propClass:       true,
	propDescription: true,
	propLabel:       true,
	propMediaUrl:    true,
}

type class string

const (
	classProcess class = "process"
	classStep          = "step"
	classEffect        = "effect"
	classMedia         = "media"
)

var validClasses = map[class]bool{
	classProcess: true,
	classStep:    true,
	classEffect:  true,
	classMedia:   true,
}