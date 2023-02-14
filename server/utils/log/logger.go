package log

import (
	"github.com/chanyipiaomiao/hlog"
	"time"
)

const (
	Url = "url"
	Err = "err"
)

// func Text(url string, err error) logrus.Fields {
// 	if err != nil {
// 		return logrus.Fields{
// 			"msgtype": "text",
// 			"text": map[string]string{
// 				"content": fmt.Sprintf("url: %s \nerror: %s", url, err.Error()),
// 			},
// 		}
// 	}
// 	return logrus.Fields{
// 		"msgtype": "text",
// 		"text": map[string]string{
// 			"content": fmt.Sprintf("url: %s \nerror: %s", url, ""),
// 		},
// 	}
// }

func init() {
	_, err := hlog.New(&hlog.Option{
		LogPath:            "/tmp/logs/hlog.log",
		LogType:            hlog.JSON,
		FileNameDateFormat: hlog.FileNameDateFormat,
		TimestampFormat:    hlog.TimestampFormat,
		LogLevel:           hlog.DebugLevel,
		MaxAge:             7 * 24 * time.Hour,
		RotationTime:       24 * time.Hour,
		JSONPrettyPrint:    true,
	})

	if err != nil {
		hlog.StderrFatalf("error: %s", err)
	}

	hlog.Debug(hlog.D{"hello": "world"}, "hello")
	hlog.Info(hlog.D{"hello": "world"}, "hello")
	hlog.Warn(hlog.D{"username": "warn"}, "呵呵")
	hlog.Error(hlog.D{"username": "Error"}, "呵呵")
}
