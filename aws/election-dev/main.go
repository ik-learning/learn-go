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
	// Get the active kubernetes context
	cfg, err := ctrl.GetConfig()
	if err != nil {
		panic(err.Error())
	}

	// Create a new lock. This will be used to create a Lease resource in the cluster.
	l, err := rl.NewFromKubeconfig(
		rl.LeasesResourceLock,
		lockNamespace,
		lockName,
		rl.ResourceLockConfig{
			Identity: identity,
		},
		cfg,
		time.Second*10,
	)
	if err != nil {
		panic(err)
	}

	run := func(ctx context.Context) {
		fmt.Printf("i am the leader %s\n", identity)
		select {}
	}

	// Create a new leader election configuration with a 15 second lease duration.
	// Visit https://pkg.go.dev/k8s.io/client-go/tools/leaderelection#LeaderElectionConfig
	// for more information on the LeaderElectionConfig struct fields
	el, err := leaderelection.NewLeaderElector(leaderelection.LeaderElectionConfig{
		Lock:          l,
		LeaseDuration: time.Second * 15,
		RenewDeadline: time.Second * 10,
		RetryPeriod:   time.Second * 2,
		Name:          lockName,
		// election checker
		WatchDog: leaderelection.NewLeaderHealthzAdaptor(time.Second * 10),
		//
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: run,
			OnStoppedLeading: func() {
				println("leader election lost")
			},
			OnNewLeader: func(identity string) {
				fmt.Printf("the leader is %s\n", identity)
			},
		},
	})
	if err != nil {
		panic(err)
	}

	// Begin the leader election process. This will block.
	el.Run(context.Background())
}
