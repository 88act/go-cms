package example

import (
	"errors"

	"go-cms/global"
	"go-cms/model/example"

	"gorm.io/gorm"
)

type FileUploadAndDownloadService struct {
}

//@author: [88act](https://github.com/88act)
//@function: FindOrCreateFile
//@description: 上传文件时检测当前文件属性，如果没有文件则创建，有则返回文件的当前切片
//@param: fileMd5 string, fileName string, chunkTotal int
//@return: err error, file model.ExaFile

func (e *FileUploadAndDownloadService) FindOrCreateFile(fileMd5 string, fileName string, chunkTotal int) (err error, file example.ExaFile) {
	var cfile example.ExaFile
	cfile.FileMd5 = fileMd5
	cfile.FileName = fileName
	cfile.ChunkTotal = chunkTotal

	if errors.Is(global.DB.Where("file_md5 = ? AND is_finish = ?", fileMd5, true).First(&file).Error, gorm.ErrRecordNotFound) {
		err = global.DB.Where("file_md5 = ? AND file_name = ?", fileMd5, fileName).Preload("ExaFileChunk").FirstOrCreate(&file, cfile).Error
		return err, file
	}
	cfile.IsFinish = true
	cfile.FilePath = file.FilePath
	err = global.DB.Create(&cfile).Error
	return err, cfile
}

//@author: [88act](https://github.com/88act)
//@function: CreateFileChunk
//@description: 创建文件切片记录
//@param: id uint, fileChunkPath string, fileChunkNumber int
//@return: error

func (e *FileUploadAndDownloadService) CreateFileChunk(id uint, fileChunkPath string, fileChunkNumber int) error {
	var chunk example.ExaFileChunk
	chunk.FileChunkPath = fileChunkPath
	chunk.ExaFileID = id
	chunk.FileChunkNumber = fileChunkNumber
	err := global.DB.Create(&chunk).Error
	return err
}

//@author: [88act](https://github.com/88act)
//@function: DeleteFileChunk
//@description: 删除文件切片记录
//@param: fileMd5 string, fileName string, filePath string
//@return: error

func (e *FileUploadAndDownloadService) DeleteFileChunk(fileMd5 string, fileName string, filePath string) error {
	var chunks []example.ExaFileChunk
	var file example.ExaFile
	err := global.DB.Where("file_md5 = ? AND file_name = ?", fileMd5, fileName).First(&file).Update("IsFinish", true).Update("file_path", filePath).Error
	if err != nil {
		return err
	}
	err = global.DB.Where("exa_file_id = ?", file.ID).Delete(&chunks).Unscoped().Error
	return err
}
