package types_test

import (
	"testing"

	"github.com/vlaner/notes_share/internal/core/types"
	"github.com/vlaner/notes_share/internal/utils"
)

func TestNoteEncryptDecrypt(t *testing.T) {
	key, _ := utils.GenerateRandomBytes(32)
	content := "test content"
	note := types.Note{
		Title:         "test title",
		Content:       content,
		Encrypted:     true,
		EncryptionKey: key,
	}

	err := note.Encrypt(key)
	if err != nil {
		t.Fatal(err)
	}

	if note.Content == content {
		t.Fatal("content is not encrypted")
	}
	t.Log(note.Content)

	err = note.Decrypt(key)
	if err != nil {
		t.Fatal(err)
	}

	if note.Content != content {
		t.Fatal("content is not decrypted")
	}

	t.Log(note.Content)
}
