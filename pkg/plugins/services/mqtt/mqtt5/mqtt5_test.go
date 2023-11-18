// Copyright 2022 Praetorian Security, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mqtt5

import (
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/omarbdrn/fingerprintx/pkg/plugins"
	"github.com/omarbdrn/fingerprintx/pkg/test"
)

func TestMqtt5(t *testing.T) {
	testcases := []test.Testcase{
		{
			Description: "mqtt",
			Port:        1883,
			Protocol:    plugins.TCP,
			Expected: func(res *plugins.Service) bool {
				return res != nil
			},
			RunConfig: dockertest.RunOptions{
				Repository: "efrecon/mosquitto",
			},
		},
	}

	p := &MQTT5Plugin{}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.Description, func(t *testing.T) {
			t.Parallel()
			err := test.RunTest(t, tc, p)
			if err != nil {
				t.Errorf(err.Error())
			}
		})
	}
}
