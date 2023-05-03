package service

import (
	"bufio"
	"fmt"
	"io"
	"penta/model"
)

type dirListService struct{}

func NewDirListService() model.DirListService {
	return &dirListService{}
}

func (s *dirListService) GetTopDirList(top uint) (r []string) {
	return r
}

func (s *dirListService) ImportUrlList(file io.Reader) error {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}
	return nil
}
