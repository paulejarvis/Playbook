package ontology

import (
	"fmt"
	"github.com/knakk/rdf"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

// Tags
const (
	classTag   = "class"
	predTag    = "pred"
	propTag    = "prop"
	predMapTag = "predMap"
	propMapTag = "propMap"
)

const dl = "dl"

// DL iris are the internal format we use for disco lab resource IDs
type dlIRI struct {
	nameSpace string // id, pred, prop, class etc
	name      string // next, owns, literal-id etc
	subName   string // arbitrary map keys
}

func toiriString(relType, relName, relSubname string) string {
	base := fmt.Sprintf("%s:%s/%s", dl, relType, relName)
	if relSubname != "" {
		base = fmt.Sprintf("%s/%s", base, relSubname)
	}
	return base
}

func fromIriString(iri string) (*dlIRI, error) {
	fmt.Println(iri)
	domainParts := strings.Split(iri, ":")
	if len(domainParts) > 1 && domainParts[0] == dl {
		relParts := strings.Split(domainParts[1], "/")
		if len(relParts) < 2 || len(relParts) > 3 {
			return nil, errors.New(fmt.Sprintf("%s is not a valid dl iri tag (wrong length)", iri))
		}
		out := &dlIRI{
			nameSpace: relParts[0],
			name:      relParts[1],
		}

		if len(relParts) == 3 {
			out.subName = relParts[2]
		}

		return out, nil

	}

	return nil, errors.New(fmt.Sprintf("%s is not a valid dl iri tag (wrong prefix)", iri))
}

func ToRDF(n Resource) ([]rdf.Triple, error) {
	var out []rdf.Triple


	subject, err := rdf.NewIRI(n.GetIri())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get IRI")
	}

	nVal := reflect.ValueOf(n)
	if nVal.Kind() == reflect.Ptr {
		nVal = nVal.Elem()
	}

	nType := reflect.TypeOf(n)
	if nType.Kind() == reflect.Ptr {
		nType = nType.Elem()
	}

	serializePred := func(fieldVal reflect.Value, name, subname string) error {
		objAn, ok := fieldVal.Interface().(Resource)
		if !ok {
			return errors.New("predicates can only refer to addressable nodes")
		}

		pred, err := rdf.NewIRI(toiriString(predTag, name, subname))
		if err != nil {
			return errors.Wrap(err, "failed to get iri from pred")
		}

		obj, err := rdf.NewIRI(objAn.GetIri())
		if err != nil {
			return errors.Wrap(err, "failed to get IRI from Resource")
		}
		out = append(out, rdf.Triple{Subj: subject, Pred: pred, Obj: obj})

		return nil
	}

	serializeProp := func(fieldVal reflect.Value, name, subname string) error {
		obj, err := rdf.NewLiteral(fieldVal.Interface())
		if err != nil {
			return errors.Wrapf(err, "unable to get rdf literal from prop %s", name)
		}

		pred, err := rdf.NewIRI(toiriString(propTag, name, subname))
		if err != nil {
			return errors.Wrap(err, "failed to get iri from prop")
		}

		trip := rdf.Triple{
			Subj: subject,
			Pred: pred,
			Obj:  obj,
		}

		out = append(out, trip)
		return nil
	}

	for i := 0; i < nType.NumField(); i++ {
		field := nType.Field(i)
		fieldVal := nVal.Field(i)
		fType := field.Type
		tags := field.Tag

		c, ok := tags.Lookup(classTag)
		if ok {
			if valid := validClasses[class(c)]; !valid {
				return nil, errors.New(fmt.Sprintf("class %q is not valid", c))
			}
			trip := rdf.Triple{
				Subj: subject,
				Pred: mustIRI(toiriString(propTag, string(propClass), "")),
				Obj:  mustLiteral(c),
			}
			out = append(out, trip)
			continue
		}

		pre, ok := tags.Lookup(predTag)
		if ok {
			// Need to handle if this is a slice, or a map, or pointer
			switch fType.Kind() {
			case reflect.Slice:
				for vi := 0; vi < fieldVal.Len(); vi++ {
					if err := serializePred(fieldVal.Index(vi), pre, ""); err != nil {
						return nil, errors.Wrap(err, "failed to serialize predicate")
					}
				}

			case reflect.Ptr:
				if err := serializePred(fieldVal, pre, ""); err != nil {
					return nil, errors.Wrap(err, "failed to serialize predicate")
				}

			case reflect.Struct:
				if err := serializePred(fieldVal, pre, ""); err != nil {
					return nil, errors.Wrap(err, "failed to serialize predicate")
				}

			default:
				return nil, errors.New(fmt.Sprintf("%s is not a valid predicate object", fType.Kind().String()))

			}
			continue
		}

		mt, ok := tags.Lookup(predMapTag)
		if ok {
			for _, k := range fieldVal.MapKeys() {
				if err := serializePred(fieldVal.MapIndex(k), mt, k.String()); err != nil {
					return nil, errors.Wrap(err, "failed to serialize predicate")
				}
			}
		}

		pro, ok := tags.Lookup(propTag)
		if ok {
			if reflect.Zero(fieldVal.Type()).Interface() == fieldVal.Interface() {
				continue
			}

			if err := serializeProp(fieldVal, pro, ""); err != nil {
				return nil, errors.Wrap(err, "failed to serialize prop")
			}

			continue
		}

		mt, ok = tags.Lookup(propMapTag)
		if ok {
			for _, k := range fieldVal.MapKeys() {
				if err := serializeProp(fieldVal.MapIndex(k), mt, k.String()); err != nil {
					return nil, errors.Wrap(err, "failed to serialize property")
				}
			}
			continue
		}
	}

	return out, nil
}

type PredObjPair struct {
	P rdf.Term
	O rdf.Term
}

func DecodeTerms(r Resource, pairs []PredObjPair) error {
	for _, p := range pairs {
		if err := DecodeTerm(r, p.P, p.O); err != nil {
			return errors.Wrap(err, "failed to decode term")
		}
	}
	return nil
}


// DecodeTerm parses a predicate and an object and binds them to the correct
// field on a resource, for a given set of struct tags.
func DecodeTerm(r Resource, predicate, object rdf.Term) error {
	pIri, err := fromIriString(predicate.String())
	if err != nil {
		return errors.Wrap(err, "failed to get IRI from predicate")
	}

	if pIri.nameSpace == propTag && pIri.name == string(propClass) {
		// Class is weirdly encoded as a shortcut, so need to carve out exception.
		return nil
	}

	rVal := reflect.ValueOf(r)
	if  rVal.Kind() != reflect.Ptr || rVal.Elem().Kind() != reflect.Struct || !rVal.Elem().CanSet() {
		return errors.New("resource must be a settable struct ptr")
	}
	rVal = rVal.Elem()
	rType := rVal.Type()

	for i := 0; i < rType.NumField(); i++ {
		// Iterate through the struct fields until we find the correctly tagged relationship,
		// then set it.
		field := rType.Field(i)
		fVal := rVal.Field(i)

		tagVal, ok := field.Tag.Lookup(pIri.nameSpace)
		if !ok {
			// Wrong kind of relationship
			continue
		}

		if tagVal != pIri.name {
			// Write kind of relationship, but wrong relationship
			continue
		}

		var oVal reflect.Value
		switch pIri.nameSpace {
		case propTag, propMapTag:

			// In this case our object is a literal
			oLit, ok := object.(rdf.Literal)
			if !ok {
				return errors.New("object should be an rdf.literal, it is not")
			}
			o, err := oLit.Typed()
			if err != nil {
				return errors.Wrap(err, "failed to convert o to proper type")
			}

			oVal = reflect.ValueOf(o)

		case predTag,predMapTag:
			// In this case our object is an iri pointing to an entity
			// represented as a go type, we have to make a new version of that type and set its id.
			// TODO(henry) this is going to explode, make it not

			if fVal.Type().Kind() == reflect.Map || fVal.Type().Kind() == reflect.Slice {
				oVal = reflect.New(fVal.Type().Elem().Elem())
			} else {
				oVal = reflect.New(fVal.Type().Elem())
			}


			objectIdIri, err := fromIriString(object.String())
			if err != nil {
				return errors.Wrap(err, "failed to parse object iri")
			}
			oVal.Elem().FieldByName("Id").Set(reflect.ValueOf(objectIdIri.name))
		}

		// Now actually set the value at the field.
		switch pIri.nameSpace {
		case propTag, predTag:
			if fVal.Type().Kind() == reflect.Slice {
				slice := fVal
				if fVal.IsNil() {
					slice = reflect.MakeSlice(fVal.Type(), 0,1)
				}
				fVal.Set(reflect.Append(slice, oVal))
			} else {
				fVal.Set(oVal)
			}
			return nil

		case propMapTag, predMapTag:
			if fVal.IsNil() {
				fVal.Set(reflect.MakeMap(fVal.Type()))
			}
			fVal.SetMapIndex(reflect.ValueOf(pIri.subName), oVal)
			return nil
		}

	}
	return nil
}
