package pgxadapter

import (
	"context"

	"github.com/casbin/casbin/v2/model"
)

// ContextAdapter provides a context-aware interface for Casbin adapters.
type ContextAdapter interface {
	// LoadPolicyCtx loads all policy rules from the storage with context.
	LoadPolicyCtx(ctx context.Context, model model.Model) error
	// SavePolicyCtx saves all policy rules to the storage with context.
	SavePolicyCtx(ctx context.Context, model model.Model) error

	// AddPolicyCtx adds a policy rule to the storage with context.
	// This is part of the Auto-Save feature.
	AddPolicyCtx(ctx context.Context, sec string, ptype string, rule []string) error
	// RemovePolicyCtx removes a policy rule from the storage with context.
	// This is part of the Auto-Save feature.
	RemovePolicyCtx(ctx context.Context, sec string, ptype string, rule []string) error
	// RemoveFilteredPolicyCtx removes policy rules that match the filter from the storage with context.
	// This is part of the Auto-Save feature.
	RemoveFilteredPolicyCtx(ctx context.Context, sec string, ptype string, fieldIndex int, fieldValues ...string) error
}
