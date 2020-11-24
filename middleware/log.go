package middleware

import (
	"gitlab.stagingvip.net/publicGroup/public/common"
)

func OutLog(logName, logStr string) {
	common.LogsWithFileName(logPath, logName, logStr)
}
