package memory

import (
	"github.com/google/uuid"
	"github.com/vlaner/notes_share/internal/core/types"
)

type NoteMemoryStorage struct {
	Notes []*types.Note
}

func NewNoteStorage() *NoteMemoryStorage {
	return &NoteMemoryStorage{
		Notes: []*types.Note{},
	}
}

func (s *NoteMemoryStorage) GetNote(id uuid.UUID) (*types.Note, error) {
	for _, note := range s.Notes {
		if note.Id == id {
			return note, nil
		}
	}
	return nil, nil
}

func (s *NoteMemoryStorage) CreateNote(note *types.Note) (*types.Note, error) {
	s.Notes = append(s.Notes, note)
	return note, nil
}
