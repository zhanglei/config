// Copyright 2012 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package shconf

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var testdata = []struct {
	k string
	v string
}{
	{"Bool", "true"},
	{"Int", "2"},
	{"Float", "3.3"},
	{"String", "small"},
}

func Test(t *testing.T) {
	// == Create temporary file
	file, _ := ioutil.TempFile("", "test")

	buf := bufio.NewWriter(file)
	buf.WriteString("# main comment\n\n")
	buf.WriteString(fmt.Sprintf("%s=%s\n", testdata[0].k, testdata[0].v))
	buf.WriteString(fmt.Sprintf("%s=%s\n\n", testdata[1].k, testdata[1].v))
	buf.WriteString("# Another comment\n")
	buf.WriteString(fmt.Sprintf("%s=%s\n", testdata[2].k, testdata[2].v))
	buf.WriteString(fmt.Sprintf("%s=%s\n", testdata[3].k, testdata[3].v))
	buf.Flush()
	file.Close()

	// == Parser
	cfg, err := ParseFile(file.Name())
	if err != nil {
		goto _exit
	}

	for k, _ := range cfg.data {
		switch k {
		case "Bool":
			_, err = cfg.Bool(k)
		case "Int":
			_, err = cfg.Int(k)
		case "Float":
			_, err = cfg.Float(k)
		case "String":
			_ = cfg.String(k)
		}
		if err != nil {
			t.Errorf("%q got wrong value", k)
		}
	}

	// == Editing
	if err = cfg.WriteValue("String", "big"); err != nil {
		goto _exit
	}
	if cfg.data["String"] != "big" {
		t.Errorf("value %q could not be set in key %q", "big", "String")
	}

	if err = cfg.WriteValue("Not", ""); err == nil {
		t.Errorf("key %q should not exist", "Not")
	} else {
		err = nil
	}

_exit:
	files, _ := filepath.Glob(file.Name() + "*")
	for _, v := range files {
		os.Remove(v)
	}

	if err != nil {
		t.Fatal(err)
	}
}
