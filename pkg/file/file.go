package file

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
)

// GetSize 获取文件大小
func GetSize(f multipart.File) (int64, error) {
	// 使用 Seek 和 Stat 替代读取整个文件
	if seeker, ok := f.(io.Seeker); ok {
		// 保存当前位置
		current, err := seeker.Seek(0, io.SeekCurrent)
		if err != nil {
			return 0, err
		}
		// 获取文件大小
		size, err := seeker.Seek(0, io.SeekEnd)
		if err != nil {
			return 0, err
		}
		// 恢复位置
		_, err = seeker.Seek(current, io.SeekStart)
		if err != nil {
			return 0, err
		}
		return size, nil
	}
	
	// 如果不支持 Seek，则读取整个文件
	content, err := io.ReadAll(f)
	if err != nil {
		return 0, err
	}
	return int64(len(content)), nil
}

// GetExt 获取文件扩展名
func GetExt(fileName string) string {
	return path.Ext(filepath.Clean(fileName))
}

// CheckNotExist 检查文件是否不存在
func CheckNotExist(src string) bool {
	_, err := os.Stat(filepath.Clean(src))
	return os.IsNotExist(err)
}

// CheckPermission 检查文件权限
func CheckPermission(src string) bool {
	_, err := os.Stat(filepath.Clean(src))
	return os.IsPermission(err)
}

// IsNotExistMkDir 如果目录不存在则创建
func IsNotExistMkDir(src string) error {
	src = filepath.Clean(src)
	if CheckNotExist(src) {
		return MkDir(src)
	}
	return nil
}

// MkDir 创建目录
func MkDir(src string) error {
	src = filepath.Clean(src)
	return os.MkdirAll(src, os.ModePerm)
}

// Open 根据指定模式打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	name = filepath.Clean(name)
	return os.OpenFile(name, flag, perm)
}

// MustOpen 尝试打开文件，如果需要则创建目录
func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("获取工作目录失败: %w", err)
	}

	// 使用 filepath.Join 处理路径
	src := filepath.Join(dir, filePath)
	src = filepath.Clean(src)

	// 检查路径是否在工作目录下
	if !isSubPath(dir, src) {
		return nil, fmt.Errorf("非法的文件路径: %s", src)
	}

	if CheckPermission(src) {
		return nil, fmt.Errorf("没有权限访问目录: %s", src)
	}

	if err := IsNotExistMkDir(src); err != nil {
		return nil, fmt.Errorf("创建目录失败 %s: %w", src, err)
	}

	fullPath := filepath.Join(src, fileName)
	fullPath = filepath.Clean(fullPath)

	// 再次检查完整路径
	if !isSubPath(dir, fullPath) {
		return nil, fmt.Errorf("非法的文件路径: %s", fullPath)
	}

	f, err := Open(fullPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("打开文件失败: %w", err)
	}

	return f, nil
}

// isSubPath 检查 child 是否是 parent 的子路径
func isSubPath(parent, child string) bool {
	parent = filepath.Clean(parent)
	child = filepath.Clean(child)
	
	relative, err := filepath.Rel(parent, child)
	if err != nil {
		return false
	}
	
	// 检查是否包含 ".."
	if relative == ".." || relative[:3] == "../" {
		return false
	}
	
	return true
}
