# CSI Cinder driver for tacoctl

CSI Cinder driver can create a volume using OpenStack Cinder and attach to 
Kubernetes Pod. So If you want to mount a volume in Pod, you should be try this.

## Deploy

You can use the manifests under `addons/cinder-csi-plugin`

```sh
$ kubectl apply -f addons/cinder-csi-plugin
serviceaccount/csi-cinder-controller-sa created
clusterrole.rbac.authorization.k8s.io/csi-attacher-role created
clusterrolebinding.rbac.authorization.k8s.io/csi-attacher-binding created
clusterrole.rbac.authorization.k8s.io/csi-provisioner-role created
clusterrolebinding.rbac.authorization.k8s.io/csi-provisioner-binding created
clusterrole.rbac.authorization.k8s.io/csi-snapshotter-role created
clusterrolebinding.rbac.authorization.k8s.io/csi-snapshotter-binding created
clusterrole.rbac.authorization.k8s.io/csi-resizer-role created
clusterrolebinding.rbac.authorization.k8s.io/csi-resizer-binding created
role.rbac.authorization.k8s.io/external-resizer-cfg created
rolebinding.rbac.authorization.k8s.io/csi-resizer-role-cfg created
service/csi-cinder-controller-service created
statefulset.apps/csi-cinder-controllerplugin created
serviceaccount/csi-cinder-node-sa created
clusterrole.rbac.authorization.k8s.io/csi-nodeplugin-role created
clusterrolebinding.rbac.authorization.k8s.io/csi-nodeplugin-binding created
daemonset.apps/csi-cinder-nodeplugin created
csidriver.storage.k8s.io/cinder.csi.openstack.org created
secret/cloud-config created

kubectl get po -n kube-system|grep csi
csi-cinder-controllerplugin-0                                 5/5     Running   0          5m50s
csi-cinder-nodeplugin-sr58k                                   2/2     Running   0          18m
csi-cinder-nodeplugin-xmkn7                                   2/2     Running   0          18m
```

`csi-cinder-nodeplugin` is daemonset so it runs as many as the number of nodes.

## Usage

Try following [document](https://github.com/kubernetes/cloud-provider-openstack/blob/master/docs/using-cinder-csi-plugin.md#example-nginx-application-usage), you can create StorageClass, PersistentVolumeClaim and Pod to consume it.
