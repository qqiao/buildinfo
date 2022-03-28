// Copyright 2017 Qian Qiao
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package buildinfo_test

import (
	"reflect"
	"testing"

	"github.com/qqiao/buildinfo"
)

var expected = &buildinfo.BuildInfo{
	BuildTime: 1648396360413459000,
	Revision:  "e594b8e",
}

func TestLoad(t *testing.T) {
	if got, err := buildinfo.Load("./test_data.json"); err != nil {
		t.Errorf("Error loading test data: %v", err)
	} else {
		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected: %v\nGot: %v", expected, got)
		}
	}
}

func TestLoadAsync(t *testing.T) {
	gotCh, errCh := buildinfo.LoadAsync("./test_data.json")

	select {
	case err := <-errCh:
		t.Errorf("Error loading test data: %v", err)
	case got := <-gotCh:
		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected: %v\nGot: %v", expected, got)
		}
	}
}
