### Tests

The PolicyGenTemplate binary is necessary for local testing.

The source-crs folder is necessary for policy-gen to generate the appropriate manifest. This folder is in the repo for test purposes. The container packaging this KRM function is fetching the source-crs from the published image.

```bash
docker cp quay.io/openshift-kni/ztp-site-generator:/kustomize/plugin/ran.openshift.io/v1/policygentemplate/source-crs /usr/local/bin/source-crs
docker cp quay.io/openshift-kni/ztp-site-generator:/kustomize/plugin/ran.openshift.io/v1/policygentemplate/PolicyGenTemplate /usr/local/bin/PolicyGenTemplate
```