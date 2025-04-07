// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gdb

import (
	"context"
)

type (
	BeforeHookFuncSelect func(ctx context.Context, in *BeforeHookSelectInput) (err error)
	BeforeHookFuncInsert func(ctx context.Context, in *BeforeHookInsertInput) (err error)
	BeforeHookFuncUpdate func(ctx context.Context, in *BeforeHookUpdateInput) (err error)
	BeforeHookFuncDelete func(ctx context.Context, in *BeforeHookDeleteInput) (err error)
)

// BeforeHookHandler manages all supported hook functions for Model.
type BeforeHookHandler struct {
	Select BeforeHookFuncSelect
	Insert BeforeHookFuncInsert
	Update BeforeHookFuncUpdate
	Delete BeforeHookFuncDelete
}

// BeforeHookSelectInput holds the parameters for select hook operation.
// Note that, COUNT statement will also be hooked by this feature,
// which is usually not be interesting for upper business hook handler.
type BeforeHookSelectInput struct {
	handler BeforeHookFuncSelect
	Model   *Model // Current operation Model.
	Table   string // The table name that to be used. Update this attribute to change target table name.
}

// BeforeHookInsertInput holds the parameters for insert hook operation.
type BeforeHookInsertInput struct {
	handler BeforeHookFuncInsert
	Model   *Model // Current operation Model.
	Table   string // The table name that to be used. Update this attribute to change target table name.
}

// BeforeHookUpdateInput holds the parameters for update hook operation.
type BeforeHookUpdateInput struct {
	handler BeforeHookFuncUpdate
	Model   *Model // Current operation Model.
	Table   string // The table name that to be used. Update this attribute to change target table name.
}

// BeforeHookDeleteInput holds the parameters for delete hook operation.
type BeforeHookDeleteInput struct {
	handler BeforeHookFuncDelete
	Model   *Model // Current operation Model.
	Table   string // The table name that to be used. Update this attribute to change target table name.
}

// Next calls the next hook handler.
func (h *BeforeHookSelectInput) Next(ctx context.Context) (err error) {
	if h.handler != nil {
		safeOld := h.Model.safe
		h.Model.Safe(false)
		err := h.handler(ctx, h)
		h.Model.Safe(safeOld)
		return err
	}
	return nil
}

// Next calls the next hook handler.
func (h *BeforeHookInsertInput) Next(ctx context.Context) (err error) {
	if h.handler != nil {
		safeOld := h.Model.safe
		h.Model.Safe(false)
		err := h.handler(ctx, h)
		h.Model.Safe(safeOld)
		return err
	}
	return nil
}

// Next calls the next hook handler.
func (h *BeforeHookUpdateInput) Next(ctx context.Context) (err error) {
	if h.handler != nil {
		safeOld := h.Model.safe
		h.Model.Safe(false)
		err := h.handler(ctx, h)
		h.Model.Safe(safeOld)
		return err
	}
	return nil

}

// Next calls the next hook handler.
func (h *BeforeHookDeleteInput) Next(ctx context.Context) (err error) {
	if h.handler != nil {
		safeOld := h.Model.safe
		h.Model.Safe(false)
		err := h.handler(ctx, h)
		h.Model.Safe(safeOld)
		return err
	}
	return nil

}

// BeforeHook sets the hook functions for current model.
func (m *Model) BeforeHook(hook BeforeHookHandler) *Model {
	model := m.getModel()
	model.beforeHookHandler = hook
	return model
}
