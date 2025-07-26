package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileInfo struct {
	Path     string
	Filename string
	Size     int64
}

func NewFileInfo() *FileInfo {
	return &FileInfo{}
}

func (f *FileInfo) SetPath(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	f.Path = path
	f.Filename = filepath.Base(path)
	f.Size = info.Size()
	return nil
}

func (f *FileInfo) HumanSize() string {
	const (
		_          = iota
		KB float64 = 1 << (10 * iota)
		MB
		GB
		TB
	)

	fs := float64(f.Size)

	switch {
	case fs >= TB:
		return fmt.Sprintf("%.2f ТБ", fs/TB)
	case fs >= GB:
		return fmt.Sprintf("%.2f ГБ", fs/GB)
	case fs >= MB:
		return fmt.Sprintf("%.2f МБ", fs/MB)
	case fs >= KB:
		return fmt.Sprintf("%.2f кБ", fs/KB)
	default:
		return fmt.Sprintf("%d Б", f.Size)
	}
}

func (f *FileInfo) ReplaceExt(name, ext string) string {
	fileExt := filepath.Ext(name)
	return fmt.Sprintf("%s.%s", strings.TrimSuffix(name, fileExt), ext)
}

func (f *FileInfo) RatioPercent(from, to int64) float64 {
	return (float64(to) / float64(from)) * 100
}
