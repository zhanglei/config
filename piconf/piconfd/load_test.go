// Copyright 2012 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"path"
	"strings"
	"testing"
)

var (
	expressionErrs = []string{
		"err-func.cfg",
		"err-operation.cfg",
	}

	genericErrs = []string{
		"err-mconst.cfg",
		"err-mvar.cfg",
	}
)

func TestLoad(t *testing.T) {
	filename := "../testdata/ok.cfg"
	err := Load(filename)

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	for _, f := range genericErrs {
		file := path.Join("../testdata", f)
		if err = Load(file); err == nil || !strings.HasPrefix(err.Error(), "configuration in") {
			t.Errorf("expected error on file %q", f)
		}
	}

	for _, f := range expressionErrs {
		file := path.Join("../testdata", f)
		if err = Load(file); err == nil || !strings.HasPrefix(err.Error(), "expression no") {
			t.Errorf("expected error on file %q", file)
		}
	}

	filename = "../testdata/err-doc.cfg"
	if err = Load(filename); err == nil || !strings.Contains(err.Error(), "documentation") {
		t.Errorf("expected error on file %q", filename)
	}
}
