// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import "sort"

// BoolValidators type defines BoolValidator types
type BoolValidators []BoolValidator

// Equal returns true if the given BoolValidators is the same
// length, and each of the BoolValidator entries is equal.
func (v BoolValidators) Equal(other BoolValidators) bool {
	if v == nil && other == nil {
		return true
	}

	if v == nil || other == nil {
		return false
	}

	if len(v) != len(other) {
		return false
	}

	var validators BoolValidators

	var otherValidators BoolValidators

	// Remove nils otherwise sort will panic.
	for _, validator := range v {
		if validator.Custom != nil {
			validators = append(validators, validator)
		}
	}

	// Remove nils otherwise sort will panic.
	for _, validator := range other {
		if validator.Custom != nil {
			otherValidators = append(otherValidators, validator)
		}
	}

	// Compare length after removing nils.
	if len(validators) != len(otherValidators) {
		return false
	}

	// SchemaDefinition is required by the spec JSON schema.
	sort.Slice(validators, func(i, j int) bool {
		return validators[i].Custom.SchemaDefinition < validators[j].Custom.SchemaDefinition
	})

	// SchemaDefinition is required by the spec JSON schema.
	sort.Slice(otherValidators, func(i, j int) bool {
		return otherValidators[i].Custom.SchemaDefinition < otherValidators[j].Custom.SchemaDefinition
	})

	for k, validator := range validators {
		if !validator.Equal(otherValidators[k]) {
			return false
		}
	}

	return true
}

// BoolValidator type defines type and function that provides validation
// functionality.
type BoolValidator struct {
	// Custom defines a schema definition, and optional imports.
	Custom *CustomValidator `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given BoolValidator equal.
func (v BoolValidator) Equal(other BoolValidator) bool {
	return v.Custom.Equal(other.Custom)
}
