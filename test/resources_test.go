// Copyright (c) Alex Ellis 2017. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package test

import (
	"testing"

	"github.com/openfaas/faas-swarm/handlers"
)

// Test_ParseMemory exploratory testing to document how to convert
// from Docker limits notation to bytes value.
func Test_ParseMemory(t *testing.T) {
	value := "512 m"

	val, err := handlers.ParseMemory(value)
	if err != nil {
		t.Error(err)
	}

	if val != 1024*1024*512 {
		t.Errorf("want: %d got: %d", 1024, val)
	}
}
