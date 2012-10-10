// Copyright 2012 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package shconf implements a shell-variable configuration parser.
//
// The configuration file consists on entries with the format "key"="value".
// The comments are indicated by "#" at the beginning of a line and upon the keys;
// the first comment is about the file.
package shconf

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"strconv"
	"sync"
	"unicode"

	"github.com/kless/shout/file"
)

var (
	bComment = []byte{'#'}
	bEmpty   = []byte{}
	bEqual   = []byte{'='}
	bDQuote  = []byte{'"'}
)

// A Config represents the configuration.
type Config struct {
	filename string
	comment  map[int][]string  // id: []{comment, key...}; id 1 is for main comment.
	data     map[string]string // key: value
	offset   map[string]int64  // key: offset; for editing.
	sync.RWMutex
}

// ParseFile creates a new Config and parses the file configuration from the
// named file.
func ParseFile(name string) (*Config, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		file.Name(),
		make(map[int][]string),
		make(map[string]string),
		make(map[string]int64),
		sync.RWMutex{},
	}
	cfg.Lock()
	defer cfg.Unlock()
	defer file.Close()

	var comment bytes.Buffer
	buf := bufio.NewReader(file)

	for nComment, off := 0, int64(1); ; {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		if bytes.Equal(line, bEmpty) {
			continue
		}

		off += int64(len(line))

		if bytes.HasPrefix(line, bComment) {
			line = bytes.TrimLeft(line, "#")
			line = bytes.TrimLeftFunc(line, unicode.IsSpace)
			comment.Write(line)
			comment.WriteByte('\n')
			continue
		}
		if comment.Len() != 0 {
			cfg.comment[nComment] = []string{comment.String()}
			comment.Reset()
			nComment++
		}

		val := bytes.SplitN(line, bEqual, 2)
		if bytes.HasPrefix(val[1], bDQuote) {
			val[1] = bytes.Trim(val[1], `"`)
		}

		key := string(val[0])
		cfg.comment[nComment-1] = append(cfg.comment[nComment-1], key)
		cfg.data[key] = string(val[1])
		cfg.offset[key] = off
	}
	return cfg, nil
}

// Bool returns the boolean value for a given key.
func (c *Config) Bool(key string) (bool, error) {
	return strconv.ParseBool(c.data[key])
}

// Int returns the integer value for a given key.
func (c *Config) Int(key string) (int, error) {
	return strconv.Atoi(c.data[key])
}

// Float returns the float value for a given key.
func (c *Config) Float(key string) (float64, error) {
	return strconv.ParseFloat(c.data[key], 64)
}

// String returns the string value for a given key.
func (c *Config) String(key string) string {
	return c.data[key]
}

// WriteValue writes a new value for key.
func (c *Config) WriteValue(key, value string) error {
	c.Lock()
	defer c.Unlock()

	if _, found := c.data[key]; !found {
		return errors.New("key not found: " + key)
	}

	replAt := []file.ReplacerAtLine{
		{key + "=", "=.*", "=" + value},
	}

	if err := file.ReplaceAtLine(c.filename, replAt); err != nil {
		return err
	}
	c.data[key] = value
	return nil
}
