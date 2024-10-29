package constants

import (
	"fmt"
	"strings"
)

const (
	TransactionQueue                   = "TRANSACTION_QUEUE"
	TransactionWorkflowName            = "TransactionWorkflow"
	TransactionApproveWorkflowName     = "TransactionApproveWorkflow"
	TransactionRejectWorkflowName      = "TransactionRejectWorkflow"
	TransactionWorkflowSearchAttribute = "UUID"
	TransactionWorkflowStateIdentifier = "TransactionWorkflowState"
)

type WorkflowIdentifier interface {
	string | uint | int
}

func CreateWorkflowID[ID WorkflowIdentifier](workflowName string, ident ID) string {
	id := fmt.Sprintf("%s-%v", workflowName, ident)
	if len(id) > 250 {
		id = id[:250]
	}
	return id
}

func GetWorkflowNameFromID(workflowID string) string {
	return workflowID[:strings.Index(workflowID, "-")]
}
