// Package ontology provides a way of representing process/ nodes
// as go types. This provides the ability to write methods on top
// of node types/ collections of nodes. Node types in ontology are able
// to translate themselves into proto-api types and be correctly sent
// to the front end.
package ontology

import (
	"fmt"
	"github.com/knakk/rdf"
)

// TODO(hdh) write transformation logic so that a process can be "written" to db (updates and creates)
// TODO(hdh) write transformation logic so that a process can be read from the db
// TODO(hdh) write transformation logic so that a process can be sent along via proto

type uid string

// Any field that is designated with a predicate (pred:foo)
// Must be an addressable node, or a collection of addressableNodes
// Any field that is designated with a prop must the correct type for that predicate
// (haven't done this yet)
type node interface {
	getIri() (rdf.IRI, error)
}

type Process struct {
	Id          uid     `class:"process"`
	Label       string  `prop:"label"`
	Description string  `prop:"description"`
	Owns        []*Step `pred:"owns"`
	Start       *Step   `pred:"next"`
	CanBeDoneBy []*Role `pred:"canBeDoneBy"`
}

func (v *Process) getIri() (rdf.IRI, error) {
	return rdf.NewIRI(fmt.Sprintf("id:%s", v.Id))
}

type Effect struct {
	Id         uid                   `class:"effect"`
	RequiredBy []*Step               `pred:"requiredBy"`
	CausedBy   []*Step               `pred:"causedBy"`
	HasSubject *Thing                `pred:"hasSubject"`
	StateThing map[string]stateThing `predMap:"state"`
}

func (v *Effect) getIri() (rdf.IRI, error) {
	return rdf.NewIRI(fmt.Sprintf("id:%s", v.Id))
}

type Media struct {
	Id       uid    `class:"media"`
	Label    string `prop:"label"`
	Url      string `prop:"url"`
	Next     *Media `pred:"next"`
	Previous *Media `pred:"previous"`
}

func (v *Media) getIri() (rdf.IRI, error) {
	return rdf.NewIRI(fmt.Sprintf("id:%s", v.Id))
}

type Person struct {
	Id          uid     `class:"person"`
	Label       string  `prop:"label"`
	Description string  `prop:"description"`
	CanFill     []*Role `pred:"canFill"`
}

func (v *Person) getIri() (rdf.IRI, error) {
	return rdf.NewIRI(fmt.Sprintf("id:%s", v.Id))
}

type Role struct {
	Id            uid       `class:"role"`
	Label         string    `prop:"label"`
	Description   string    `prop:"description"`
	CanBeFilledBy []*Person `pred:"canBeFilledBy"`
}

func (v *Role) getIri() (rdf.IRI, error) {
	return rdf.NewIRI(fmt.Sprintf("id:%s", v.Id))
}

type stateThing interface {
	isStateThing()
}

type State struct {
	Id    uid    `class:"state"`
	Label string `prop:"label"`
}

func (v *State) getIri() (rdf.IRI, error) {
	return rdf.NewIRI(fmt.Sprintf("id:%s", v.Id))
}

func (v *State) isStateThing() {}

type Thing struct {
	Id          uid                    `class:"thing"`
	Label       string                 `prop:"label"`
	Description string                 `prop:"description"`
	SubjectOf   []*Effect              `pred:"subjectOf"`
	Properties  map[string]interface{} `propMap:""`
}

func (v *Thing) getIri() (rdf.IRI, error) {
	return rdf.NewIRI(fmt.Sprintf("id:%s", v.Id))
}

func (v *Thing) isStateThing() {}

type Workflow struct {
	Type        class
	Id          uid        `class:"workflow"`
	Label       string     `prop:"label"`
	Description string     `prop:"description"`
	Owns        []*Process `pred:"owns"`
}

func (v *Workflow) getIri() (rdf.IRI, error) {
	return rdf.NewIRI(fmt.Sprintf("id:%s", v.Id))
}

type StepLike interface {
	Next() []struct {
		Choice string // Can be blank
		Next   StepLike
	}
	Previous() StepLike
}

type Step struct {
	Id          uid        `class:"step"`
	Label       string     `prop:"label"`
	Description string     `prop:"description"`
	Next        []StepLike `pred:"next"`
	OwnedBy     *Process   `pred:"ownedBy"`
	Requires    []*Effect  `pred:"requires"`
	Causes      []*Effect  `pred:"causes"`
	HasMedia    *Media     `pred:"hasMedia"`
}

func (v *Step) getIri() (rdf.IRI, error) {
	return rdf.NewIRI(fmt.Sprintf("id:%s", v.Id))
}

type Decision struct {
	Id          uid                 `class:"step"`
	Label       string              `prop:"label"`
	Description string              `prop:"description"`
	OwnedBy     *Process            `pred:"ownedBy"`
	Next        map[string]StepLike `predMap:"next"`
}

func (v *Decision) getIri() (rdf.IRI, error) {
	return rdf.NewIRI(fmt.Sprintf("id:%s", v.Id))
}
