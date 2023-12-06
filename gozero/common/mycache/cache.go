package mycache

import (
	"bytes"
	"encoding/gob"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/core/logx"
)

type MyCache struct {
	Cache      *redis.Client
	CacheTimes time.Duration
	//MyBigCache *bigcache.BigCache
	//MyRides    *redis.Client
}

var once_MyCache sync.Once = sync.Once{}
var obj_MyCache *MyCache

func InitObj(addr string, pwd string, db int) {
	once_MyCache.Do(func() {
		obj_MyCache = new(MyCache)
		//obj_MyCache.MyBigCache, _ = bigcache.NewBigCache(bigcache.DefaultConfig(120 * time.Minute))
		client := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: pwd, // no password set
			DB:       db,  // use default DB
		})
		obj_MyCache.Cache = client
		obj_MyCache.CacheTimes = 30 * time.Minute //2 * time.Hour // 2 小时
		//可连接性检测
		_, err := client.Ping().Result()
		if err != nil {
			logx.Errorf("redis链接错误 ,err =  %s\n", err.Error())
			os.Exit(0)
			return
		} else {
			logx.Info("redis链接成功 addr=", addr)
		}

	})
}

// 获取单例

func GetCache() *MyCache {
	return obj_MyCache
}

// func (m *MyCache) InitRedis() {
// 	//redisCfg := global.CONFIG.Redis
// 	client := redis.NewClient(&redis.Options{
// 		Addr:     "127.0.0.1:6759",
// 		Password: "", // no password set
// 		DB:       0,  // use default DB
// 	})
// 	m.MyRides = client
// 	// pong, err := client.Ping(context.Background()).Result()
// 	// if err != nil {
// 	// 	//logx.Error("redis connect ping failed, err:", zap.Any("err", err))
// 	// } else {
// 	// 	//logx.Error("redis connect ping response:", zap.String("pong", pong))
// 	// 	m.MyRides = client
// 	// }
// }

// 设置字符串，  注意 设置val 是字符串
func (m *MyCache) Set(key string, value interface{}, second int) error {
	// 将 value 序列化为 bytes
	//logx.Error("Set cacheKey  =" + key)
	var cmd *redis.StatusCmd
	if second > 0 {
		cmd = m.Cache.Set(key, value, time.Duration(second)*time.Second)
	} else {
		cmd = m.Cache.Set(key, value, m.CacheTimes)
	}
	if cmd.Err() != nil {
		logx.Errorf(" 缓存设置出错 %s", cmd.Err().Error())
		//panic(err)
	}
	return cmd.Err()
}

// 获取字符串， 注意返回 是字符串
func (m *MyCache) Get(key string) (interface{}, error) {
	// 获取以 bytes 格式存储的 value
	//logx.Error("Get cacheKey  =" + key)
	val, err := m.Cache.Get(key).Result()
	// switch {
	// case err == redis.Nil:
	// 	fmt.Println("key不存在")
	// case err != nil:
	// 	fmt.Println("错误", err)
	// case val == "":
	// 	fmt.Println("值是空字符串")
	// }
	return val, err
}

// 设置 struct 对象 , 缓存时间 分钟
func (m *MyCache) SetObj(key string, value interface{}, second int) error {
	// 将 value 序列化为 bytes
	//logx.Error(" SetObj cacheKey  =" + key)
	//转成byte[]
	valueBytes, err := serialize(value)
	if err != nil {
		return err
	}
	var cmd *redis.StatusCmd
	if second > 0 {
		cmd = m.Cache.Set(key, valueBytes, time.Duration(second)*time.Second)
	} else {
		cmd = m.Cache.Set(key, valueBytes, m.CacheTimes)
	}
	if cmd.Err() != nil {
		logx.Errorf(" 缓存设置出错 %s", cmd.Err().Error())
		//panic(err)
	}
	return cmd.Err()
}

// 获取 struct 对象
func (m *MyCache) GetObj(key string) (interface{}, error) {
	// 获取以 bytes 格式存储的 value

	cmd := m.Cache.Get(key)
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	// 反序列化 valueBytes
	value, err := deserialize([]byte(cmd.Val()))
	if err != nil {
		return nil, err
	}
	//logx.Error("GetObj cacheKey  =" + key)
	return value, err
}

// 删除key
func (m *MyCache) Delete(key string) error {
	err := m.Cache.Del(key).Err()
	return err
}

// 删除Tag
func (m *MyCache) DelTag(tag string) error {
	//return m.Cache.Delete(key)
	return nil
}

// 删除具有特定前缀的多个键
func (m *MyCache) DeleteKeyPre(prefix string) (num int, err error) {
	keysToDelete, err := m.Cache.Keys(prefix + "*").Result()
	if err != nil {
		return 0, err
	}
	num = len(keysToDelete)
	if num > 0 {
		// 使用pipeline批量删除所有匹配的键
		pipeline := m.Cache.Pipeline()
		for _, key := range keysToDelete {
			logx.Errorf(" 批量删除 key =%s ", key)
			pipeline.Del(key)
		}
		_, err = pipeline.Exec()
		if err != nil {
			logx.Errorf(" 批量删除 err =%s ", err.Error())
			return 0, err
		}
	}
	return num, nil
}

// https://www.freeaihub.com/post/106487.html   Gob 是一种 Go 原生的数据序列化方式，它还具有序列化 interface 类型的功能
func serialize(value interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	gob.Register(value)

	err := enc.Encode(&value)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func deserialize(valueBytes []byte) (interface{}, error) {
	var value interface{}
	buf := bytes.NewBuffer(valueBytes)
	dec := gob.NewDecoder(buf)

	err := dec.Decode(&value)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// func (m *MyCache) Set(key string, obj []byte) {
// 	m.Cache.Set(key, obj)
// }

// func (m *MyCache) Get(key string) (obj []byte, err error) {
// 	obj, err = m.Cache.Get(key)
// 	return obj, err
// }

//==== struct 对象与内存byte[] 互转==================
// var sizeOfMyStruct = int(unsafe.Sizeof(SysUsers{}))

// func (m *SysUsers) ToBytes() []byte {
// 	var x reflect.SliceHeader
// 	x.Len = sizeOfMyStruct
// 	x.Cap = sizeOfMyStruct
// 	x.Data = uintptr(unsafe.Pointer(m))
// 	return *(*[]byte)(unsafe.Pointer(&x))
// }

// func (m *SysUsers) BytesToMy(b []byte) *SysUsers {
// 	return (*SysUsers)(unsafe.Pointer(
// 		(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
// 	))
// }
