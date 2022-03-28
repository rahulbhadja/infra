# Kubernetes (Helm)

Using Infra CLI: 

1. Generate the helm install command via

```
infra destinations add kubernetes example-cluster-name
```

2. Run the output Helm command on the Kubernetes cluster to be added.

Example output: 
```
helm upgrade --install infra-engine infrahq/infra --set engine.config.accessKey=2pVqDSdkTF.oSCEe6czoBWdgc6wRz0ywK8y --set engine.config.name=kubernetes.example-cluster --set engine.config.server=https://infra.acme.com
```