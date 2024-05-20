package utils

import (
	"context"
	"errors"
	"log/slog"
)

var ErrLawyerIDNotFoundInContext = errors.New("lawyer not found in context")

func GetLawyerIDFromCtx(ctx context.Context) (uint, error) {
	lawyerID, ok := ctx.Value("lawyerID").(uint)
	slog.Info("getLawyerIDFromCtx", "lawyerID", lawyerID, "ok", ok)
	if !ok {
		return 0, ErrLawyerIDNotFoundInContext
	}

	return lawyerID, nil
}
