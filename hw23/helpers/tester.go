package helpers

type Tester struct{}

func (t *Tester) GetTestDataFiles() []struct{ File string } {
	return []struct {
		File string
	}{
		{File: "../testdata/alice.txt"},
		{File: "../testdata/text_en.txt"},
		{File: "../testdata/text_pl.txt"},
		{File: "../testdata/text_ru.txt"},
		//{File: "../testdata/text_en_415kb.txt"},
		//{File: "../testdata/text_pl_562kb.txt"},
		//{File: "../testdata/text_ru_396kb.txt"},
		{File: "../testdata/archive.zip"},
		//{File: "../testdata/text_pdf.pdf"},
		//{File: "../testdata/image.jpg"},
		{File: "../testdata/nature.jpg"},
		//{File: "../testdata/till_down.mp3"},
		//{File: "../testdata/video.mp4"},
	}
}
