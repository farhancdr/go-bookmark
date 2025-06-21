package bookmarks

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	config "github.com/farhancdr/go-bookmark/internal"
)

type Bookmark struct {
	Alias     string    `json:"alias"`
	Path      string    `json:"path"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BookmarkStore struct {
	Bookmarks []Bookmark `json:"bookmarks"`
	mutex     sync.Mutex
}

func LoadBookmarks() (*BookmarkStore, error) {
	configDir, err := config.GetConfigDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get config dir: %v", err)
	}

	filePath := filepath.Join(configDir, "bookmarks.json")
	data, err := os.ReadFile(filePath)
	if os.IsNotExist(err) {
		return &BookmarkStore{Bookmarks: []Bookmark{}}, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to read bookmarks file: %v", err)
	}

	var store BookmarkStore
	if err := json.Unmarshal(data, &store); err != nil {
		return nil, fmt.Errorf("failed to parse bookmarks file: %v", err)
	}

	return &store, nil
}

func (s *BookmarkStore) Save() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	configDir, err := config.GetConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get config dir: %v", err)
	}

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config dir: %v", err)
	}

	filePath := filepath.Join(configDir, "bookmarks.json")
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal bookmarks: %v", err)
	}

	return os.WriteFile(filePath, data, 0644)
}

func (s *BookmarkStore) FindByAlias(alias string) (Bookmark, bool) {
	for _, b := range s.Bookmarks {
		if b.Alias == alias {
			return b, true
		}
	}
	return Bookmark{}, false
}

func (s *BookmarkStore) AddBookmark(alias, path string) error {
	s.mutex.Lock()
	now := time.Now()
	var updated bool
	for i, b := range s.Bookmarks {
		if b.Alias == alias {
			s.Bookmarks[i] = Bookmark{
				Alias:     alias,
				Path:      path,
				CreatedAt: b.CreatedAt,
				UpdatedAt: now,
			}
			updated = true
			break
		}
	}
	if !updated {
		s.Bookmarks = append(s.Bookmarks, Bookmark{
			Alias:     alias,
			Path:      path,
			CreatedAt: now,
			UpdatedAt: now,
		})
	}
	s.mutex.Unlock()

	return s.Save()
}

func (s *BookmarkStore) DeleteBookmark(alias string) error {
	s.mutex.Lock()
	for i, b := range s.Bookmarks {
		if b.Alias == alias {
			s.Bookmarks = append(s.Bookmarks[:i], s.Bookmarks[i+1:]...)
			s.mutex.Unlock()
			return s.Save()
		}
	}
	s.mutex.Unlock()
	return nil
}

func ClearBookmarks() error {
	store := &BookmarkStore{Bookmarks: []Bookmark{}}
	return store.Save()
}

func (s *BookmarkStore) UpdateBookmark(alias, newPath string) error {
	s.mutex.Lock()
	for i, b := range s.Bookmarks {
		if b.Alias == alias {
			s.Bookmarks[i].Path = newPath
			s.Bookmarks[i].UpdatedAt = time.Now()
			s.mutex.Unlock()
			return s.Save()
		}
	}
	s.mutex.Unlock()
	return nil
}

func (s *BookmarkStore) RenameBookmark(oldAlias, newAlias string) error {
	s.mutex.Lock()
	for i, b := range s.Bookmarks {
		if b.Alias == oldAlias {
			s.Bookmarks[i].Alias = newAlias
			s.Bookmarks[i].UpdatedAt = time.Now()
			s.mutex.Unlock()
			return s.Save()
		}
	}
	s.mutex.Unlock()
	return nil
}
