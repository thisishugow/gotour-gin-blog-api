package global

import (
	"yuwang.idv/go-tour/blog-api/pkg/logger"
	"yuwang.idv/go-tour/blog-api/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
