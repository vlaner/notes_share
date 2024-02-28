package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vlaner/notes_share/internal/core/port"
	"github.com/vlaner/notes_share/internal/core/types"
)

type noteService struct {
	noteRepo port.NoteRepository
}

func NewNoteService(noteRepo port.NoteRepository) *noteService {
	return &noteService{
		noteRepo: noteRepo,
	}
}

func (s *noteService) CreateNote(note *types.Note) (*types.Note, error) {
	note, err := s.noteRepo.CreateNote(note)
	if err != nil {
		return nil, fmt.Errorf("error creating note: %w", err)
	}

	return note, nil
}

func (s *noteService) GetNote(id uuid.UUID) (*types.Note, error) {
	note, err := s.noteRepo.GetNote(id)
	if err != nil {
		return nil, fmt.Errorf("error getting note: %w", err)
	}

	return note, nil
}

func (s *noteService) Encrypt(note *types.Note, key []byte) (*types.Note, error) {
	note.EncryptionKey = key
	if err := note.Encrypt(key); err != nil {
		return nil, err
	}

	return note, nil
}

func (s *noteService) Decrypt(note *types.Note, key []byte) (*types.Note, error) {
	if err := note.Decrypt(key); err != nil {
		return nil, err
	}

	return note, nil
}
