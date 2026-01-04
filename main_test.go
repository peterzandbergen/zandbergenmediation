// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"testing"
)

func TestMain(t *testing.T) {
	getenv := func(key string) string {
		var val string
		switch key {
		case "PORT":
			val = "8080"
		case "URL_FROM":
			val = "localhost"
		case "URL_TO":
			val = "127.0.0.1"
		}
		return val
	}
	run(context.Background(), getenv)
}
