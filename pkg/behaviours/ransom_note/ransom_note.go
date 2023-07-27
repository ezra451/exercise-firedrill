package ransom_note

import (
	"context"
	"encoding/base64"
	"os"
	"path/filepath"

	"github.com/ezra451/exercise-firedrill/pkg/sergeant"
	"github.com/ezra451/exercise-firedrill/pkg/utils/userinfo"
	"go.uber.org/zap"
)

const (
	ID   = "ransom_note"
	Name = "Ransomware note"

	ransomMessage      = `KkFsbCB5b3VyIGZpbGVzIGhhdmUgYmVlbiBlbmNyeXB0ZWQhKiAKCkFsbCB5b3VyIGZpbGVzIGhhdmUgYmVlbiBlbmNyeXB0ZWQgZHVlIHRvIGEgc2VjdXJpdHkgcHJvYmxlbSB3aXRoIHlvdSBhbmQgeW91ciBQQyBJRC4gCklmIHlvdSB3YW50IHRvIHJlc3RvcmUgdGhlbSwgY29udGFjdCB1cyB0byB0aGUgZS1tYWlsOiAwZzRjMWhjQHFja3IwMHZ5LmluZm8KCllvdSBoYXZlIHRvIHBheSBmb3IgZGVjcnlwdGlvbiBpbiBCaXRjb2lucy4gVGhlIHByaWNlIGRlcGVuZHMgb24gaG93IGZhc3QgeW91IHdyaXRlIHRvIHVzLiAKQWZ0ZXIgcGF5bWVudCwgd2Ugd2lsbCBzZW5kIHlvdSB0aGUgdG9vbCB0aGF0IHdpbGwgZGVjcnlwdCBhbGwgeW91ciBmaWxlcy4gRnJlZSBkZWNyeXB0aW9uIGFzIGd1YXJhbnRlZS4gQmVmb3JlIHBheWluZyB5b3UgY2FuIHNlbmQgdXMgdXAgdG8gNSBmaWxlcyBmb3IgZnJlZSBkZWNyeXB0aW9uLiAKVGhlIHRvdGFsIHNpemUgb2YgZmlsZXMgbXVzdCBiZSBsZXNzIHRoYW4gNE1iIChub24tYXJjaGl2ZWQpLgpBdHRlbnRpb24hIERvIG5vdCByZW5hbWUgZW5jcnlwdGVkIGZpbGVzLiAKRG8gbm90IHRyeSB0byBkZWNyeXB0IHlvdXIgZGF0YSB1c2luZyB0aGlyZC1wYXJ0eSBzb2Z0d2FyZSwgaXQgbWF5IGNhdXNlIHBlcm1hbmVudCBkYXRhIGxvc3MuIApUaGUgZGVjcnlwdGlvbiBvZiB5b3VyIGZpbGVzIHdpdGggdGhlIGhlbHAgb2YgdGhpcmQgcGFydGllcyBtYXkgY2F1c2UgaW5jcmVhc2VkIHByaWNlICh0aGV5IGFkZCB0aGVpciBmZWUpIG9yIHlvdSBjYW4gYmVjb21lIGEgdmljdGltIG9mIGEgc2NhbS4gQ2hlZXJzIQ==`
	ransomNoteFileName = "ransomnote.txt"
)

type RansomNoteOptions struct {
	NoteFileName string
	Note         string
}

type RansomNote struct {
	note         string
	noteFileName string
	base64       bool
}

func NewRansomNote(opts ...RansomNoteOptions) sergeant.Runnable {
	var options RansomNoteOptions = RansomNoteOptions{
		Note:         ransomMessage,
		NoteFileName: ransomNoteFileName,
	}

	if len(opts) > 0 {
		options = opts[0]
	}

	return &RansomNote{note: options.Note, noteFileName: options.NoteFileName, base64: true}
}

func (e *RansomNote) ID() string {
	return ID
}

func (e *RansomNote) Name() string {
	return Name
}

func (e *RansomNote) Run(ctx context.Context, logger *zap.Logger) error {
	desktopPath := userinfo.UserDesktop()
	logger.Sugar().Infof("User desktop path for ransom note: %s", desktopPath)

	ransomNoteFileName := e.noteFileName

	ransomNotePath := filepath.Join(desktopPath, ransomNoteFileName)

	if e.base64 {
		uDec, _ := base64.URLEncoding.DecodeString(e.note)

		e.note = string(uDec)
	}

	if err := os.WriteFile(ransomNotePath, []byte(e.note), 0644); err != nil {
		logger.Sugar().Warnf("Failed to drop ransom note at %s: %s", ransomNotePath, err.Error())
		return err
	}

	logger.Sugar().Infof("Dropped ransom note at %s", ransomNotePath)

	return nil
}
