### Tests

The SiteConfig binary is necessary for local testing.

The extra-manifest folder is necessary for siteconfig-gen to generate the appropriate manifest. This folder is in the repo for test purposes. The container packaging this KRM function is fetching the extra-manifest from the published image.

```bash
docker cp quay.io/openshift-kni/ztp-site-generator:/home/ztp/extra-manifest extra-manifest
docker cp quay.io/openshift-kni/ztp-site-generator:/kustomize/plugin/ran.openshift.io/v1/siteconfig/SiteConfig SiteConfig

```