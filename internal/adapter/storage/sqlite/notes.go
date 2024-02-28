package sqlite

import (
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vlaner/notes_share/internal/core/types"
)

type NoteSqliteStorage struct {
	conn *sql.DB
}

func NewNoteSqliteStorage(conn *sql.DB) *NoteSqliteStorage {
	return &NoteSqliteStorage{conn: conn}
}

func (s *NoteSqliteStorage) CreateNote(note *types.Note) (*types.Note, error) {
	_, err := s.conn.Exec("INSERT INTO notes (id, title, content, encrypted, encryption_key) VALUES (?, ?, ?, ?, ?)", note.Id, note.Title, note.Content, note.Encrypted, note.EncryptionKey)
	return note, err
}

func (s *NoteSqliteStorage) GetNote(id uuid.UUID) (*types.Note, error) {
	var note types.Note
	var ID string

	err := s.conn.QueryRow("SELECT id, title, content, encrypted, encryption_key FROM notes WHERE id = ?", id).Scan(&ID, &note.Title, &note.Content, &note.Encrypted, &note.EncryptionKey)
	if err != nil {
		return nil, err
	}

	uid, err := uuid.Parse(ID)
	if err != nil {
		return nil, err
	}

	note.Id = uid

	return &note, nil
}
