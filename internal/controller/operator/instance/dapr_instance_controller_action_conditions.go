package instance

import (
	"context"
	"fmt"

	"github.com/dapr/kubernetes-operator/pkg/conditions"
	"github.com/dapr/kubernetes-operator/pkg/controller/gc"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"

	"github.com/dapr/kubernetes-operator/pkg/controller/client"
	"github.com/dapr/kubernetes-operator/pkg/helm"
	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/builder"
)

func NewConditionsAction(l logr.Logger) Action {
	return &ConditionsAction{
		engine:        helm.NewEngine(),
		l:             l.WithName("action").WithName("conditions"),
		subscriptions: make(map[string]struct{}),
		gc:            gc.New(),
	}
}

type ConditionsAction struct {
	engine        *helm.Engine
	gc            *gc.GC
	l             logr.Logger
	subscriptions map[string]struct{}
}

func (a *ConditionsAction) Configure(_ context.Context, _ *client.Client, b *builder.Builder) (*builder.Builder, error) {
	return b, nil
}

func (a *ConditionsAction) Run(ctx context.Context, rc *ReconciliationRequest) error {
	crs, err := CurrentReleaseSelector(rc)
	if err != nil {
		return fmt.Errorf("cannot compute current release selector: %w", err)
	}

	deployments, err := rc.Client.AppsV1().Deployments(rc.Resource.Namespace).List(ctx, metav1.ListOptions{
		LabelSelector: crs.String(),
	})

	if err != nil {
		return fmt.Errorf("cannot list deployments: %w", err)
	}

	ready := 0

	for i := range deployments.Items {
		if conditions.ConditionStatus(deployments.Items[i], appsv1.DeploymentAvailable) == corev1.ConditionTrue {
			ready++
		}
	}

	var readyCondition metav1.Condition

	if len(deployments.Items) > 0 {
		if ready == len(deployments.Items) {
			readyCondition = metav1.Condition{
				Type:               conditions.TypeReady,
				Status:             metav1.ConditionTrue,
				Reason:             "Ready",
				Message:            fmt.Sprintf("%d/%d deployments ready", ready, len(deployments.Items)),
				ObservedGeneration: rc.Resource.Generation,
			}
		} else {
			readyCondition = metav1.Condition{
				Type:               conditions.TypeReady,
				Status:             metav1.ConditionFalse,
				Reason:             "InProgress",
				Message:            fmt.Sprintf("%d/%d deployments ready", ready, len(deployments.Items)),
				ObservedGeneration: rc.Resource.Generation,
			}
		}
	} else {
		readyCondition = metav1.Condition{
			Type:               conditions.TypeReady,
			Status:             metav1.ConditionFalse,
			Reason:             "InProgress",
			Message:            "no deployments",
			ObservedGeneration: rc.Resource.Generation,
		}
	}

	meta.SetStatusCondition(&rc.Resource.Status.Conditions, readyCondition)

	return nil
}

func (a *ConditionsAction) Cleanup(_ context.Context, _ *ReconciliationRequest) error {
	return nil
}
