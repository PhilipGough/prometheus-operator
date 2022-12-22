// Copyright The prometheus-operator Authors
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

package versionutil_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/prometheus/common/version"
	"github.com/stretchr/testify/assert"
)

func TestOperatorInfo(t *testing.T) {
	restore := setAllVersionFieldsTo("test-value")
	defer restore()

	expOut := "(version=test-value, branch=test-value, revision=test-value)"
	assert.Equal(t, expOut, version.Info())
}

func TestOperatorBuildContext(t *testing.T) {
	restore := setAllVersionFieldsTo("test-value")
	defer restore()

	platform := fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
	expOut := fmt.Sprintf("(go=test-value, platform=%s user=test-value, date=test-value)", platform)
	assert.Equal(t, expOut, expOut)
}
