package temporal

import (
	"github.com/hjoshi123/temporal-loan-app/pkg/constants"
	"go.temporal.io/api/enums/v1"
	temporal2 "go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func GetChildWorkflowOptions(ctx workflow.Context, uuid string, workflowName string) workflow.Context {
	return workflow.WithChildOptions(ctx, workflow.ChildWorkflowOptions{
		ParentClosePolicy:     enums.PARENT_CLOSE_POLICY_ABANDON,
		WorkflowIDReusePolicy: enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE,
		TypedSearchAttributes: temporal2.NewSearchAttributes(
			temporal2.NewSearchAttributeKeyString(constants.TransactionWorkflowSearchAttribute).ValueSet(uuid),
		),
		WorkflowID: constants.CreateWorkflowID(workflowName, uuid),
	})
}
