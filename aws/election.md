# Elections Things

- [Leader election](https://dev.to/sklarsa/how-to-add-kubernetes-powered-leader-election-to-your-go-apps-57jh)
- [Leader election 2](https://carlosbecker.com/posts/k8s-leader-election/)
- [Leader Election 3](https://taesunny.github.io/kubernetes/kubernetes-controllers-leader-election-with-go-library/)
- [worth to consider](https://github.com/kubernetes/kubernetes/blob/dff82dc0de47299ab66c83c626e08b245ab19037/cmd/kube-controller-manager/app/controllermanager.go#L255)
- [blog](https://dev.to/sklarsa/how-to-add-kubernetes-powered-leader-election-to-your-go-apps-57jh)

```snippets
id = id + "_" + string(uuid.NewUUID())
```

- [sigs leader election](https://github.com/search?q=org%3Akubernetes-sigs%20k8s.io%2Fclient-go%2Ftools%2Fleaderelection%2Fresourcelock&type=code)
- [Code Examples](https://github.com/kubernetes-sigs/descheduler/blob/fa3fb4e954f02e837068f73e7520a9cc02f8bc52/pkg/descheduler/leaderelection.go#L27)
- [fake leader election](https://github.com/kubernetes-sigs/controller-runtime/blob/91e102ffc63c9b415b177eee448316556fa8d315/pkg/leaderelection/fake/leader_election.go)
- [karpenter](https://github.com/kubernetes-sigs/karpenter/blob/c380935b0d9a14aca8b5925c6b276e88fb93f21f/pkg/operator/operator.go#L42)
- [kubemark manager](https://github.com/kubernetes-sigs/cluster-api-provider-kubemark/blob/main/main.go)
- [how to find out that in cluster running](https://github.com/kubernetes-sigs/controller-runtime/blob/91e102ffc63c9b415b177eee448316556fa8d315/pkg/leaderelection/leader_election.go)
- [cert manager](https://github.com/cert-manager/cert-manager/blob/027b692921688b0afd57916cc771878f7ed8532a/internal/apis/config/shared/types_leaderelection.go)
- [cert manager v2](https://github.com/cert-manager/cert-manager/blob/027b692921688b0afd57916cc771878f7ed8532a/pkg/webhook/server/server.go#L140)
- [deschedule](https://github.com/kubernetes-sigs/descheduler/blob/fa3fb4e954f02e837068f73e7520a9cc02f8bc52/pkg/descheduler/leaderelection.go#L41)
- [required rbac](https://carlosbecker.com/posts/k8s-leader-election/)
- [go eleaction with kind](https://github.com/olamilekan000/leader-election)
- [descheduler](https://github.com/kubernetes-sigs/descheduler/blob/fa3fb4e954f02e837068f73e7520a9cc02f8bc52/pkg/descheduler/leaderelection.go#L28)
