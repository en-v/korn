package holder

import (
	"reflect"

	"github.com/en-v/korn/inset"
	"github.com/pkg/errors"
)

type Reference struct {
	Type    reflect.Type
	Name    string
	Pointer reflect.Type
}

func makeRef(obj interface{}) (*Reference, error) {

	if obj == nil {
		return nil, errors.New("Reference object can't to be a nil pointer")
	}

	refType := reflect.TypeOf(obj)
	typeName := refType.String()

	if refType.Kind() != reflect.Struct {
		return nil, errors.New("Reference object has to be a structure only, " + typeName)
	}

	// INSET FEILD
	insetField, found := refType.FieldByName(inset.NAME)
	if !found {
		return nil, errors.New("Embedded korn.Inset not found in refernce object, " + typeName)
	}

	bsonTag, found := insetField.Tag.Lookup("bson")
	if !found {
		return nil, errors.New("Embedded korn.Inset have to have 'bson' tag, " + typeName)
	}

	if bsonTag != ",inline" {
		return nil, errors.New("Embedded korn.Inset have to have 'bson' tag with the value ',inline', " + typeName)
	}

	kornTag, found := insetField.Tag.Lookup("korn")
	if !found {
		return nil, errors.New("Embedded korn.Inset have to have 'korn' tag, " + typeName)
	}

	if kornTag != "-" {
		return nil, errors.New("Embedded korn.Inset have to have 'bson' with the value '-', " + typeName)
	}

	// ID FIELD
	idField, found := refType.FieldByName("Id")
	if !found {
		return nil, errors.New("Id field not found in refernce object, " + typeName)
	}

	idRequiredTag, found := idField.Tag.Lookup(inset.REQUIRED)
	if !found {
		return nil, errors.New("Id field have to have 'required' tag, " + typeName)
	}

	if idRequiredTag != "true" {
		return nil, errors.New("Id field have to have 'required' tag value 'true', " + typeName)
	}

	// UPDATE FIELD
	updateField, found := refType.FieldByName("Updated")
	if !found {
		return nil, errors.New("Update field not found in refernce object, " + typeName)
	}

	updateRequiredTag, found := updateField.Tag.Lookup(inset.REQUIRED)
	if !found {
		return nil, errors.New("Update field have to have 'required' tag, " + typeName)
	}

	if updateRequiredTag != "true" {
		return nil, errors.New("Update field have to have 'required' tag value 'true', " + typeName)
	}

	refPointer := reflect.TypeOf(reflect.New(refType).Interface())

	return &Reference{
		Type:    refType,
		Pointer: refPointer,
		Name:    typeName,
	}, nil
}
