# kubectl-switch
kubectl-switch is a zero-dependeny kubectl plugin to provide a fast and flexible way to change the namespace or the context of the kubeconfig file.

# How to use it
The binary provides two simple commands and a help function:
The binary provides two simple commands and a help function:
	
the ns command provides a way to switch or list all namespaces, so if you want to list all namespaces you specify no more parameter just the ns argument.
When you want to switch you add the namespaces right after the ns command

```bash
# with kubectl plugin
kubectl switch ns # get a list of all namespaces
kubectl switch ns kube-system # switch to the kube-system namespace for the actual context

# standalone
kubectl-switch ns # get a list of all namespaces
kubectl-switch ns kube-system # switch to the kube-system namespace for the actual context
```

the ctx command provides a way to switch or list all contexts, so if you want to list all contexts you specify no more parameter just the ctx argument.
When you want to switch you add the contexts right after the ctx command

```bash
# with kubectl plugin
kubectl switch ctx # get a list of all contexts
kubectl switch ctx kubernetes-admin # switch to the kubernetes-admin context

# standalone
kubectl-switch ctx # get a list of all contexts
kubectl-switch ctx kubernetes-admin # switch to the kubernetes-admin context
```

## prebuild binaries
You can use prebuild binaries for every platform on the []()

## build from source
You can build the program from source with following commands:
```bash
# Clone the repository
git clone https://github.com/Dominik-Robert/kubectl-switch.git

# compile the code
cd kubectl-switch; go build -o kubectl-switch . 
```

# compatiblity
The program runs on nearly every platform and kubernetes version without any dependency.

# TODOs
Feel free to add any todo via issue or PR 