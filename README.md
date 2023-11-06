# LLMNETES

A Kubernetes controller that allows users to operate their kubernetes using natural language.

## Description

llmnetes is a Kubernetes controller that allows users to operate their kubernetes
cluster using a natural language interface. It is backed by LLM (Language Learning Model)
which translates natural language commands into kubernetes API calls. It is capable of
understanding commands like `create a deployment with 3 replicas` or `delete all pods in the default namespace`.
Or even triggering chaos experiments like `kill a pod in the default namespace`. and much more.

llmenetes supports multiple LLM backends. Currently, it supports
[OpenAI's GPT-3](https://openai.com/blog/openai-api/) and  a local model that is trained on the
[TODO](TODO) dataset. It is also possible to add your own LLM backend by implementing
the [LLM interface]().

## Support table

| LLM backend | Supported | Notes |
| ----------- | --------- | ----- |
| OpenAI GPT-X | Yes | |
| llama local model | WIP | |

## Status

It is currently in development and is not ready for production use.

## Installation

### Prerequisites

- A Kubernetes cluster with a version >= 1.16.0
- [helm](https://helm.sh/docs/intro/install/) installed and configured to access your cluster
- An OPENAI API key

Modify the `deploy/helm-chart/llmnetes/values.yaml` file to add your OPENAI API key
and other configuration options. Then, install the helm chart:

```bash
$ helm install llmnetes deploy/helm-chart/llmnetes
```


## Examples

#### Create a deployment with 3 replicas

To deploy new pods using llmnetes, you can deploy the following manifest:

```yaml
apiVersion: batch.yolo.ahilaly.dev/v1alpha1
kind: Command
metadata:
  name: my-command
spec:
  input: Create 3 nginx pods that will serve traffic on port 80.
```

This will create 3 nginx pods that will serve traffic on port 80.

#### Chaos experiments

llmnetes can also be used to trigger chaos experiments. For example, to kill a pod in the default namespace, you can deploy the following manifest:

```yaml
apiVersion: batch.yolo.ahilaly.dev/v1alpha1
kind: ChaosSimulation
metadata:
  name: chaos-simulation-cr
spec:
  level: 10 # 1 being the lowest and 10 the highest
  command: break my cluster networking layer (or at least try to)
```

## License

Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

