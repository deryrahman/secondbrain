package model

import "github.com/google/uuid"

type RecordOnStorage struct {
	ID      uuid.UUID
	Content string
}

type RecordTagOnStorage struct {
	RecordID uuid.UUID
	TagID    string
}

type TagOnStorage struct {
	ID string
}
