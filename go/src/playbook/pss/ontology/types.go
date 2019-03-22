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

type stateThing interface {
	isStateThing()
}

type Process struct {
	Id          uid     `class:"class:process"`
	Label       string  `prop:"prop:label"`
	Description string  `prop:"prop:description"`
	Owns        []*Step `pred:"pred:owns"`
	Start       *Step   `pred:"pred:next"`
	CanBeDoneBy []*Role `pred:"pred:canBeDoneBy"`
}

func (v *Process) getIri() (rdf.IRI, error) {
	return rdf.NewIRI(fmt.Sprintf("id:%s", v.Id))
}

type Effect struct {
	Type       class
	Id         uid                   `class:"class:effect"`
	RequiredBy []*Step               `pred:"pred:requiredBy"`
	CausedBy   []*Step               `pred:"pred:causedBy"`
	HasSubject *Thing                `pred:"pred:hasSubject"`
	StateThing map[string]stateThing `predMap:""`
}

func (v *Effect) getIri() (rdf.IRI, error) {
	return rdf.NewIRI(fmt.Sprintf("id:%s", v.Id))
}

type Media struct {
	Id       uid    `class:"class:media"`
	Label    string `prop:"prop:label"`
	Url      string `prop:"prop:url"`
	Next     *Media `pred:"pred:next"`
	Previous *Media `pred:"pred:previous"`
}

func (v *Media) getIri() (rdf.IRI, error) {
	return rdf.NewIRI(fmt.Sprintf("id:%s", v.Id))
}

type Person struct {
	Id          uid     `class:"class:person"`
	Label       string  `prop:"prop:label"`
	Description string  `prop:"prop:description"`
	CanFill     []*Role `pred:"pred:canFill"`
}

func (v *Person) getIri() (rdf.IRI, error) {
	return rdf.NewIRI(fmt.Sprintf("id:%s", v.Id))
}

type Role struct {
	Id            uid       `class:"class:role"`
	Label         string    `prop:"prop:label"`
	Description   string    `prop:"prop:description"`
	CanBeFilledBy []*Person `pred:"pred:canBeFilledBy"`
}

func (v *Role) getIri() (rdf.IRI, error) {
	return rdf.NewIRI(fmt.Sprintf("id:%s", v.Id))
}

type State struct {
	Id    uid    `class:"class:state"`
	Label string `prop:"prop:label"`
}

func (v *State) getIri() (rdf.IRI, error) {
	return rdf.NewIRI(fmt.Sprintf("id:%s", v.Id))
}

type Thing struct {
	Id          uid                    `class:"class:thing"`
	Label       string                 `prop:"prop:label"`
	Description string                 `prop:"prop:description"`
	SubjectOf   []*Effect              `pred:"pred:subjectOf"`
	Properties  map[string]interface{} `propMap:""`
}

func (v *Thing) getIri() (rdf.IRI, error) {
	return rdf.NewIRI(fmt.Sprintf("id:%s", v.Id))
}

type Workflow struct {
	Type        class
	Id          uid        `class:"class:workflow"`
	Label       string     `prop:"prop:label"`
	Description string     `prop:"prop:description"`
	Owns        []*Process `pred:"pred:owns"`
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
	Id          uid        `class:"class:step"`
	Label       string     `prop:"prop:label"`
	Description string     `prop:"prop:description"`
	Next        []StepLike `pred:"pred:next"`
	OwnedBy     *Process   `pred:"pred:ownedBy"`
	Requires    []*Effect  `pred:"pred:requires"`
	Causes      []*Effect  `pred:"pred:causes"`
	HasMedia    *Media     `pred:"pred:hasMedia"`
}

func (v *Step) getIri() (rdf.IRI, error) {
	return rdf.NewIRI(fmt.Sprintf("id:%s", v.Id))
}

type Decision struct {
	Id          uid                 `class:"class:step"`
	Label       string              `prop:"prop:label"`
	Description string              `prop:"prop:description"`
	OwnedBy     *Process            `pred:"ownedBy"`
	Next        map[string]StepLike `predMap:""`
}

func (v *Decision) getIri() (rdf.IRI, error) {
	return rdf.NewIRI(fmt.Sprintf("id:%s", v.Id))
}
