// Copyright 2012 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

/*
Command piconfd is a system configuration server.

Advantages

Centralized system: it is more easy to backup/restore, and modify the same
values in different servers.

Editing: from a web interface with support for localization and validation.

Localization: can show the help messages for each key in different languages
(via web UI).

Validation: can validate the values at creating or updating a configuration.

Revision: the initial configuration for an application is not overwritten.

Log: logs all changes done in the configuration files.
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"os/user"
//	"reflect"
	"strconv"
	"sync"
	"time"

	"github.com/kless/netutil/start"
	"github.com/kless/netutil/watch"
	"github.com/kless/piconf/defconf"
)

// Time to save database to disk after of a change.
const TIMEOUT_SAVE = 30 * time.Second

// Configuration.
var config struct {
	Lang   string // language by default
	//Server *netutil.Config
}

// Database
var (
	HEADER = [3]byte{'7', '0', '7'}
	db     = Conf{m: make(map[int]map[string]*Map)}
	once   sync.Once // to write to disk
)

// == Errors

// SameConfigError is returned by Conf.Add when the command configuration for
// the given user ID has been already added.
type SameConfigError struct {
	uid     int
	cmdPath string
}

func (e SameConfigError) Error() string {
	uidString := strconv.Itoa(e.uid)

	user_, err := user.LookupId(uidString)
	if err == nil {
		return "user " + user_.Username + " has program: " + e.cmdPath
	}
	return "userid " + uidString + " has program: " + e.cmdPath
}

// SameConfigError is returned by Conf.Add when the command configuration for
// the given user ID does not exist.
type UnknownConfigError struct {
	uid     int
	cmdPath string
}

func (e UnknownConfigError) Error() string {
	uidString := strconv.Itoa(e.uid)

	user_, err := user.LookupId(uidString)
	if err == nil {
		return "user " + user_.Username + " has not program: " + e.cmdPath
	}
	return "userid " + uidString + " has not program: " + e.cmdPath
}
// ==

type Void struct{}

// Conf represents the configuration per user.
// The map's key indicates the user identifier; -1 is used for all users so
// the global configuration.
type Conf struct {
	sync.RWMutex
	m map[int]map[string]*Map // user id: program path: configuration data
}

type ArgsConf struct {
	uid     int
	cmdPath string
	m       *Map
}

// Add registers the user's configuration of a program installed in the given path.
func (c Conf) Add(args ArgsConf, reply *Void) error {
	c.Lock()

	if _, exist := c.m[args.uid][args.cmdPath]; exist {
		c.Unlock()

		err := &SameConfigError{args.uid, args.cmdPath}
		log.Println(err) // TODO: better logging
		return err
	}
	c.m[args.uid] = make(map[string]*Map)
	c.m[args.uid][args.cmdPath] = args.m

	c.Unlock()
	return nil
}

// Get returns the Map for the user id and command path given.
func (c Conf) Get(args ArgsConf, m *Map) error {
	c.RLock()

	m, exist := c.m[args.uid][args.cmdPath]
	if !exist {
		c.RUnlock()

		err := &UnknownConfigError{args.uid, args.cmdPath}
		log.Println(err)
		return err
	}

	c.RUnlock()
	return nil
}

// Save writes database in memory to disk after of the given time in TIMEOUT_SAVE.
func (c Conf) Save(*Void, *Void) error {
	go func() {
		//select {
		//case <-time.After(TIMEOUT_SAVE):
		//}
		<-time.After(TIMEOUT_SAVE)

		c.Lock()

		/*if err := ioutil.WriteFile("foo", []byte(s), 0666); err != nil {
			log.Printf("failed to saved data: %v", err)
		} else {
			log.Printf("saved %q", s)
		}*/

		c.Unlock()
		once = sync.Once{} // ready for next writing
	}()
	return nil
}


func (c Conf) Ping(args *Void, reply *string) error {
	*reply = "pong"
	return nil
}

// == Commands
//
/*func init() {
	commands := map[string]netutil.CmdDescriptor{
		"GET":  netutil.Descriptor(cmdGet, reflect.String, reflect.Int),
		"PING": netutil.Descriptor(cmdPing),
	}
	config.Server = netutil.NewConfig(defconf.GREETING, commands)
}

// cmdGet returns the configuration for a program.
func cmdGet(cmd *netutil.Command) (string, error) {
	m, err := db.Get(cmd.Values[0].(string), cmd.Values[1].(int))
	if err != nil {
		return "", err
	}

	return m.String(), nil
}

// cmdPing checks if the server is responding.
func cmdPing(*netutil.Command) (string, error) {
	return "PONG", nil
}
*/
// * * *

func printUsage() {
	fmt.Fprintf(os.Stderr, `System configuration server

Usage: piconfd [-v] -tcp [-h -p] -unix [-s] [-wui -http]

`)
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	// == Command-line arguments
	var (
		fVerbose = flag.Bool("v", false, "Verbose")
		fDevel   = flag.Bool("devel", false, "Restart server. Use it only during development!")

		fUseTCP = flag.Bool("tcp", false, "TCP server")
		fHost   = flag.String("h", defconf.HOST, "TCP Host")
		fPort   = flag.Uint("p", defconf.PORT, "TCP Port")

		fUseUnix = flag.Bool("unix", false, "Unix socket server")
		fSocket  = flag.String("s", defconf.SOCKET_FILE, "Unix socket file")

		//fUseWUI = flag.Bool("wui", false, "Web interface")
		//fHTTP   = flag.Uint("http", defconf.HTTP_PORT, "Web port")
	)

	flag.Usage = printUsage
	flag.Parse()

	if len(os.Args) == 1 || (!*fUseUnix && !*fUseTCP) {
		printUsage()
	}
	// ==

	start.Verbose, watch.Verbose = *fVerbose, *fVerbose
	defer start.Exit()

	rpc.Register(db)

	if *fUseUnix {
		listen, err := net.Listen("unix", *fSocket)
		if err != nil {
			log.Fatal("listen error:", err)
			return
		}

		go rpc.Accept(listen)
		//go rpc.ServeConn(listen)

		/*unixServer, err := netutil.NewUnixServer(config.Server, *fSocket, nil)
		if err != nil {
			unixServer.Log.Print(err)
			return
		}
		defer unixServer.Close()

		go unixServer.Serve()*/
	}
	if *fUseTCP {
		listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *fHost, *fPort))
		if err != nil {
			log.Fatal("listen error:", err)
			return
		}

		go rpc.Accept(listen)

		/*tcpServer, err := netutil.NewTCPServer(config.Server, *fHost, *fPort, nil)
		if err != nil {
			tcpServer.Log.Print(err)
			return
		}
		defer tcpServer.Close()

		go tcpServer.Serve()*/
	}
/*	if *fUseWUI {
		httpServer, err := netutil.NewHTTPserver(*fHTTP, nil)
		if err != nil {
			httpServer.Log.Print(err)
			return
		}
		defer httpServer.Close()

		// Serve static files
//		if err = httpServer.FileServer("/static/", wui.StaticDir); err != nil {
//			httpServer.Log.Print(err)
//			return
//		}

		// Handlers
		httpServer.SetHandlers(map[string]netutil.HandleFunc{
			//"/": wui.RootHandler,
			//"/get/": wui.GetHandler,
		})

		go httpServer.Serve()
	}
*/
	// Detect packages updated during development so the server can be restarted.
	if *fDevel {
		// Run "start3 -up piconfd [options]".
		watcher, err := watch.Start("github.com/kless/piconf/piconfd", nil,
			"github.com/kless/netutil",
			"github.com/kless/piconf",
		)

		if err != nil {
			watcher.Log.Print(err)
			return
		}
		defer watcher.Close()
	}

	start.Wait(nil)
}
