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
	relType string // pred, prop, class etc
	relName string // next, owns etc
	relSubName string // arbitrary map keys
}

func toiriString(relType, relName, relSubname string) string {
	base := fmt.Sprintf("%s:%s/%s", dl, relType, relName)
	if relSubname != "" {
		base = fmt.Sprintf("%s/%s", base, relSubname)
	}
	return base
}

func fromIriString(iri string) (*dlIRI, error) {
	domainParts := strings.Split(iri, ":")
	if len(domainParts) > 1 && domainParts[0] == dl {
		relParts := strings.Split(iri, "/")
		if len(relParts) < 2 || len(relParts) > 3 {
			return nil, errors.New(fmt.Sprintf("%s is not a valid dl iri tag (wrong length)", iri))
		}
		out := &dlIRI{
			relType: relParts[0],
			relName: relParts[1],
		}

		if len(relParts) == 3 {
			out.relSubName = relParts[3]
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
			return errors.Wrapf(err, "unable to get rdf literal from prop %s", p)
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
				Pred: mustIRI(toiriString(classTag,"instanceOf", "")),
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


func BindTermToField(r Resource, predicate, object rdf.Term) error {
	p := predicate.String()

}