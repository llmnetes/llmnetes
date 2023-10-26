package condition

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ConditionType string

const (
	// ConditionTypeSuccessful indicates that the CR has reached a successful state
	ConditionTypeSuccessful ConditionType = "Successful"

	// ConditionTypeRecoverable indicates that the CR has reached a failed but recoverable
	// state. Meaning the controller will attempt to reconcile it again.
	ConditionTypeRecoverable ConditionType = "Recoverable"

	// ConditionTypeFailed indicates that the CR has reached a failed state. Meaning
	// the controller will not attempt to reconcile it again.
	ConditionTypeFailed ConditionType = "Failed"
)

// Condition contains details for the current condition of this resource.
type Condition struct {
	// Type is the type of the Condition
	Type ConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status corev1.ConditionStatus `json:"status"`
	// Last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty"`
	// The reason for the condition's last transition.
	// +optional
	Reason *string `json:"reason,omitempty"`
	// A human readable message indicating details about the transition.
	// +optional
	Message *string `json:"message,omitempty"`
}

// NewCondition returns a new Condition instance.
func NewCondition(t ConditionType, status corev1.ConditionStatus, reason, message *string) Condition {
	return Condition{
		Type:               t,
		Status:             status,
		LastTransitionTime: &metav1.Time{},
		Reason:             reason,
		Message:            message,
	}
}

// NewSuccessfulCondition returns a new Successful Condition instance.
func NewSuccessfulCondition(message string) Condition {
	return NewCondition(ConditionTypeSuccessful, corev1.ConditionTrue, nil, &message)
}

// NewRecoverableCondition returns a new Recoverable Condition instance.
func NewRecoverableCondition(reason, message string) Condition {
	return NewCondition(ConditionTypeRecoverable, corev1.ConditionTrue, &reason, &message)
}

// NewFailedCondition returns a new Failed Condition instance.
func NewFailedCondition(reason, message string) Condition {
	return NewCondition(ConditionTypeFailed, corev1.ConditionTrue, &reason, &message)
}

func (c Condition) IsSuccessful() bool {
	return c.Type == ConditionTypeSuccessful
}

func (c Condition) IsRecoverable() bool {
	return c.Type == ConditionTypeRecoverable
}

func (c Condition) IsFailed() bool {
	return c.Type == ConditionTypeFailed
}

func HaveFailedCondition(conditions []Condition) bool {
	for _, c := range conditions {
		if c.IsFailed() {
			return true
		}
	}
	return false
}

func HaveRecoverableCondition(conditions []Condition) bool {
	for _, c := range conditions {
		if c.IsRecoverable() {
			return true
		}
	}
	return false
}

func HaveSuccessfulCondition(conditions []Condition) bool {
	for _, c := range conditions {
		if c.IsSuccessful() {
			return true
		}
	}
	return false
}

func GetFailedCondition(conditions []Condition) *Condition {
	for _, c := range conditions {
		if c.IsFailed() {
			return &c
		}
	}
	return nil
}

func GetRecoverableCondition(conditions []Condition) *Condition {
	for _, c := range conditions {
		if c.IsRecoverable() {
			return &c
		}
	}
	return nil
}

func GetSuccessfulCondition(conditions []Condition) *Condition {
	for _, c := range conditions {
		if c.IsSuccessful() {
			return &c
		}
	}
	return nil
}
