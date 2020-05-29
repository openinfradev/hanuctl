# Set up Ingress on TACO with the NGINX Ingress Controller

## Prerequisite
A NodePort is used in this manual, so you should allow NodePorts(30000-32767) to
your security group in TACO OpenStack.

## Deploy the Ingress Controller
Nginx Ingress Controller provides some yaml for various enviroment. Because TACO
does not support Load Balancer yet, you should use yaml for baremetal. 
(See [Install guide](https://kubernetes.github.io/ingress-nginx/deploy/#bare-metal) for more information)

```sh
$ kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-0.32.0/deploy/static/provider/baremetal/deploy.yaml
namespace/ingress-nginx created
serviceaccount/ingress-nginx created
configmap/ingress-nginx-controller created
clusterrole.rbac.authorization.k8s.io/ingress-nginx created
clusterrolebinding.rbac.authorization.k8s.io/ingress-nginx created
role.rbac.authorization.k8s.io/ingress-nginx created
rolebinding.rbac.authorization.k8s.io/ingress-nginx created
service/ingress-nginx-controller-admission created
service/ingress-nginx-controller created
deployment.apps/ingress-nginx-controller created
validatingwebhookconfiguration.admissionregistration.k8s.io/ingress-nginx-admission created
clusterrole.rbac.authorization.k8s.io/ingress-nginx-admission created
clusterrolebinding.rbac.authorization.k8s.io/ingress-nginx-admission created
job.batch/ingress-nginx-admission-create created
job.batch/ingress-nginx-admission-patch created
role.rbac.authorization.k8s.io/ingress-nginx-admission created
rolebinding.rbac.authorization.k8s.io/ingress-nginx-admission created
serviceaccount/ingress-nginx-admission created
 
$ kubectl get po -n ingress-nginx
NAME                                       READY   STATUS      RESTARTS   AGE
ingress-nginx-admission-create-sggkx       0/1     Completed   0          71m
ingress-nginx-admission-patch-csf25        0/1     Completed   0          71m
ingress-nginx-controller-f8d756996-57ltx   1/1     Running     0          71m
 
$ kubectl get svc -n ingress-nginx
 
 
NAME                                 TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                      AGE
ingress-nginx-controller             NodePort    10.97.197.215   <none>        80:32402/TCP,443:31299/TCP   140m
ingress-nginx-controller-admission   ClusterIP   10.108.181.45   <none>        443/TCP                      140m

```

## Deploy the sample app

`hello-app` is sample application that execute a simple http server.
You can access this server through URL(http://<VM_PUBLIC_IP>:<NODEPORT>)
(Below example, URL is http://<VM_PUBLIC_IP>:31670)

```sh
$ kubectl create deployment web --image=gcr.io/google-samples/hello-app:1.0
deployment.apps/web created

$ kubectl expose deployment web --type=NodePort --port=8080

$ kubectl get svc
NAME         TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
kubernetes   ClusterIP   10.96.0.1        <none>        443/TCP          3d1h
web          NodePort    10.99.88.2       <none>        8080:31670/TCP   162m

$ curl http://192.168.97.62:31670
Hello, world!
Version: 1.0.0
Hostname: web-557f59c6cf-7mnmr
```

## Create an Ingress resource

For ingress resource, you can access your application via DNS. For example, 
`http://<YOUR_DOMAIN>`. In this manual, we use [nip.io](https://nip.io) for use 
DNS.

```sh
$ vi ingress.yaml
 apiVersion: networking.k8s.io/v1beta1
 kind: Ingress
 metadata:
   name: example-ingress
   annotations:
     nginx.ingress.kubernetes.io/rewrite-target: /$1
 spec:
   rules:
   - host: hello-world-192-168-97-62.nip.io
     http:
       paths:
       - path: /
         backend:
           serviceName: web
           servicePort: 8080

$ kubectl apply -f ingress.yaml
NAME              HOSTS                        ADDRESS   PORTS   AGE
example-ingress   nginx-192-168-97-62.nip.io             80      11s

```

## Test

Now, you can access this server via DNS.
```sh 
$ curl http://hello-world-192-168-97-62.nip.io:32402
Hello, world!
Version: 1.0.0
Hostname: web-557f59c6cf-7mnmr
```
