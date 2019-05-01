// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package elemental

import (
	"fmt"
	"net/http"
	"reflect"
)

const (
	readOnlyErrorTitle     = "Read Only Error"
	creationOnlyErrorTitle = "Creation Only Error"
)

// ValidateAdvancedSpecification verifies advanced specifications attributes like ReadOnly and CreationOnly.
//
// For instance, it will check if the given Manipulable has field marked as
// readonly, that it has not changed according to the db.
func ValidateAdvancedSpecification(obj AttributeSpecifiable, pristine AttributeSpecifiable, op Operation) error {

	errors := NewErrors()

	for _, field := range extractFieldNames(obj) {

		spec := obj.SpecificationForAttribute(field)

		// If the field is not exposed, we don't enforce anything.
		if !spec.Exposed || spec.Transient {
			continue
		}

		switch op {
		case OperationCreate:
			if spec.ReadOnly && !isFieldValueZero(field, obj) && !areFieldsValueEqualValue(field, obj, spec.DefaultValue) {

				// Special case here. If we have a pristine object, and the fields are equal, it is fine.
				if pristine != nil && areFieldValuesEqual(field, obj, pristine) {
					continue
				}

				e := NewError(
					readOnlyErrorTitle,
					fmt.Sprintf("Field %s is read only. You cannot set its value.", spec.Name),
					"elemental",
					http.StatusUnprocessableEntity,
				)
				e.Data = map[string]string{"attribute": spec.Name}
				errors = append(errors, e)
			}

		case OperationUpdate:
			if !spec.CreationOnly && spec.ReadOnly && !areFieldValuesEqual(field, obj, pristine) {
				e := NewError(
					readOnlyErrorTitle,
					fmt.Sprintf("Field %s is read only. You cannot modify its value.", spec.Name),
					"elemental",
					http.StatusUnprocessableEntity,
				)
				e.Data = map[string]string{"attribute": spec.Name}
				errors = append(errors, e)
			}

			if spec.CreationOnly && !areFieldValuesEqual(field, obj, pristine) {
				e := NewError(
					creationOnlyErrorTitle,
					fmt.Sprintf("Field %s can only be set during creation. You cannot modify its value.", spec.Name),
					"elemental",
					http.StatusUnprocessableEntity,
				)
				e.Data = map[string]string{"attribute": spec.Name}
				errors = append(errors, e)
			}
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// BackportUnexposedFields copy the values of unexposed fields from src to dest.
func BackportUnexposedFields(src, dest AttributeSpecifiable) {

	for _, field := range extractFieldNames(src) {

		spec := src.SpecificationForAttribute(field)

		if !spec.Exposed {
			reflect.Indirect(reflect.ValueOf(dest)).FieldByName(field).Set(reflect.Indirect(reflect.ValueOf(src)).FieldByName(field))
		}

		if spec.Secret && isFieldValueZero(field, dest) {
			reflect.Indirect(reflect.ValueOf(dest)).FieldByName(field).Set(reflect.Indirect(reflect.ValueOf(src)).FieldByName(field))
		}
	}
}

// ResetDefaultForZeroValues reset the default value from the specification when a field is Zero.
func ResetDefaultForZeroValues(obj AttributeSpecifiable) {

	for _, field := range extractFieldNames(obj) {

		spec := obj.SpecificationForAttribute(field)

		if spec.DefaultValue == nil || !isFieldValueZero(field, obj) {
			continue
		}

		reflect.Indirect(reflect.ValueOf(obj)).FieldByName(field).Set(reflect.ValueOf(spec.DefaultValue))
	}
}
