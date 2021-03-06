// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package sqlite

import (
	"reflect"
	"testing"
)

func TestSqlite_createBuildService(t *testing.T) {
	// setup types
	want := &Service{
		Create:  CreateBuildTable,
		Indexes: []string{CreateBuildRepoIDIndex, CreateBuildRepoIDNumberIndex, CreateBuildStatusIndex},
	}

	// run test
	got := createBuildService()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("createBuildService is %v, want %v", got, want)
	}
}
