package helpers

import "fmt"

type Stat struct {
	fi           *FileInfo
	isDecompress bool
	action       string
	from         string
	to           string
}

func NewStat(isDecompress bool) *Stat {
	action := "Compression"
	if isDecompress {
		action = "Decompression"
	}
	return &Stat{fi: NewFileInfo(), isDecompress: isDecompress, action: action}
}

func (s *Stat) Print(from, to string) error {
	sizeFrom, err := s.PrintSize(from, "Source file")
	if err != nil {
		return err
	}

	sizeTo, err := s.PrintSize(to, s.action)
	if err != nil {
		return err
	}

	s.PrintRation(sizeFrom, sizeTo)
	return nil
}

func (s *Stat) PrintSize(file, msg string) (int64, error) {
	err := s.fi.SetPath(file)
	if err != nil {
		return 0, err
	}
	fmt.Printf("%-15s %-20s| %6s\n", msg, s.fi.Filename, s.fi.HumanSize())
	return s.fi.Size, nil
}

func (s *Stat) PrintRation(sizeFrom, sizeTo int64) {
	ratio := s.fi.RatioPercent(sizeFrom, sizeTo)
	if s.isDecompress {
		ratio = s.fi.RatioPercent(sizeTo, sizeFrom)
	}
	fmt.Printf("%-35s | %6.2f%%\n", "Ratio % ", ratio)
}
