# siteconfig-gen-fn

## Overview

<!--mdtogo:Short-->

This function is the SiteConfig generator and is based of [openshift-kni/cnf-features-deploy](https://github.com/openshift-kni/cnf-features-deploy) which is productized in Red Hat Advanced Cluster Management product.

<!--mdtogo-->

The siteconfig-generator library makes cluster deployment easier by generating the following CRs based on a SiteConfig CR instance;

    AgentClusterInstall
    ClusterDeployment
    NMStateConfig
    KlusterletAddonConfig
    ManagedCluster
    InfraEnv
    BareMetalHost
    HostFirmwareSettings
    ConfigMap for extra-manifest configurations

The [SiteConfig](https://github.com/openshift-kni/cnf-features-deploy/blob/release-4.13/ztp/ran-crd/site-config-crd.yaml) is a Custom Resource created to facilitate the creation of those CRs and avoid repeating the configuration names.

<!--mdtogo:Long-->

## Usage

To use this function, define a [SiteConfig](https://github.com/openshift-kni/cnf-features-deploy/blob/release-4.13/ztp/ran-crd/site-config-crd.yaml) resource.

### FunctionConfig

```yaml
apiVersion: ran.openshift.io/v1
kind: SiteConfig
metadata:
  name: "site-plan-sno-du-ex"
  namespace: "clusters-sub"
spec:
  baseDomain: "example.com"
  pullSecretRef:
    name: "pullSecret-name"
  clusterImageSetNameRef: "openshift-v4.8.0"
  sshPublicKey: "ssh-rsa "
  sshPrivateKeySecretRef:
    name: "privKeySecret-name"
  clusters:
    - clusterName: "du-sno-ex"
      clusterLabels:
        group-du-sno: ""
        common: true
        sites : "site-plan-sno-du-ex"
      clusterNetwork:
        - cidr: 10.128.0.0/14
          hostPrefix: 23
      machineNetwork:
        - cidr: 10.16.231.0/24
      serviceNetwork:
        - 172.30.0.0/16
      additionalNTPSources:
        - NTP.server1
        - 10.16.231.22
      diskEncryption:
        type: "nbde"
        tang:
          - url: "http://10.0.0.1:7500"
            thumbprint: "1c3wJKh6TQKTghTjWgS4MlIXtGk"
          - url: "http://10.0.0.2:7500"
            thumbprint: "WOjQYkyK7DxY_T5pMncMO5w0f6E"
      nodes:
        - hostName: "node-ex"
          bmcAddress: "idrac-virtualmedia+https://10.16.231.87/redfish/v1/Systems/System.Embedded.1"
          bmcCredentialsName:
            name: "bmcCredentialsSecret-Name"
          bootMACAddress: "00:00:00:01:20:30"
          bootMode: "UEFI"
          rootDeviceHints:
            hctl: '1:2:0:0'
          # userData contain the user data to be passed to the host before it boots
          userData:
            bootKey: value1
          cpuset: "2-19,22-39"
          nodeNetwork:
            interfaces:
              - name: eno1
                macAddress: 00:00:00:01:20:30
              - name: eth0
                macAddress: 02:00:00:80:12:14
              - name: eth1
                macAddress: 02:00:00:80:12:15
            config:
              # Example for the nmstate config. The interface names must match the defined interfaces above.
              interfaces:
                - name: eno1
                  type: ethernet
                  ipv4:
                    enabled: true
                    dhcp: false
                    address:
                      - 10.16.231.3/24
                      - 10.16.231.28/24
                      - 10.16.231.31/24
                  ipv6:
                    enabled: true
                    dhcp: false
                    address:
                      - 2620:52:0:10e7:e42:a1ff:fe8a:601/64
                      - 2620:52:0:10e7:e42:a1ff:fe8a:602/64
                      - 2620:52:0:10e7:e42:a1ff:fe8a:603/64
                - name: bond99
                  type: bond
                  state: up
                  ipv6:
                    address:
                      - ip:2620:52:0:1302::100
                    prefix-length: 64
                    enabled: true
                    link-aggregation:
                      mode: balance-rr
                      options:
                        miimon: '140'
                      slaves:
                        - eth0
                        - eth1
              dns-resolver:
                config:
                  server:
                    - 10.19.42.41
              routes:
                config:
                  - destination: 0.0.0.0/0
                    next-hop-address: 10.16.231.254
                    next-hop-interface: eno1
                    table-id: 254
```

The function config above will generate all the required manifest for OpenShift to create the cluster and import in the management cluster.

<!--mdtogo-->