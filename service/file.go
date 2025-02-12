package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"time"
)

var FileS = &FileService{}

type FileService struct {
	BaseService
}

func (s *FileService) UploadFile(c *gin.Context) {
	// 获取上传的文件
	file, _ := c.FormFile("file") // 获取文件表单字段 "file"
	if file == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// 保存文件到服务器
	filename := filepath.Join("./files", file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 获取最新的文件列表
	files, err := getFiles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回文件的下载链接
	c.JSON(http.StatusOK, gin.H{
		"message":  "File uploaded successfully!",
		"file_url": "/download/" + file.Filename,
		"files":    files, // 返回最新的文件列表
	})
}

func (s *FileService) Index(c *gin.Context) {
	files, err := getFiles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 模板名为新定义的模板名字
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"files": files})
}
func (s *FileService) Download(c *gin.Context) {
	filename := c.Param("filename")
	filePath := "./files/" + filename // 你可以根据实际需求修改路径

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// 下载文件
	c.File(filePath)
}

type FileInfo struct {
	Name    string
	Path    string
	ModTime time.Time
}

func getFiles() ([]FileInfo, error) {
	var files []FileInfo

	// 获取当前目录下所有文件
	dir := "./files" // 当前目录
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 排除目录，只列出文件
		if !info.IsDir() {
			files = append(files, FileInfo{Name: info.Name(), Path: path, ModTime: info.ModTime()})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	// 按照 ModTime 倒序排序文件列表
	sort.Slice(files, func(i, j int) bool {
		// 倒序排列：如果 ModTime 比较小，则排在后面
		return files[i].ModTime.After(files[j].ModTime)
	})
	return files, nil
}
