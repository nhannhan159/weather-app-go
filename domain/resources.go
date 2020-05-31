package domain

import "go.uber.org/zap"

type Resources struct {
	DaoManager  IDaoManager
	HTTPManager IHttpManager
	GinLogger   *zap.SugaredLogger
	BizLogger   *zap.SugaredLogger
}
