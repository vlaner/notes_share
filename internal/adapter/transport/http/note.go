package http_transport

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/vlaner/notes_share/internal/core/port"
	"github.com/vlaner/notes_share/internal/core/types"
	"github.com/vlaner/notes_share/internal/utils"
	"github.com/vlaner/notes_share/pkg/html"
)

type NoteHandler struct {
	noteSvc      port.NoteService
	htmlRenderer *html.Renderer
}

func NewNoteHandler(noteSvc port.NoteService, htmlRenderer *html.Renderer) *NoteHandler {
	return &NoteHandler{
		noteSvc:      noteSvc,
		htmlRenderer: htmlRenderer,
	}
}

type createNoteReq struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	Encrypted bool   `json:"encrypted,string"`
}

type createNoteResp struct {
	Id         uuid.UUID `json:"id"`
	Encrypted  bool      `json:"encrypted"`
	EncryptKey string    `json:"encryption_key,omitempty"`
}

func (nh *NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var req createNoteReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	note, err := types.NewNote(uuid.New(), req.Title, req.Content, req.Encrypted, nil)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if req.Encrypted {
		encKey, err := utils.GenerateRandomBytes(32)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		note, err = nh.noteSvc.Encrypt(note, encKey)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}

	createdNote, err := nh.noteSvc.CreateNote(note)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := WriteJsonResponse(w, 201, Response{Success: true, Data: createNoteResp{
		Id:         createdNote.Id,
		Encrypted:  req.Encrypted,
		EncryptKey: base64.RawURLEncoding.EncodeToString(createdNote.EncryptionKey),
	}}); err != nil {
		log.Println("error writing response:", err)
	}
}

func (nh *NoteHandler) GetNote(w http.ResponseWriter, r *http.Request) {
	noteID := r.PathValue("id")
	uid, err := uuid.Parse(noteID)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	note, err := nh.noteSvc.GetNote(uid)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	if note == nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	noteCpy := *note

	if note.Encrypted {
		enckey := r.URL.Query().Get("key")
		encKey, err := base64.RawURLEncoding.DecodeString(enckey)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		note, err = nh.noteSvc.Decrypt(&noteCpy, []byte(encKey))
		if err != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
	}

	err = nh.htmlRenderer.Render(w, "note", note)
	if err != nil {
		log.Println("error rendering", err)
	}
}
