package ctxdata

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

// CtxKeyJwtUserId get uid from ctx
var CtxKeyJwtUserId = "jwtUserId"
var CtxKeyJwtUserType = "jwtUserType"
var CtxKeyCuId = "jwtCuId"

func GetUidFromCtx(ctx context.Context) int64 {
	var uid int64
	uid = 0
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
		}
	}
	return uid
}

func GetUTypeFromCtx(ctx context.Context) int64 {
	var uid int64
	uid = 0
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserType).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("GetUTypeFromCtx err : %+v", err)
		}
	}
	return uid
}
func GeCuIdFormCtx(ctx context.Context) int64 {
	var uid int64
	uid = 0
	if jsonUid, ok := ctx.Value(CtxKeyCuId).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("GeCuIdFormCtx err : %+v", err)
		}
	}
	return uid
}
func GetUserLogInfo(ctx context.Context) string {
	uid := GetUidFromCtx(ctx)
	ut := GetUTypeFromCtx(ctx)
	cuid := GeCuIdFormCtx(ctx)
	return fmt.Sprintf("{'uid':%d,'ut':%d,'cuid':%d}", uid, ut, cuid)
}
