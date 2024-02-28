package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/vlaner/notes_share/internal/adapter/storage/sqlite"
	http_transport "github.com/vlaner/notes_share/internal/adapter/transport/http"
	"github.com/vlaner/notes_share/internal/core/service"
	"github.com/vlaner/notes_share/pkg/html"
)

func main() {
	hmtlRend, err := html.NewHtmlRenderer("./www/templates")
	if err != nil {
		log.Fatalln("error parsing html templates:", err)
	}

	db, err := sql.Open("sqlite3", "./main.db")
	if err != nil {
		log.Fatalln("error opening sqlite db:", err)
	}
	defer db.Close()

	sqliteStorage := sqlite.NewNoteSqliteStorage(db)
	notesSvc := service.NewNoteService(sqliteStorage)
	notesHandler := http_transport.NewNoteHandler(notesSvc, hmtlRend)

	mux := http.NewServeMux()
	h := http_transport.NewHandler(hmtlRend, notesHandler)
	h.RegisterRoutes("./www/static", mux)

	srv := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	sig := <-signalCh
	log.Printf("Received signal: %v\n", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v\n", err)
	}

	log.Println("Server shutdown gracefully")
}
