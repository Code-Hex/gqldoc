package graphql

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetOperationContext(t *testing.T) {
	rc := &OperationContext{}

	ctx := WithOperationContext(context.Background(), rc)

	got := GetOperationContext(ctx)

	if diff := cmp.Diff(rc, got); diff != "" {
		t.Errorf("(-want, +got)\n%s", diff)
	}
}
