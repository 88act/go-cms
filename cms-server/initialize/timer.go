package initialize

import (
	"fmt"

	"github.com/88act/go-cms/server/config"
	"github.com/88act/go-cms/server/global"
	"github.com/88act/go-cms/server/utils"
)

func Timer() {
	if global.CONFIG.Timer.Start {
		for i := range global.CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				global.Timer.AddTaskByFunc("ClearDB", global.CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})
			}(global.CONFIG.Timer.Detail[i])
		}
	}
}
