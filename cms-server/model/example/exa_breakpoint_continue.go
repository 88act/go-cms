package example

import (
	"github.com/88act/go-cms/server/global"
)

// file struct, 文件结构体
type ExaFile struct {
	global.GO_MODEL
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
}

// file chunk struct, 切片结构体
type ExaFileChunk struct {
	global.GO_MODEL
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string
}
