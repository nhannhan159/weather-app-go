package domain

import "go.uber.org/zap"

type Resources struct {
	DaoManager  IDaoManager
	HTTPManager IHttpManager
	GRPCManager IGRPCManager
	GinLogger   *zap.SugaredLogger
	BizLogger   *zap.SugaredLogger
}
