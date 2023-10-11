# policy-gen-fn

## Overview

<!--mdtogo:Short-->

This function is the PolicyGen generator and is based of [openshift-kni/cnf-features-deploy](https://github.com/openshift-kni/cnf-features-deploy) which is productized in Red Hat Advanced Cluster Management product.

<!--mdtogo-->

The policy-generator library makes cluster deployment easier by generating the following CRs based on a PolicyGenTemplate CR instance;

The [PolicyGenTemplate](https://github.com/openshift-kni/cnf-features-deploy/blob/release-4.13/ztp/ran-crd/policy-gen-template-crd.yaml) is a Custom Resource created to facilitate the creation of those CRs and avoid repeating the configuration names.

<!--mdtogo:Long-->

## Usage

To use this function, define a [SiteConfig](https://github.com/openshift-kni/cnf-features-deploy/blob/release-4.13/ztp/ran-crd/policy-gen-template-crd.yaml) resource.

### FunctionConfig

```yaml
apiVersion: ran.openshift.io/v1
kind: PolicyGenTemplate
metadata:
  name: "example-sno"
  namespace: "ztp-site"
spec:
  bindingRules:
    # These policies will correspond to all clusters with this label:
    sites: "example-sno"
  mcp: "master"
  sourceFiles:
    - fileName: SriovNetwork.yaml
      policyName: "config-policy"
      metadata:
        name: "sriov-nw-du-fh"
      spec:
        resourceName: du_fh
        vlan: 140
    - fileName: SriovNetworkNodePolicy.yaml
      policyName: "config-policy"
      metadata:
        name: "sriov-nnp-du-fh"
      spec:
        deviceType: netdevice
        isRdma: true
        nicSelector:
          pfNames: ["ens5f0"]
        numVfs: 8
        priority: 10
        resourceName: du_fh
```

The function config above will generate all the required manifests.

<!--mdtogo-->