// Copyright 2012 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package defconf sets the configuration by default for piconf.
package defconf

const (
	GREETING = "piconf server"

	// Port for the web user interface
	HTTP_PORT = 7070

	// TCP
	HOST = "127.0.0.1"
	PORT = 707

	// Unix socket
	SOCKET_FILE = "/tmp/conf" // "/dev/conf"
)
