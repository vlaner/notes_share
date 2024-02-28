package types

import (
	"errors"

	"github.com/google/uuid"
	"github.com/vlaner/notes_share/internal/utils"
)

type Note struct {
	Id            uuid.UUID `json:"id"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	Encrypted     bool      `json:"encrypted"`
	EncryptionKey []byte    `json:"-"`
}

func NewNote(id uuid.UUID, title string, content string, encrypted bool, encryptionKey []byte) (*Note, error) {
	if len(title) == 0 {
		return nil, errors.New("title cannot be empty")
	}

	if len(content) == 0 {
		return nil, errors.New("content cannot be empty")
	}

	return &Note{
		Id:            id,
		Title:         title,
		Content:       content,
		Encrypted:     encrypted,
		EncryptionKey: encryptionKey,
	}, nil

}

func (n *Note) Encrypt(encryptionKey []byte) error {
	encContent, err := utils.EncryptData([]byte(n.Content), encryptionKey)
	if err != nil {
		return nil
	}

	n.Content = string(encContent)

	return nil

}

func (n *Note) Decrypt(decryptionKey []byte) error {
	decContent, err := utils.DecryptData([]byte(n.Content), decryptionKey)
	if err != nil {
		return err
	}

	n.Content = string(decContent)

	return nil
}
