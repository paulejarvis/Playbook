package dbtranslation

import ps "dl/apis/procstore"

// The types in this file are directly reflective of how the data is stored, and should only be used
// Inside the scope of this package. They do not have any guarantees of sanity or correctness
// Vis-a-vis the process being created. This is brittle and needs a lot more thought/ testing.
// Sanity and correctness is currently enforced by the procstore service types.

// nRef is how nodes reference each other in the context of a dgraph transaction
type nref struct {
	Uid string `json:"uid"`
}

// Mappers translate themselves into slices
// Of maps that are correctly formatted to produce
// the underlying structure desired in dgraph.
// We Need to create a structure that looks like the following:
// Where foo and bar are the content hashes we care about,
// or the literal uid's of nodes in the case of update.
//
// [
// {uid:_:foo .... all the props}
// {uid:_:bar .... all the props}
// {uid:_:bar :relation_to {"uid":"_foo"}}
// ]

type Mapper interface {
	ToMaps() []map[string]interface{}
}

func fmtUid(s string) string {
	return "_:" + s
}

func fmtNref(n nref) nref {
	return nref{Uid:fmtUid(n.Uid)}
}

func relMapTranslation(fmtedUid string, rels map[string]nref) []map[string]interface{} {
	var out []map[string]interface{}
	for k, v := range rels {
		out = append(out, map[string]interface{}{
			"uid": fmtedUid,
			k: fmtNref(v),
		})
	}
	return out
}

func relSliceTranslation(fmtedUid, relType string, rels []nref) []map[string]interface{} {
	var out []map[string]interface{}
	for _, v := range rels {
		out = append(out,map[string]interface{}{
			"uid": fmtedUid,
			relType: fmtNref(v),
		})
	}
	return out
}

func appendIfNotEmpty(slice []map[string]interface{}, uid, rel string, ref nref) []map[string]interface{}{
	if ref.Uid != "" {
		return append(slice,map[string]interface{}{"uid":uid, rel: fmtNref(ref)})
	}
	return slice
}

type process struct {
	// System-level descriptors
	nref
	NodeKind string `json:"nodeKind,omitempty"`
	HashId string   `json:"hashId,omitempty"`

	// Properties on process
	Name string        `json:"name,omitempty"`
	Description string `json:"description,omitempty"`

	// Relations
	StartsWith nref `json:"startsWith,omitempty"`
}

func (p *process) ToMaps() []map[string]interface{} {
	uid := fmtUid(p.Uid)
	out := []map[string]interface{}{{
		"uid":         uid,
		"nodeKind":    p.NodeKind,
		"hashId":      p.HashId,
		"name":        p.Name,
		"description": p.Description,
	}}

	if p.StartsWith.Uid != "" {
		out = append(out, map[string]interface{}{"uid": uid, "StartsWith": fmtNref(p.StartsWith)})
	}

	return out
}


type step struct {
	// System-level descriptors
	nref
	NodeKind string `json:"nodeKind,omitempty"`
	HashId string   `json:"hashId,omitempty"`

	// Properties on step
	Name string        `json:"name,omitempty"`
	Description string `json:"description,omitempty"`

	// Relations
	FollowedBy map[string]nref `json:"followedBy,omitempty"`
	PrecededBy []nref          `json:"precededBy,omitempty"`
	ExecutedBy nref            `json:"executedBy,omitempty"`
	Causes []nref              `json:"causes,omitempty"`
	Requires []nref            `json:"requires,omitempty"`
}

func (s *step) ToMaps() []map[string]interface{} {
	uid := fmtUid(s.Uid)
	// Create the node that contains all the properties
	out := []map[string]interface{}{{
		"uid":         uid,
		"nodeKind":    s.NodeKind,
		"hashId":      s.HashId,
		"name":        s.Name,
		"description": s.Description,
	}}

	if s.ExecutedBy.Uid != "" {
		out = append(out, map[string]interface{}{"uid": uid, "executedBy": fmtNref(s.ExecutedBy)})
	}

	out = append(out, relMapTranslation(uid, s.FollowedBy)...)
	out = append(out, relSliceTranslation(uid, "precededBy", s.PrecededBy)...)
	out = append(out, relSliceTranslation(uid, "causes", s.Causes)...)
	out = append(out, relSliceTranslation(uid, "requires", s.Requires)...)

	return out
}

type object struct {
	nref
	NodeKind string `json:"nodeKind,omitempty"`
	HashId string   `json:"hashId,omitempty"`
	ObjectType string `json:"objectType,omitempty"`

	// Properties on object class
	Name string        `json:"name,omitempty"`
	Description string `json:"description,omitempty"`

	ExistsIn nref

	// only one of these should be set.
	IsIn nref
	IsNotIn nref

	// Thing Fields
	Thing_Props map[string]string // Make this interface

	// Tool Fields
	Tool_LoginPage string

	// Qualification Fields
	Qualification_CertifiedBy nref

	// Employee Fields
	Employee_RoleName string
	Employee_CanBe []nref

	// Machine
	Machine_ProgramName string


}

func (o *object) ToMaps() []map[string]interface{} {
	uid := fmtUid(o.Uid)
	// Create the node that contains all the properties
	out := []map[string]interface{}{{
		"uid":         uid,
		"nodeKind":    o.NodeKind,
		"objectType":  o.ObjectType,
		"hashId":      o.HashId,
		"name":        o.Name,
		"description": o.Description,

	}}

	out = appendIfNotEmpty(out, uid, "existsIn", o.ExistsIn)
	out = appendIfNotEmpty(out, uid, "isIn", o.IsIn)
	out = appendIfNotEmpty(out, uid, "isNotIn", o.IsNotIn)

	// SubType Properties
	sub := make(map[string]interface{})
	sub["uid"] = uid

	switch ps.Object_ObjectType(ps.Object_ObjectType_value[o.ObjectType]) {
	case ps.Object_OT_THING:
		for k, v := range o.Thing_Props {
			sub[k] = v
		}

	case ps.Object_OT_TOOL:
		sub["loginPage"] = o.Tool_LoginPage

	case ps.Object_OT_QUALIFICATION:
		out = appendIfNotEmpty(out,uid, "certifiedBy", o.Qualification_CertifiedBy)

	case ps.Object_OT_EMPLOYEE:
		sub["roleName"] = o.Employee_RoleName
		out = append(out, relSliceTranslation(uid, "canBe", o.Employee_CanBe)...)

	case ps.Object_OT_MACHINE:
		sub["programName"] = o.Machine_ProgramName
	}

	if len(sub) > 1 {
		out = append(out, sub)
	}

	return out
}

 type state struct {
 	subject   nref
 	predicate string
 	object    struct {
 		nref
 		Name string
	}
 }

 func (s *state) ToMaps() []map[string]interface{} {
 	obj := make(map[string]string)
 	if s.object.Name != "" {
 		obj["name"] = s.object.Name
	} else {
		obj["uid"] = fmtUid(s.object.Uid)
	}

	 return []map[string]interface{}{{
		"uid":       fmtUid(s.subject.Uid),
		s.predicate: obj,
	 }}
 }



