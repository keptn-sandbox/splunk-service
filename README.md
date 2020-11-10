# README

This is a Keptn Splunk Service  written in GoLang. 

 

# splunk-service



### Deploy in your Kubernetes cluster

To deploy the current version of the *splunk-service* in your Keptn Kubernetes cluster, apply the [`deploy/service.yaml`](deploy/service.yaml) file:

```console
kubectl apply -f deploy/service.yaml
```

This should install the `splunk-service` together with a Keptn `distributor` into the `keptn` namespace, which you can verify using

```console
kubectl -n keptn get deployment splunk-service -o wide
kubectl -n keptn get pods -l run=splunk-service
```


### Uninstall

To delete a deployed *splunk-service*, use the file `deploy/*.yaml` files from this repository and delete the Kubernetes resources:

```console
kubectl delete -f deploy/service.yaml
```



 
## License

Please find more information in the [LICENSE](LICENSE) file.
