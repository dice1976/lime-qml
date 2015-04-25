// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package main

import (
	"runtime"

	"gopkg.in/qml.v1"

	"github.com/limetext/gopy/lib"
	"github.com/limetext/lime-backend/lib"
	_ "github.com/limetext/lime-backend/lib/commands"
	"github.com/limetext/lime-backend/lib/log"
	"github.com/limetext/lime-backend/lib/textmate"
)

const (
	console_height  = 20
	render_chan_len = 2
)

var (
	t *qmlfrontend

	scheme *textmate.Theme
	blink  bool
)

func main() {
	// Need to lock the OS thread as OSX GUI requires GUI stuff to run in the main thread
	runtime.LockOSThread()

	log.AddFilter("file", log.DEBUG, log.NewConsoleLogWriter())
	defer func() {
		py.NewLock()
		py.Finalize()
	}()

	t = &qmlfrontend{windows: make(map[*backend.Window]*frontendWindow)}
	go t.qmlBatchLoop()
	qml.Run(t.loop)
}
