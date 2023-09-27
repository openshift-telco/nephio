/*
Copyright 2023 Nephio.

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

package fn

import (
	"testing"

	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"github.com/stretchr/testify/assert"
)

func TestErrorCases(t *testing.T) {
	cases := map[string]struct {
		input       []byte
		errExpected string
	}{
		"GoTemplateExecuteError": {
			input:       []byte(``),
			errExpected: "",
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			ko, err := fn.ParseKubeObject(tc.input)
			assert.NoError(t, err)

			rl := &fn.ResourceList{FunctionConfig: ko}

			_, err = Process(rl)
			if tc.errExpected != "" {
				assert.EqualError(t, err, tc.errExpected)
			} else {
				assert.NoError(t, err)
			}

			b, _ := rl.ToYAML()
			println(string(b))
		})
	}
}
