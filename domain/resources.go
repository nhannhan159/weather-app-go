package domain

import "go.uber.org/zap"

type Resources struct {
	DaoManager  IDaoManager
	HTTPManager IHttpManager
	Logger      *zap.SugaredLogger
}
