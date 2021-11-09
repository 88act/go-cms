package response

import "github.com/88act/go-cms/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
