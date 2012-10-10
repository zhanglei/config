// Copyright 2012 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

// == Basic

func TestBool(t *testing.T) {
	v := NewBool()
	fname := "NewBool"
	want := true

	v.Set(want, 0)

	if got := v.Value; got != want {
		t.Errorf("%s.Set got %v, want %v", fname, got, want)
	}
	if got := v.Get(); got != want {
		t.Errorf("%s.Get got %v, want %v", fname, got, want)
	}
	if got := v.String(); got != fmt.Sprintf("%v", want) {
		t.Errorf("%s.String got %q, want %q", fname, got, want)
	}
}

func TestInt(t *testing.T) {
	v := NewInt()
	fname := "NewInt"
	want := -99

	v.Set(want, 0)

	if got := v.Value; got != want {
		t.Errorf("%s.Set got %v, want %v", fname, got, want)
	}
	if got := v.Get(); got != want {
		t.Errorf("%s.Get got %v, want %v", fname, got, want)
	}
	if got := v.String(); got != fmt.Sprintf("%v", want) {
		t.Errorf("%s.String got %q, want %q", fname, got, want)
	}
}

func TestInt64(t *testing.T) {
	v := NewInt64()
	fname := "NewInt64"
	want := int64(-99)

	v.Set(want, 0)

	if got := v.Value; got != want {
		t.Errorf("%s.Set got %v, want %v", fname, got, want)
	}
	if got := v.Get(); got != want {
		t.Errorf("%s.Get got %v, want %v", fname, got, want)
	}
	if got := v.String(); got != fmt.Sprintf("%v", want) {
		t.Errorf("%s.String got %q, want %q", fname, got, want)
	}
}

func TestUint(t *testing.T) {
	v := NewUint()
	fname := "NewUint"
	want := uint(99)

	v.Set(want, 0)

	if got := v.Value; got != want {
		t.Errorf("%s.Set got %v, want %v", fname, got, want)
	}
	if got := v.Get(); got != want {
		t.Errorf("%s.Get got %v, want %v", fname, got, want)
	}
	if got := v.String(); got != fmt.Sprintf("%v", want) {
		t.Errorf("%s.String got %q, want %q", fname, got, want)
	}
}

func TestUint64(t *testing.T) {
	v := NewUint64()
	fname := "NewUint64"
	want := uint64(99)

	v.Set(want, 0)

	if got := v.Value; got != want {
		t.Errorf("%s.Set got %v, want %v", fname, got, want)
	}
	if got := v.Get(); got != want {
		t.Errorf("%s.Get got %v, want %v", fname, got, want)
	}
	if got := v.String(); got != fmt.Sprintf("%v", want) {
		t.Errorf("%s.String got %q, want %q", fname, got, want)
	}
}

func TestFloat64(t *testing.T) {
	v := NewFloat64()
	fname := "NewFloat64"
	want := -99.99

	v.Set(want, 0)

	if got := v.Value; got != want {
		t.Errorf("%s.Set got %v, want %v", fname, got, want)
	}
	if got := v.Get(); got != want {
		t.Errorf("%s.Get got %v, want %v", fname, got, want)
	}
	if got := v.String(); got != fmt.Sprintf("%v", want) {
		t.Errorf("%s.String got %q, want %q", fname, got, want)
	}
}

func TestComplex128(t *testing.T) {
	v := NewComplex128()
	fname := "NewComplex128"
	want := 128.12 + 3i

	v.Set(want, 0)

	if got := v.Value; got != want {
		t.Errorf("%s.Set got %v, want %v", fname, got, want)
	}
	if got := v.Get(); got != want {
		t.Errorf("%s.Get got %v, want %v", fname, got, want)
	}
	if got := v.String(); got != fmt.Sprintf("[%v, %v]", real(want), imag(want)) {
		t.Errorf("%s.String got %v, want %v", fname, got, want)
	}
}

func TestString(t *testing.T) {
	v := NewString()
	fname := "NewString"
	want := "foo"

	v.Set(want, 0)

	if got := v.Value; got != want {
		t.Errorf("%s.Set got %v, want %v", fname, got, want)
	}
	if got := v.Get(); got != want {
		t.Errorf("%s.Get got %v, want %v", fname, got, want)
	}
	if got := v.String(); got != fmt.Sprintf("%q", want) {
		t.Errorf("%s.String got %q, want %q", fname, got, want)
	}
}

// == Slices

func TestRawBytes(t *testing.T) {
	v := NewRawBytes()
	fname := "NewRawBytes"
	want := []byte{1, 2}

	v.Set(want, 0)

	got := v.Value
	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("%s.Set got %v, want %v", fname, got, want)
	}
	got = v.Get()
	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("%s.Get got %v, want %v", fname, got, want)
	}
	if got := v.String(); got != fmt.Sprintf("%v", want) {
		t.Errorf("%s.String got %q, want %q", fname, got, want)
	}
}

func TestIntSlice(t *testing.T) {
	v := NewIntSlice()
	fname := "NewIntSlice"
	want := []int{-1, 2}

	v.Set(want, 0)

	got := v.Value
	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("%s.Set got %v, want %v", fname, got, want)
	}
	got = v.Get()
	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("%s.Get got %v, want %v", fname, got, want)
	}
}

func TestInt64Slice(t *testing.T) {
	v := NewInt64Slice()
	fname := "NewInt64Slice"
	want := []int64{-1, 2}

	v.Set(want, 0)

	got := v.Value
	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("%s.Set got %v, want %v", fname, got, want)
	}
	got = v.Get()
	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("%s.Get got %v, want %v", fname, got, want)
	}
}

func TestUintSlice(t *testing.T) {
	v := NewUintSlice()
	fname := "NewUintSlice"
	want := []uint{1, 2}

	v.Set(want, 0)

	got := v.Value
	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("%s.Set got %v, want %v", fname, got, want)
	}
	got = v.Get()
	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("%s.Get got %v, want %v", fname, got, want)
	}
}

func TestUint64Slice(t *testing.T) {
	v := NewUint64Slice()
	fname := "NewUint64Slice"
	want := []uint64{1, 2}

	v.Set(want, 0)

	got := v.Value
	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("%s.Set got %v, want %v", fname, got, want)
	}
	got = v.Get()
	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("%s.Get got %v, want %v", fname, got, want)
	}
}

func TestFloat64Slice(t *testing.T) {
	v := NewFloat64Slice()
	fname := "NewFloat64Slice"
	want := []float64{-1.14, 2.38}

	v.Set(want, 0)

	got := v.Value
	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("%s.Set got %v, want %v", fname, got, want)
	}
	got = v.Get()
	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("%s.Get got %v, want %v", fname, got, want)
	}
}

func TestStringSlice(t *testing.T) {
	v := NewStringSlice()
	fname := "NewStringSlice"
	want := []string{"foo", "bar"}

	v.Set(want, 0)

	got := v.Value
	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("%s.Set got %v, want %v", fname, got, want)
	}
	got = v.Get()
	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("%s.Get got %v, want %v", fname, got, want)
	}
}

// == Map

func TestMap(t *testing.T) {
	name := "program piconf"
	ver := "1.0.1"

	cfg := NewMap(name, true)
	cfg.Setversion(ver)

	if cfg.Name != name {
		t.Errorf("NewMap got %v, want %v", cfg.Name, name)
	}
	if cfg.Ver != ver {
		t.Errorf("%s.Setversion got %v, want %v", cfg.Ver, ver)
	}

	boolKey, boolValue := "car", true
	{
		v := NewBool()
		v.Set(boolValue, 0)
		cfg.Set(boolKey, v)
	}

	intKey, intValue := "axis", -11
	{
		v := NewInt()
		v.Set(intValue, 0)
		cfg.Set(intKey, v)
	}

	uintKey, uintValue := "age", uint(11)
	{
		v := NewUint()
		v.Set(uintValue, 0)
		cfg.Set(uintKey, v)
	}

	floatkey, floatValue := "price", 11.11
	{
		v := NewFloat64()
		v.Set(floatValue, 0)
		cfg.Set(floatkey, v)
	}

	complexKey, complexValue := "complex", 23.12+12i
	{
		v := NewComplex128()
		v.Set(complexValue, 0)
		cfg.Set(complexKey, v)
	}

	strKey, strValue := "name", "foo"
	{
		v := NewString()
		v.Set(strValue, 0)
		cfg.Set(strKey, v)
	}

	if x := cfg.Value[boolKey].(*Bool).Value; x != boolValue {
		t.Errorf("cfg[%q] got %v, want %v", boolKey, x, boolValue)
	}
	if x := cfg.Value[intKey].(*Int).Value; x != intValue {
		t.Errorf("cfg[%q] got %v, want %v", intKey, x, intValue)
	}
	if x := cfg.Value[uintKey].(*Uint).Value; x != uintValue {
		t.Errorf("cfg[%q] got %v, want %v", uintKey, x, uintValue)
	}
	if x := cfg.Value[floatkey].(*Float64).Value; x != floatValue {
		t.Errorf("cfg[%q] got %v, want %v", floatkey, x, floatValue)
	}
	if x := cfg.Value[complexKey].(*Complex128).Value; x != complexValue {
		t.Errorf("cfg[%q] got %v, want %v", complexKey, x, complexValue)
	}
	if x := cfg.Value[strKey].(*String).Value; x != strValue {
		t.Errorf("cfg[%q] got %v, want %v", strKey, x, strValue)
	}

	// Check parsing to JSON.

	cfgString := cfg.String()
	var cfgJSON interface{}

	if err := json.Unmarshal([]byte(cfgString), &cfgJSON); err != nil {
		t.Errorf("cfg.String() is not valid JSON: %v", err)
	}

	m, ok := cfgJSON.(map[string]interface{})
	if !ok {
		t.Error("cfg.String() did not get a map")
	}

	car := m["car"]
	x, ok := car.(bool)
	if !ok {
		t.Error("car.(bool) is not a boolean")
	}
	if x != true {
		t.Errorf("car got %v, want %v", x, true)
	}
}
