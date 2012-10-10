// Copyright 2012 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// piconf can handle many basic types and groups of elements of one type.
//
// Basic types:
//
//   bool
//   int
//   int64
//   uint
//   uint64
//   float64
//   complex128
//   string
//
// Slices:
//
//   []byte
//   []int
//   []int64
//   []uint
//   []uint64
//   []float64
//   []string
//
// Map with elements of types defined above:
//
//   map
//
package main

import (
	"bytes"
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Valuer is the interface to the value stored in the database.
type Valuer interface {
	String() string
}

// == Common fields
//

// common represents the fields to add to every value type.
type common struct {
	// Last modifications
	LastUIDs  []int // user identifiers
	LastTimes []time.Time

	// Informatin about actual value
	UID  int
	Time time.Time

	Help map[string]string // language: text
	sync.RWMutex
}

func (c common) initCommon() {
	c.LastUIDs = make([]int, 0)
	c.LastTimes = make([]time.Time, 0)
	c.Help = make(map[string]string)
}

// Gethelp returns the text corresponding to the given language; if it is
// empty or it does not exist then it is used the language by default.
// It returns an empty string if the language does not exist.
func (c common) Gethelp(lang string) string {
	c.RLock()
	defer c.RUnlock()

	if lang != "" {
		if val, exist := c.Help[lang]; exist {
			return val
		}
	}
	if val, exist := c.Help[config.Lang]; exist {
		return val
	}
	return ""
}

// Sethelp adds a help text for the given language.
func (c common) Sethelp(lang, text string) {
	c.Lock()
	c.Help[lang] = text
	c.Unlock()
}

// == Basic types
//

// Bool represents a bool that implements the Valuer interface.
type Bool struct {
	Value      bool
	LastValues []bool
	common
}

// NewBool returns a new bool Value.
func NewBool() *Bool {
	v := new(Bool)
	v.LastValues = make([]bool, 0)
	v.initCommon()
	return v
}

// Get returns the value.
func (v *Bool) Get() bool {
	v.RLock()
	defer v.RUnlock()
	return v.Value
}

// Set sets the value, and saves the given user who is updating it; it also
// saves the time at setting.
// Before of to do setting, it is backed up the actual values, if any.
func (v *Bool) Set(value bool, uid int) {
	v.Lock()

	if !v.Time.IsZero() {
		v.LastTimes = append(v.LastTimes, v.Time)
		v.LastUIDs = append(v.LastUIDs, v.UID)
		v.LastValues = append(v.LastValues, v.Value)
	}
	v.UID, v.Value = uid, value
	v.Time = time.Now()

	v.Unlock()
	//<-updated
	once.Do(func() { db.Save(&Void{}, &Void{}) })
}

// String implements the Valuer interface.
func (v *Bool) String() string {
	v.RLock()
	defer v.RUnlock()
	return strconv.FormatBool(v.Value)
}

// == int Value
type Int struct {
	Value      int
	LastValues []int
	common
}

func NewInt() *Int {
	v := new(Int)
	v.LastValues = make([]int, 0)
	v.initCommon()
	return v
}

func (v *Int) Get() int {
	v.RLock()
	defer v.RUnlock()
	return v.Value
}

func (v *Int) Set(value int, uid int) {
	v.Lock()

	if !v.Time.IsZero() {
		v.LastTimes = append(v.LastTimes, v.Time)
		v.LastUIDs = append(v.LastUIDs, v.UID)
		v.LastValues = append(v.LastValues, v.Value)
	}
	v.UID, v.Value = uid, value
	v.Time = time.Now()

	v.Unlock()
	//<-updated
	once.Do(func() { db.Save(&Void{}, &Void{}) })
}

func (v *Int) String() string {
	v.RLock()
	defer v.RUnlock()
	return strconv.FormatInt(int64(v.Value), 10)
}

// == int64 Value
type Int64 struct {
	Value      int64
	LastValues []int64
	common
}

func NewInt64() *Int64 {
	v := new(Int64)
	v.LastValues = make([]int64, 0)
	v.initCommon()
	return v
}

func (v *Int64) Get() int64 {
	v.RLock()
	defer v.RUnlock()
	return v.Value
}

func (v *Int64) Set(value int64, uid int) {
	v.Lock()

	if !v.Time.IsZero() {
		v.LastTimes = append(v.LastTimes, v.Time)
		v.LastUIDs = append(v.LastUIDs, v.UID)
		v.LastValues = append(v.LastValues, v.Value)
	}
	v.UID, v.Value = uid, value
	v.Time = time.Now()

	v.Unlock()
	//<-updated
	once.Do(func() { db.Save(&Void{}, &Void{}) })
}

func (v *Int64) String() string {
	v.RLock()
	defer v.RUnlock()
	return strconv.FormatInt(v.Value, 10)
}

// == uint Value
type Uint struct {
	Value      uint
	LastValues []uint
	common
}

func NewUint() *Uint {
	v := new(Uint)
	v.LastValues = make([]uint, 0)
	v.initCommon()
	return v
}

func (v *Uint) Get() uint {
	v.RLock()
	defer v.RUnlock()
	return v.Value
}

func (v *Uint) Set(value uint, uid int) {
	v.Lock()

	if !v.Time.IsZero() {
		v.LastTimes = append(v.LastTimes, v.Time)
		v.LastUIDs = append(v.LastUIDs, v.UID)
		v.LastValues = append(v.LastValues, v.Value)
	}
	v.UID, v.Value = uid, value
	v.Time = time.Now()

	v.Unlock()
	//<-updated
	once.Do(func() { db.Save(&Void{}, &Void{}) })
}

func (v *Uint) String() string {
	v.RLock()
	defer v.RUnlock()
	return strconv.FormatUint(uint64(v.Value), 10)
}

// == uint64 Value
type Uint64 struct {
	Value      uint64
	LastValues []uint64
	common
}

func NewUint64() *Uint64 {
	v := new(Uint64)
	v.LastValues = make([]uint64, 0)
	v.initCommon()
	return v
}

func (v *Uint64) Get() uint64 {
	v.RLock()
	defer v.RUnlock()
	return v.Value
}

func (v *Uint64) Set(value uint64, uid int) {
	v.Lock()

	if !v.Time.IsZero() {
		v.LastTimes = append(v.LastTimes, v.Time)
		v.LastUIDs = append(v.LastUIDs, v.UID)
		v.LastValues = append(v.LastValues, v.Value)
	}
	v.UID, v.Value = uid, value
	v.Time = time.Now()

	v.Unlock()
	//<-updated
	once.Do(func() { db.Save(&Void{}, &Void{}) })
}

func (v *Uint64) String() string {
	v.RLock()
	defer v.RUnlock()
	return strconv.FormatUint(v.Value, 10)
}

// == float64 Value
type Float64 struct {
	Value      float64
	LastValues []float64
	common
}

func NewFloat64() *Float64 {
	v := new(Float64)
	v.LastValues = make([]float64, 0)
	v.initCommon()
	return v
}

func (v *Float64) Get() float64 {
	v.RLock()
	defer v.RUnlock()
	return v.Value
}

func (v *Float64) Set(value float64, uid int) {
	v.Lock()

	if !v.Time.IsZero() {
		v.LastTimes = append(v.LastTimes, v.Time)
		v.LastUIDs = append(v.LastUIDs, v.UID)
		v.LastValues = append(v.LastValues, v.Value)
	}
	v.UID, v.Value = uid, value
	v.Time = time.Now()

	v.Unlock()
	//<-updated
	once.Do(func() { db.Save(&Void{}, &Void{}) })
}

func (v *Float64) String() string {
	v.RLock()
	defer v.RUnlock()
	return strconv.FormatFloat(v.Value, 'g', -1, 64)
}

// == complex128 Value
type Complex128 struct {
	Value      complex128
	LastValues []complex128
	common
}

func NewComplex128() *Complex128 {
	v := new(Complex128)
	v.LastValues = make([]complex128, 0)
	v.initCommon()
	return v
}

func (v *Complex128) Get() complex128 {
	v.RLock()
	defer v.RUnlock()
	return v.Value
}

func (v *Complex128) Set(value complex128, uid int) {
	v.Lock()

	if !v.Time.IsZero() {
		v.LastTimes = append(v.LastTimes, v.Time)
		v.LastUIDs = append(v.LastUIDs, v.UID)
		v.LastValues = append(v.LastValues, v.Value)
	}
	v.UID, v.Value = uid, value
	v.Time = time.Now()

	v.Unlock()
	//<-updated
	once.Do(func() { db.Save(&Void{}, &Void{}) })
}

func (v *Complex128) String() string {
	v.RLock()
	defer v.RUnlock()
	return fmt.Sprintf("[%v, %v]", real(v.Value), imag(v.Value)) // for JSON
}

// == string Value
type String struct {
	Value      string
	LastValues []string
	common
}

func NewString() *String {
	v := new(String)
	v.LastValues = make([]string, 0)
	v.initCommon()
	return v
}

func (v *String) Get() string {
	v.RLock()
	defer v.RUnlock()
	return v.Value
}

func (v *String) Set(value string, uid int) {
	v.Lock()

	if !v.Time.IsZero() {
		v.LastTimes = append(v.LastTimes, v.Time)
		v.LastUIDs = append(v.LastUIDs, v.UID)
		v.LastValues = append(v.LastValues, v.Value)
	}
	v.UID, v.Value = uid, value
	v.Time = time.Now()

	v.Unlock()
	//<-updated
	once.Do(func() { db.Save(&Void{}, &Void{}) })
}

func (v *String) String() string {
	v.RLock()
	defer v.RUnlock()
	return strconv.Quote(v.Value)
}

// == Composite types
//

// == []byte Value
type RawBytes struct {
	Value      []byte
	LastValues [][]byte
	common
}

func NewRawBytes() *RawBytes {
	v := new(RawBytes)
	v.Value = make([]byte, 0)
	v.LastValues = make([][]byte, 0)
	v.initCommon()
	return v
}

func (v *RawBytes) Get() []byte {
	v.RLock()
	defer v.RUnlock()
	return v.Value
}

func (v *RawBytes) Set(value []byte, uid int) {
	v.Lock()

	if !v.Time.IsZero() {
		v.LastTimes = append(v.LastTimes, v.Time)
		v.LastUIDs = append(v.LastUIDs, v.UID)
		v.LastValues = append(v.LastValues, v.Value)
	}
	v.UID, v.Value = uid, value
	v.Time = time.Now()

	v.Unlock()
	//<-updated
	once.Do(func() { db.Save(&Void{}, &Void{}) })
}

func (v *RawBytes) String() string {
	v.RLock()
	defer v.RUnlock()
	return fmt.Sprintf("%v", v.Value)
}

// == []int Value
type IntSlice struct {
	Value      []int
	LastValues [][]int
	common
}

func NewIntSlice() *IntSlice {
	v := new(IntSlice)
	v.Value = make([]int, 0)
	v.LastValues = make([][]int, 0)
	v.initCommon()
	return v
}

func (v *IntSlice) Get() []int {
	v.RLock()
	defer v.RUnlock()
	return v.Value
}

func (v *IntSlice) Set(value []int, uid int) {
	v.Lock()

	if !v.Time.IsZero() {
		v.LastTimes = append(v.LastTimes, v.Time)
		v.LastUIDs = append(v.LastUIDs, v.UID)
		v.LastValues = append(v.LastValues, v.Value)
	}
	v.UID, v.Value = uid, value
	v.Time = time.Now()

	v.Unlock()
	//<-updated
	once.Do(func() { db.Save(&Void{}, &Void{}) })
}

func (v *IntSlice) String() string {
	var b bytes.Buffer
	first := true

	fmt.Fprintf(&b, "[")

	v.RLock()
	for _, val := range v.Value {
		if !first {
			fmt.Fprintf(&b, ", ")
		}
		fmt.Fprintf(&b, "%d", val)
		first = false
	}
	v.RUnlock()

	fmt.Fprintf(&b, "]")
	return b.String()
}

// == []int64 Value
type Int64Slice struct {
	Value      []int64
	LastValues [][]int64
	common
}

func NewInt64Slice() *Int64Slice {
	v := new(Int64Slice)
	v.Value = make([]int64, 0)
	v.LastValues = make([][]int64, 0)
	v.initCommon()
	return v
}

func (v *Int64Slice) Get() []int64 {
	v.RLock()
	defer v.RUnlock()
	return v.Value
}

func (v *Int64Slice) Set(value []int64, uid int) {
	v.Lock()

	if !v.Time.IsZero() {
		v.LastTimes = append(v.LastTimes, v.Time)
		v.LastUIDs = append(v.LastUIDs, v.UID)
		v.LastValues = append(v.LastValues, v.Value)
	}
	v.UID, v.Value = uid, value
	v.Time = time.Now()

	v.Unlock()
	//<-updated
	once.Do(func() { db.Save(&Void{}, &Void{}) })
}

func (v *Int64Slice) String() string {
	var b bytes.Buffer
	first := true

	fmt.Fprintf(&b, "[")

	v.RLock()
	for _, val := range v.Value {
		if !first {
			fmt.Fprintf(&b, ", ")
		}
		fmt.Fprintf(&b, "%v", val)
		first = false
	}
	v.RUnlock()

	fmt.Fprintf(&b, "]")
	return b.String()
}

// == []uint Value
type UintSlice struct {
	Value      []uint
	LastValues [][]uint
	common
}

func NewUintSlice() *UintSlice {
	v := new(UintSlice)
	v.Value = make([]uint, 0)
	v.LastValues = make([][]uint, 0)
	v.initCommon()
	return v
}

func (v *UintSlice) Get() []uint {
	v.RLock()
	defer v.RUnlock()
	return v.Value
}

func (v *UintSlice) Set(value []uint, uid int) {
	v.Lock()

	if !v.Time.IsZero() {
		v.LastTimes = append(v.LastTimes, v.Time)
		v.LastUIDs = append(v.LastUIDs, v.UID)
		v.LastValues = append(v.LastValues, v.Value)
	}
	v.UID, v.Value = uid, value
	v.Time = time.Now()

	v.Unlock()
	//<-updated
	once.Do(func() { db.Save(&Void{}, &Void{}) })
}

func (v *UintSlice) String() string {
	var b bytes.Buffer
	first := true

	fmt.Fprintf(&b, "[")

	v.RLock()
	for _, val := range v.Value {
		if !first {
			fmt.Fprintf(&b, ", ")
		}
		fmt.Fprintf(&b, "%v", val)
		first = false
	}
	v.RUnlock()

	fmt.Fprintf(&b, "]")
	return b.String()
}

// == []uint64 Value
type Uint64Slice struct {
	Value      []uint64
	LastValues [][]uint64
	common
}

func NewUint64Slice() *Uint64Slice {
	v := new(Uint64Slice)
	v.Value = make([]uint64, 0)
	v.LastValues = make([][]uint64, 0)
	v.initCommon()
	return v
}

func (v *Uint64Slice) Get() []uint64 {
	v.RLock()
	defer v.RUnlock()
	return v.Value
}

func (v *Uint64Slice) Set(value []uint64, uid int) {
	v.Lock()

	if !v.Time.IsZero() {
		v.LastTimes = append(v.LastTimes, v.Time)
		v.LastUIDs = append(v.LastUIDs, v.UID)
		v.LastValues = append(v.LastValues, v.Value)
	}
	v.UID, v.Value = uid, value
	v.Time = time.Now()

	v.Unlock()
	//<-updated
	once.Do(func() { db.Save(&Void{}, &Void{}) })
}

func (v *Uint64Slice) String() string {
	var b bytes.Buffer
	first := true

	fmt.Fprintf(&b, "[")

	v.RLock()
	for _, val := range v.Value {
		if !first {
			fmt.Fprintf(&b, ", ")
		}
		fmt.Fprintf(&b, "%v", val)
		first = false
	}
	v.RUnlock()

	fmt.Fprintf(&b, "]")
	return b.String()
}

// == []Float64 Value
type Float64Slice struct {
	Value      []float64
	LastValues [][]float64
	common
}

func NewFloat64Slice() *Float64Slice {
	v := new(Float64Slice)
	v.Value = make([]float64, 0)
	v.LastValues = make([][]float64, 0)
	v.initCommon()
	return v
}

func (v *Float64Slice) Get() []float64 {
	v.RLock()
	defer v.RUnlock()
	return v.Value
}

func (v *Float64Slice) Set(value []float64, uid int) {
	v.Lock()

	if !v.Time.IsZero() {
		v.LastTimes = append(v.LastTimes, v.Time)
		v.LastUIDs = append(v.LastUIDs, v.UID)
		v.LastValues = append(v.LastValues, v.Value)
	}
	v.UID, v.Value = uid, value
	v.Time = time.Now()

	v.Unlock()
	//<-updated
	once.Do(func() { db.Save(&Void{}, &Void{}) })
}

func (v *Float64Slice) String() string {
	var b bytes.Buffer
	first := true

	fmt.Fprintf(&b, "[")

	v.RLock()
	for _, val := range v.Value {
		if !first {
			fmt.Fprintf(&b, ", ")
		}
		fmt.Fprintf(&b, "%v", val)
		first = false
	}
	v.RUnlock()

	fmt.Fprintf(&b, "]")
	return b.String()
}

// == []string Value
type StringSlice struct {
	Value      []string
	LastValues [][]string
	common
}

func NewStringSlice() *StringSlice {
	v := new(StringSlice)
	v.Value = make([]string, 0)
	v.LastValues = make([][]string, 0)
	v.initCommon()
	return v
}

func (v *StringSlice) Get() []string {
	v.RLock()
	defer v.RUnlock()
	return v.Value
}

func (v *StringSlice) Set(value []string, uid int) {
	v.Lock()

	if !v.Time.IsZero() {
		v.LastTimes = append(v.LastTimes, v.Time)
		v.LastUIDs = append(v.LastUIDs, v.UID)
		v.LastValues = append(v.LastValues, v.Value)
	}
	v.UID, v.Value = uid, value
	v.Time = time.Now()

	v.Unlock()
	//<-updated
	once.Do(func() { db.Save(&Void{}, &Void{}) })
}

func (v *StringSlice) String() string {
	var b bytes.Buffer
	first := true

	fmt.Fprintf(&b, "[")

	v.RLock()
	for _, val := range v.Value {
		if !first {
			fmt.Fprintf(&b, ", ")
		}
		fmt.Fprintf(&b, "%q", val)
		first = false
	}
	v.RUnlock()

	fmt.Fprintf(&b, "]")
	return b.String()
}

// == map Value

// Map represents a map whose keys in Value are the variable names of the
// configuration and the values implement the Valuer interface.
//
// The Map is the container to hold the configuration of every program, or
// program's section.
type Map struct {
	IsMain bool   // whether the configuration is for a program or section
	Name   string // program name or configuration's section
	Ver    string // program version
	Value  map[string]Valuer
	sync.RWMutex
}

// NewMap defines a map with the specified name.
func NewMap(name string, isMain bool) *Map {
	return &Map{
		IsMain: isMain,
		Name:   name,
		Value:  make(map[string]Valuer),
	}
}

// Get returns the value of given key.
func (v *Map) Get(key string) Valuer {
	v.RLock()
	defer v.RUnlock()
	return v.Value[key]
}

// Set sets value in key.
func (v *Map) Set(key string, val Valuer) {
	v.Lock()
	v.Value[key] = val

	v.Unlock()
	//<-updated
	once.Do(func() { db.Save(&Void{}, &Void{}) })
}

// Setversion sets the program version.
func (v *Map) Setversion(ver string) *Map {
	v.Lock()
	v.Ver = ver

	v.Unlock()
	//<-updated
	once.Do(func() { db.Save(&Void{}, &Void{}) })
	return v
}

// String implements the Valuer interface.
func (v *Map) String() string {
	var b bytes.Buffer
	first := true

	fmt.Fprintf(&b, "{")

	v.RLock()
	for key, value := range v.Value {
		if !first {
			fmt.Fprintf(&b, ", ")
		}
		fmt.Fprintf(&b, "%q: %v", key, value)
		first = false
	}
	v.RUnlock()

	fmt.Fprintf(&b, "}")
	return b.String()
}
