package globals

import "go.uber.org/zap"

var (
	Logger        *zap.Logger        = nil
	SugaredLogger *zap.SugaredLogger = nil
)
