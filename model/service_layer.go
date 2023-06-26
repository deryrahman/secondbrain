package model

type RecordOnService struct {
	ID      string
	Content string
	Tags    []TagOnService
}

type TagOnService string
