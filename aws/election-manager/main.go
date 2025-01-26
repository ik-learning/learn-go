package main

// https://dev.to/sklarsa/how-to-add-kubernetes-powered-leader-election-to-your-go-apps-57jh
// https://github.com/kubernetes/kubernetes/blob/dff82dc0de47299ab66c83c626e08b245ab19037/cmd/kube-controller-manager/app/controllermanager.go#L255
// k get lease
import (
	"context"
	"fmt"
	"os"
	"time"

	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/client-go/tools/leaderelection"
	rl "k8s.io/client-go/tools/leaderelection/resourcelock"
	ctrl "sigs.k8s.io/controller-runtime"

	k8sruntime "k8s.io/apimachinery/pkg/runtime"
)

var (
	// lockName and lockNamespace need to be shared across all running instances
	lockName      = fmt.Sprintf("my-lock-election-dev")
	lockNamespace = "default"

	// identity is unique to the individual process. This will not work for anything,
	// outside of a toy example, since processes running in different containers or
	// computers can share the same pid.
	identity = fmt.Sprintf("%d-%s", os.Getpid(), uuid.NewUUID())
)

func main() {
	fmt.Printf(">> IDENTITY: %s\n", identity)

	scheme := k8sruntime.NewScheme()
	fmt.Println("scheme:", scheme)

}
