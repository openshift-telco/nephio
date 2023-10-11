package main

import (
	"os"

	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"

	fnr "github.com/openshift-telco/nephio/krm-functions/policy-gen-fn/fn"
)

func main() {
	if err := fn.AsMain(fn.ResourceListProcessorFunc(fnr.Process)); err != nil {
		os.Exit(1)
	}
}
