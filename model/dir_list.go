package model

import "io"

type DirRecord struct {
	Path  string
	Count uint
}

type DirListService interface {
	GetTopDirList(top uint) []string
	ImportUrlList(file io.Reader) error
}
