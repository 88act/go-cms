package logic

import (
	"fmt"
	"go-cms/app/mqueue/cmd/job/jobtype"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

// scheduler job ------> go-zero-go-cms/app/mqueue/cmd/job/internal/logic/settleRecord.go.
func (l *MqueueScheduler) settleRecordScheduler() {

	task := asynq.NewTask(jobtype.ScheduleSettleRecord, nil)
	// every one minute exec
	entryID, err := l.svcCtx.Scheduler.Register("*/1 * * * *", task)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!MqueueSchedulerErr!!! ====> 【settleRecordScheduler】 registered  err:%+v , task:%+v", err, task)
	}
	fmt.Printf("【settleRecordScheduler】 registered an  entry: %q \n", entryID)
}
