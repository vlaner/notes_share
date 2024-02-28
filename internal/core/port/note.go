package port

import (
	"github.com/google/uuid"
	"github.com/vlaner/notes_share/internal/core/types"
)

type NoteService interface {
	CreateNote(*types.Note) (*types.Note, error)
	GetNote(id uuid.UUID) (*types.Note, error)
	Decrypt(*types.Note, []byte) (*types.Note, error)
	Encrypt(*types.Note, []byte) (*types.Note, error)
}

type NoteRepository interface {
	CreateNote(*types.Note) (*types.Note, error)
	GetNote(id uuid.UUID) (*types.Note, error)
}
