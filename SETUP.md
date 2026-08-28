# Local setup

Use Bash on macOS or Linux. These steps use a cluster named `kind`.

## Prerequisites

Install [Go 1.27.0+](https://go.dev/doc/install),
[Docker](https://docs.docker.com/get-started/get-docker/),
[Kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation),
[kubectl](https://kubernetes.io/docs/tasks/tools/), and
[Pulumi CLI v3](https://www.pulumi.com/docs/install/).
You also need Git, Bash, and curl. Start Docker before continuing.

## 1. Prepare the repository and cluster

Skip cluster creation if `kind` already exists.

```sh
git clone https://github.com/dgsamper/echo-service.git
cd echo-service
kind create cluster --name kind --wait 30s
kubectl config use-context kind-kind
```

## 2. Test, build, and load the image

```sh
./scripts/ci.sh
```

The script tests and vets the application, validates the Pulumi Go module, builds the Docker image, then checks Kind and loads `echo-service:local`.

## 3. Deploy

Pulumi uses the current Kubernetes context; keep `kind-kind` selected.

```sh
cd infra
pulumi login --local
```

If Pulumi shows “You don't have any stacks yet”, choose **Skip for now**.
The project already exists in this repository; the next command will create the
local stack.

```sh
pulumi stack init dev
pulumi up
cd ..
```

Keep the passphrase Pulumi requests and confirm the deployment preview.
State stays in `~/.pulumi`; no Pulumi Cloud account is needed.
If the stack already exists, use `pulumi stack select dev` instead of `stack init`.

## 4. Try the service

```sh
kubectl rollout status deployment/echo-service --timeout=60s
kubectl port-forward service/echo-service 8080:8080
```

Keep that terminal running. In another terminal:

```sh
curl -fsS 'http://localhost:8080/hello?name=world' --data-binary 'hello'
```

Expect JSON containing `headers`, `params`, `body`, and `path`.

After code changes, rerun `./scripts/ci.sh`, then
`kubectl rollout restart deployment/echo-service` because the image tag stays the
same. Restart port forwarding after the rollout.

## 5. Cleanup

Stop port forwarding with Ctrl+C. From the repository root, remove the resources
before the cluster. Only delete the cluster if it was created for this exercise.

```sh
kubectl config use-context kind-kind
cd infra
pulumi stack select dev
pulumi destroy
pulumi stack rm dev
cd ..
kind delete cluster --name kind
```
