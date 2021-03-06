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

package buildinfo // import "github.com/qqiao/buildinfo"

import (
	"encoding/json"
	"io/ioutil"
)

// BuildInfo is the data element representing information for a build.
type BuildInfo struct {
	BuildTime int64  `json:"buildTime"`
	Revision  string `json:"revision"`
}

// Load loads the build info from the given path.
func Load(path string) (BuildInfo, error) {
	var buildInfo BuildInfo
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return buildInfo, err
	}

	return buildInfo, json.Unmarshal(file, &buildInfo)
}
