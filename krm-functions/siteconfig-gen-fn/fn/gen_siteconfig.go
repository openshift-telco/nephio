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
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
)

func Process(rl *fn.ResourceList) (bool, error) {

	var out strings.Builder
	tmp := rl.FunctionConfig.String()
	err := os.WriteFile("/tmp/siteconfig.yaml", []byte(tmp), 0644)
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("/usr/local/bin/SiteConfig", "/tmp/siteconfig.yaml")
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	for i, cr := range strings.Split(out.String(), "---\n") {
		if i == 0 {
			continue
		}

		obj, err := fn.ParseKubeObject([]byte(cr))
		if err != nil {
			return false, err
		}

		err = rl.UpsertObjectToItems(obj, nil, true)
		if err != nil {
			return false, err
		}

	}

	return true, nil
}
