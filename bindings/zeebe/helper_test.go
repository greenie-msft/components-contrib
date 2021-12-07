/*
Copyright 2021 The Dapr Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package zeebe

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVariableStringToArrayRemovesSpaces(t *testing.T) {
	vars := VariableStringToArray("  a,   b,  c  ")
	require.Equal(t, 3, len(vars))
	assert.Equal(t, "a", vars[0])
	assert.Equal(t, "b", vars[1])
	assert.Equal(t, "c", vars[2])
}
