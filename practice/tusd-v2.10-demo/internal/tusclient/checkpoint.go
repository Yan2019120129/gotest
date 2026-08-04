package tusclient

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type FileFingerprint struct {
	Path            string `json:"path"`
	Size            int64  `json:"size"`
	ModTimeUnixNano int64  `json:"mod_time_unix_nano"`
}

type Checkpoint struct {
	Version     int               `json:"version"`
	Fingerprint FileFingerprint   `json:"fingerprint"`
	UploadURL   string            `json:"upload_url"`
	Offset      int64             `json:"offset"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

type CheckpointStore struct {
	Dir string
}

func NewCheckpointStore(dir string) *CheckpointStore {
	return &CheckpointStore{Dir: dir}
}

func (s *CheckpointStore) Load(fingerprint FileFingerprint) (*Checkpoint, error) {
	data, err := os.ReadFile(s.path(fingerprint))
	if err != nil {
		return nil, err
	}
	var checkpoint Checkpoint
	if err := json.Unmarshal(data, &checkpoint); err != nil {
		return nil, fmt.Errorf("decode checkpoint: %w", err)
	}
	if checkpoint.Version != 1 {
		return nil, fmt.Errorf("unsupported checkpoint version %d", checkpoint.Version)
	}
	if checkpoint.Fingerprint != fingerprint {
		return nil, errors.New("local file changed since checkpoint was created")
	}
	if checkpoint.UploadURL == "" {
		return nil, errors.New("checkpoint does not contain an upload URL")
	}
	return &checkpoint, nil
}

func (s *CheckpointStore) Save(checkpoint Checkpoint) error {
	if err := os.MkdirAll(s.Dir, 0o700); err != nil {
		return err
	}
	data, err := json.MarshalIndent(checkpoint, "", "  ")
	if err != nil {
		return err
	}
	path := s.path(checkpoint.Fingerprint)
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o600); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

func (s *CheckpointStore) Remove(fingerprint FileFingerprint) error {
	return os.Remove(s.path(fingerprint))
}

func (s *CheckpointStore) path(fingerprint FileFingerprint) string {
	sum := sha256.Sum256([]byte(fingerprint.Path))
	return filepath.Join(s.Dir, hex.EncodeToString(sum[:])+".json")
}
