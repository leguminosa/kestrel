package xcontext

import (
	"context"

	"github.com/leguminosa/kestrel/pkg/util/convert"
)

func UserIDFromContext(ctx context.Context) int {
	return convert.ToInt(ctx.Value(ctxKeyUserID))
}
