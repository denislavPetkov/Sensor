# Sensor

## Overview

CLI application which you can use to get information about your system.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine.

### Prerequisites

- [Git](https://git-scm.com/downloads)
- [Golang](https://golang.org/dl/)

### Downloading

Run `git clone https://github.com/denislavPetkov/sensor.git` in desired project directory

## Building the application

From `sensor/cmd/sensor` directory run `go build -o sensor`, where **sensor** is desired app binary name

### Running the binary

You can run the binary with `sensor --command` in binary directive, where **sensor** is the name of the binary

Example: `sensor -sensor_group CPU_TEMP -delta_duration 1 -total_duration 5 -format JSON -web_hook_url http://localhost:8000/measurement`

## Deploying

### Prerequisites

- [docker](https://docs.docker.com/get-docker/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)

### Docker

***When using the commands listed below you have to be in the project/sensor directory***

`docker build -f ./Dockerfile -t docker.io/sensor/cli:latest .` builds the image

`docker push docker.io/sensor/cli:latest` pushes the image to the repository

### Kubernetes

***When using the commands listed below you have to be in the project/sensor/deployments/kubernetes directory***

`kubectl apply -f ./sensor.yaml` deploys the sensor client with the specified image

## Testing the sensor cli

To test the application run `go test ./...` from `project/sensor` directory

## Available commands

### help

`-help`
*lists the commands and their usage

### get_available_sensors

`-get_available_sensors`
*returns list of sensor names for which exists an implementation

### delta_duration

`-delta_duration int`
*specifies the duration between two sensor measurements (default 5)

### total_duration

`-total_duration int`
*specifies the total duration after which the program is terminated (default 15)

### format

`-format string`
*specifies the output format (JSON or YAML) (default "JSON")

### sensor_group

`-sensor_group value`
*specifies the sensor group (CPU_TEMP or CPU_USAGE or MEMORY_USAGE)

#### Each sensor group contains different sensors

1. CPU_TEMP
   - cpu temperature sensor - in Celsius
2. CPU_USAGE
   - cpu usage percentage sensor - in %
   - cpu cores count sensor
   - cpu frequency sensor - in MHz
3. MEMORY_USAGE
   - total memory sensor - in GigaBytes
   - available memory sensor - in Bytes
   - used memory sensor - in Bytes
   - used memory percentage sensor - in %

multiple sensor_group flags support:
`-sensor_group value1 -sensor_group value2`

#### Example output from different sensor groups

##### CPU_TEMP

```
{"measuredAt":"2021-04-23T15:34:17+03:00","value":"45","sensorId":"1","deviceId":"1"}
```

##### CPU_USAGE

```
{"measuredAt":"2021-04-23T15:39:24+03:00","value":"6","sensorId":"3","deviceId":"1"}
{"measuredAt":"2021-04-23T15:39:24+03:00","value":"2600","sensorId":"4","deviceId":"1"}
{"measuredAt":"2021-04-23T15:39:24+03:00","value":"9.806264955485817","sensorId":"2","deviceId":"1"}
```

##### MEMORY_USAGE

```
{"measuredAt":"2021-04-23T15:42:06+03:00","value":"5723918336","sensorId":"6","deviceId":"1"}
{"measuredAt":"2021-04-23T15:42:06+03:00","value":"16","sensorId":"5","deviceId":"1"}
{"measuredAt":"2021-04-23T15:42:06+03:00","value":"66.68241024017334","sensorId":"8","deviceId":"1"}
{"measuredAt":"2021-04-23T15:42:06+03:00","value":"11455950848","sensorId":"7","deviceId":"1"}
```

### web_hook_url

`-web_hook_url string`
*specifies an url to be called

## Example with all commands

`sensor -sensor_group CPU_TEMP -delta_duration 1 -total_duration 5 -format JSON -web_hook_url http://localhost:8000/measurement`

# Operator

### Prerequisites

- [operator-sdk](https://sdk.operatorframework.io/docs/installation/)
- [git](https://git-scm.com/downloads)
- [go](https://golang.org/dl/) version 1.15
- [docker](https://docs.docker.com/get-docker/) version 17.03+
- [kubectl](https://kubernetes.io/docs/tasks/tools/) and access to a Kubernetes cluster of a [compatible version](https://sdk.operatorframework.io/docs/overview/#kubernetes-version-compatibility)

## Generating an operator

### Creating a new project

`operator-sdk init --domain sensor.cli --repo github.com/denislavPetkov/sensor/operator`

`--domain` flag sets name suffix for all API groups, the example above will result in `<group>.sensor.cli` API groups

`operator-sdk init` generates a `go.mod` file to be used with Go modules. The `--repo=<path>` flag is required when creating a project outside of `$GOPATH/src`, as scaffolded files require a valid module path. Ensure you activate module support by running `export GO111MODULE=on` before using the SDK.

### Creating a new API and Controller

Create a new Custom Resource Definition (CRD) API with version `v1` and Kind `SensorClient`. When prompted, enter yes y for creating both the resource and controller.

`operator-sdk create api --version v1 --kind SensorClient --resource --controller`

This will scaffold the SensorClient resource API at `api/v1/sensorclient_types.go` and the controller at `controllers/sensorclient_controller.go`.

## Configuring the operator

You can configure the operator resources' names, namespace, etc. in `operator/config/default/kustomization.yaml`

By default the operator watches all namespaces for resources, you can set the namespace the operator watches with `export SENSOR_OPERATOR_NAMESPACE=namespace` where namespace is your desired namespace

You have to set image tag base either by changing the Makefile or by env variable with `export IMAGE_TAG_BASE=tag_base` where tag_base is your desired registry, namespace, and partial name for all your image tags

By default the version is 0.0.1 and you can change it with `export VERSION=wanted_version` where wanted_version is the version you want

By default the docker image is created from the `IMAGE_TAG_BASE` and `VERSION` variables, if you need to change that you can do it in the Makefile

Example: if `IMAGE_TAG_BASE=example.com/my-operator` and `VERSION=0.0.1`, the final image will be `example.com/my-operator:0.0.1`. This image is used when pushing to docker repository and when deploying the operator

Example 2: `IMAGE_TAG_BASE=docker.io/sensor/operator` and `VERSION=latest` will result in `docker.io/sensor/operator:latest` final image

## Deploying

***When using the commands listed below you have to be in the operator directory***

### Docker

`make docker-build docker-push IMG=docker.io/sensor/operator:latest` builds and pushes the image to the repository

### Deploying the operator to kubernetes

`make deploy IMG=docker.io/sensor/operator:latest` deploys the operator in the specified namespace with the specified image

## Testing the operator

You can test the operator by executing `make test` from operator directory

***The tests use an existing cluster(current kubeconfig), so you need to provide your own namespace for the deployment of the sensor client and your own sensor client image to be used for that deployment, and own secrets needed for pulling the image***
