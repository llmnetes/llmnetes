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
	"github.com/yolo-operator/yolo-operator/pkg/k8s"
	"github.com/yolo-operator/yolo-operator/pkg/model"
	"github.com/yolo-operator/yolo-operator/pkg/parser"
)

// CommandExecReconciler reconciles a CommandExec object
type CommandExecReconciler struct {
	Model       model.K8SLLM
	ShellAccess *k8s.ShellAccess
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=batch.yolo.ahilaly.dev,resources=commandexecs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=batch.yolo.ahilaly.dev,resources=commandexecs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=batch.yolo.ahilaly.dev,resources=commandexecs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CommandExec object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.15.0/pkg/reconcile
func (r *CommandExecReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logger.FromContext(ctx)
	log.Info("Reconciling", "resource", req.NamespacedName)

	// First, we fetch the Command object.
	var command batchv1alpha1.CommandExec
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

	executeOutput := false
	// We process the command
	log.Info("Processing step", "commandexec", command.Spec.Input)
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
		executeOutput = true
	}

	log.Info("Processing comeplete", "command", command.Spec.Input, "output", output)
	if !executeOutput {
		return ctrl.Result{}, nil
	}

	resp, err := parser.ParseGPT3Response(output)
	if err != nil {
		command.Status.Conditions = append(command.Status.Conditions, condition.NewFailedCondition(err.Error(), "Unable to parse query response"))
	} else {
		resp.Sanitize()
		workDir, cleanup, err := r.ShellAccess.PrepareFiles(map[string]string{
			resp.FileName: resp.YamlFile,
		})
		if err != nil {
			command.Status.Conditions = append(command.Status.Conditions, condition.NewFailedCondition(err.Error(), "Unable to prepare files in filesystem"))
		} else {
			defer cleanup()
			_, err = r.ShellAccess.RunCommand(workDir, resp.CommandToRun)
			if err != nil {
				command.Status.Conditions = append(command.Status.Conditions, condition.NewFailedCondition(err.Error(), "Unable to run command"))
			} else {
				command.Status.Conditions = append(command.Status.Conditions, condition.NewSuccessfulCondition("Command executed successfully"))
			}
		}
	}

	err = r.Status().Update(ctx, &command)
	if err != nil {
		// Temporary error, let's see how this goes.
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CommandExecReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&batchv1alpha1.CommandExec{}).
		Complete(r)
}
