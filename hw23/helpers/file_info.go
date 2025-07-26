package helpers

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileInfo struct {
	Path     string
	Filename string
	Size     int64
}

func NewFileInfo(path string) (*FileInfo, error) {
	f := &FileInfo{}
	err := f.SetPath(path)
	if err != nil {
		return nil, err
	}

	return f, nil
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
