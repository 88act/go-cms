package system

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/88act/go-cms/server/global"
	"github.com/88act/go-cms/server/model/system"
)

type JwtService struct {
}

//@author: [88act](https://github.com/88act)
//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func (jwtService *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

//@author: [88act](https://github.com/88act)
//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
	//err := global.DB.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
	//isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	//return !isNotFound
}

//@author: [88act](https://github.com/88act)
//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: err error, redisJWT string

func (jwtService *JwtService) GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.REDIS.Get(context.Background(), userName).Result()
	return err, redisJWT
}

//@author: [88act](https://github.com/88act)
//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: jwt string, userName string
//@return: err error

func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.CONFIG.JWT.ExpiresTime) * time.Second
	err = global.REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

func LoadAll() {
	var data []string
	err := global.DB.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
}
