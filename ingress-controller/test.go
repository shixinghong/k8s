package main

import (
	"os"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"

	"myit.fun/kubernetes/ingress-controller/kubernetes"

	v1 "k8s.io/api/networking/v1"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func main() {

	logf.SetLogger(zap.New())

	var log = logf.Log.WithName("builder-examples")

	config := kubernetes.InitK8S()

	mgr, err := manager.New(config, manager.Options{})
	if err != nil {
		log.Error(err, "could not create manager")
		os.Exit(1)
	}

	if err = builder.ControllerManagedBy(mgr).For(&v1.Ingress{}).Complete(kubernetes.NewOldController()); err != nil {
		log.Error(err, "could not create controller")
		os.Exit(1)
	}

	if err = mgr.Start(signals.SetupSignalHandler()); err != nil {
		log.Error(err, "could not start manager")
		os.Exit(1)
	}
	
}
