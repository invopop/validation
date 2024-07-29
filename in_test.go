// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIn(t *testing.T) {
	var v = 1
	var v2 *int
	tests := []struct {
		tag    string
		values []interface{}
		value  interface{}
		err    string
	}{
		{"t0", []interface{}{1, 2}, 0, ""},
		{"t1", []interface{}{1, 2}, 1, ""},
		{"t2", []interface{}{1, 2}, 2, ""},
		{"t3", []interface{}{1, 2}, 3, "must be a valid value"},
		{"t4", []interface{}{}, 3, "must be a valid value"},
		{"t5", []interface{}{1, 2}, "1", "must be a valid value"},
		{"t6", []interface{}{1, 2}, &v, ""},
		{"t7", []interface{}{1, 2}, v2, ""},
		{"t8", []interface{}{[]byte{1}, 1, 2}, []byte{1}, ""},
	}

	for _, test := range tests {
		r := In(test.values...)
		err := r.Validate(test.value)
		assertError(t, test.err, err, test.tag)
	}
}

func TestIn_Generics(t *testing.T) {
	var v = "a"
	var v2 *string
	tests := []struct {
		tag    string
		values []string
		value  any
		err    string
	}{
		{"t0", []string{"a", "b", "c"}, "a", ""},
		{"t1", []string{"a", "b", "c"}, "d", "must be a valid value"},
		{"t2", []string{"a", "b"}, &v, ""},
		{"t3", []string{"a", "b"}, v2, ""},
		{"t4", []string{"a", "b"}, nil, ""},
		{"t5", []string{}, "a", "must be a valid value"},
	}

	for _, test := range tests {
		r := In(test.values...)
		err := r.Validate(test.value)
		assertError(t, test.err, err, test.tag)
	}
}

func Test_InRule_Error(t *testing.T) {
	r := In(1, 2, 3)
	val := 4
	assert.Equal(t, "must be a valid value", r.Validate(&val).Error())
	r = r.Error("123")
	assert.Equal(t, "123", r.err.Message())
}

func TestInRule_ErrorObject(t *testing.T) {
	r := In(1, 2, 3)

	err := NewError("code", "abc")
	r = r.ErrorObject(err)

	assert.Equal(t, err, r.err)
	assert.Equal(t, err.Code(), r.err.Code())
	assert.Equal(t, err.Message(), r.err.Message())
}

func TestValidateAgainstList(t *testing.T) {
	optionsList := []string{"a", "b", "c"}
	// I've always felt that this was dangerous
	assert.ErrorContains(t, Validate("a", In(optionsList)), "must be a valid value")
	// This is better, no need to convert to interface
	assert.NoError(t, Validate("a", In(optionsList...)))
}
