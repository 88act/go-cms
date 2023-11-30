package mycache

import (
	"bytes"
	"encoding/gob"
	"errors"
	"sync"
	"time"

	"github.com/allegro/bigcache/v3"
)

type MyCache struct {
	Cache *bigcache.BigCache
}

var once_MyCache sync.Once = sync.Once{}
var obj_MyCache *MyCache

// 获取单例
func GetCache() *MyCache {
	once_MyCache.Do(func() {
		obj_MyCache = new(MyCache)
		obj_MyCache.Cache, _ = bigcache.NewBigCache(bigcache.DefaultConfig(120 * time.Minute))
		//instanse.init()
	})
	return obj_MyCache
}

// 设置
func (m *MyCache) Set(key string, value interface{}, tag string) error {
	// 将 value 序列化为 bytes
	//global.LOG.Info("Set cacheKey  =" + key)
	valueBytes, err := serialize(value)
	if err != nil {
		return err
	}
	return m.Cache.Set(key, valueBytes)
}

// 获取
func (m *MyCache) Get(key string) (interface{}, error) {
	return nil, errors.New("1111")
	// 获取以 bytes 格式存储的 value
	//global.LOG.Info("Get cacheKey  =" + key)
	valueBytes, err := m.Cache.Get(key)
	if err != nil {
		return nil, err
	}
	// 反序列化 valueBytes
	value, err := deserialize(valueBytes)
	if err != nil {
		return nil, err
	}
	//global.LOG.Info(" Get cacheKey  OK ")
	return value, nil
}

// 删除key
func (m *MyCache) Del(key string) error {
	return m.Cache.Delete(key)
}

// 删除Tag
func (m *MyCache) DelTag(tag string) error {
	//return m.Cache.Delete(key)
	return nil
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
