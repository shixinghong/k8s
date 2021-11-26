package kubernetes

import (
	"context"
	v1 "k8s.io/api/networking/v1"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	annotation = "kubernetes.io/ingress.class"
)

type OldController struct {
	client.Client
}

func NewOldController() *OldController {
	return &OldController{}
}

func (a *OldController) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	ing := &v1.Ingress{}
	err := a.Get(ctx, req.NamespacedName, ing)
	if err != nil {
		return reconcile.Result{}, err
	}

	if an, ok := ing.Annotations[annotation]; ok && an == "old" {
		klog.Info(ing)
	}

	return reconcile.Result{}, nil
}

func (a *OldController) InjectClient(c client.Client) error {
	a.Client = c
	return nil
}
