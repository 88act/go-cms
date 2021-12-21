package response

import "go-cms/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
