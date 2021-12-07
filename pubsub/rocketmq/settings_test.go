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

package rocketmq

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSettingsDecode(t *testing.T) { //nolint:paralleltest
	props := map[string]string{
		"accessProto":   "http",
		"accessKey":     "**",
		"secretKey":     "***",
		"endpoint":      "http://test.endpoint",
		"nameServer":    "http://test.nameserver",
		"consumerGroup": "defaultGroup",
		"instanceId":    "defaultNamespace",
		"topics":        "defaultTopic",
	}

	var settings Settings
	err := settings.Decode(props)
	require.NoError(t, err)

	assert.Equal(t, "http", settings.AccessProto)
	assert.Equal(t, "**", settings.AccessKey)
	assert.Equal(t, "***", settings.SecretKey)
	assert.Equal(t, "http://test.endpoint", settings.Endpoint)
	assert.Equal(t, "defaultGroup", settings.ConsumerGroup)
	assert.Equal(t, "defaultNamespace", settings.InstanceID)
}

func TestParseCommonMetadata(t *testing.T) { //nolint:paralleltest
	props := map[string]string{
		"accessProto":   "http",
		"accessKey":     "**",
		"secretKey":     "***",
		"endpoint":      "http://test.endpoint",
		"nameServer":    "http://test.nameserver",
		"consumerGroup": "defaultGroup",
		"instanceId":    "defaultNamespace",
		"topics":        "defaultTopic",
	}

	var settings Settings
	err := settings.Decode(props)
	require.NoError(t, err)

	b := settings.ToRocketMQMetadata()

	assert.Equal(t, "http", b.AccessProto)
	assert.Equal(t, "**", b.AccessKey)
	assert.Equal(t, "***", b.SecretKey)
	assert.Equal(t, "http://test.endpoint", b.Endpoint)
	assert.Equal(t, "defaultGroup", b.ConsumerGroup)
	assert.Equal(t, "defaultNamespace", b.InstanceId)
}
