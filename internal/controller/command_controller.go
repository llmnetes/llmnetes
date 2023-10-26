/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logger "sigs.k8s.io/controller-runtime/pkg/log"

	batchv1alpha1 "github.com/yolo-operator/yolo-operator/api/v1alpha1"
	"github.com/yolo-operator/yolo-operator/pkg/condition"
	"github.com/yolo-operator/yolo-operator/pkg/model"
)

// CommandReconciler reconciles a Command object
type CommandReconciler struct {
	Model model.K8SLLM
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=batch.yolo.ahilaly.dev,resources=commands,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=batch.yolo.ahilaly.dev,resources=commands/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=batch.yolo.ahilaly.dev,resources=commands/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Command object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.15.0/pkg/reconcile
func (r *CommandReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logger.FromContext(ctx)
	log.Info("Reconciling", "resource", req.NamespacedName)

	// First, we fetch the Command object.
	var command batchv1alpha1.Command
	err := r.Get(ctx, req.NamespacedName, &command)
	if err != nil {
		log.Error(err, "unable to fetch Command object")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Second, we check if the object is being deleted, and if so, we skip.
	// TODO(a-hilaly): maybe we should leverage finalizers instead too.
	if !command.ObjectMeta.DeletionTimestamp.IsZero() {
		log.Info("Command is deleted")
		return ctrl.Result{}, nil
	}

	// Third, we check if the command has already been processed.
	if condition.HaveSuccessfulCondition(command.Status.Conditions) {
		log.Info("Command already processed")
		return ctrl.Result{}, nil
	}

	// We process the command
	log.Info("Processing", "command", command.Spec.Input)
	output, err := r.Model.RunQuery(command.Spec.Input)
	if err != nil {
		// instead of returning an error, we update the status of the command
		// and let the controller decide what to do with it.
		log.Error(err, "unable to run query")
		command.Status.Conditions = append(command.Status.Conditions, condition.NewFailedCondition(err.Error(), "Unable to run query"))
	} else {
		// Query processed successfully, we can set the output and the condition
		command.Status.Output = output
		command.Status.Conditions = append(command.Status.Conditions, condition.NewSuccessfulCondition("Command processed successfully"))
	}

	log.Info("Processed", "command", command.Spec.Input, "output", output)
	err = r.Status().Update(ctx, &command)
	if err != nil {
		// Temporary error, let's see how this goes.
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CommandReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&batchv1alpha1.Command{}).
		Complete(r)
}
