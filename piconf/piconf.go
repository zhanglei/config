// Copyright 2012 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package piconf 
//
package piconf

import (
	"fmt"
	"log"
	"sync"
)

// SaveData saves data into a file encoded to gob.
func SaveData(file string, set ExtensionSet) error {
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("SaveData: %s", err)
	}
	defer f.Close()

	enc := gob.NewEncoder(f)

	if err = enc.Encode(_HEADER); err != nil {
		goto _error
	}

	// Save the extensions at the beginning to can load them.
	if err = enc.Encode(set); err != nil {
		goto _error
	}

	for i, s := range allSetExts {
		if set&s != 0 {
			if err = enc.Encode(allExts[i]); err != nil {
				goto _error
			}
		}
	}

	return nil
_error:
	return fmt.Errorf("SaveData: %s", err)
}
