package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mailru/easyjson"
	"github.com/sanposhiho/gomockhandler/internal/model"
)

type Repository struct{}

func NewRepository() Repository {
	return Repository{}
}

func (r *Repository) Put(m *model.Config, path string) error {
	d, err := easyjson.Marshal(m)
	if err != nil {
		return fmt.Errorf("easyjson marshal: %w", err)
	}

	var buf bytes.Buffer
	if err := json.Indent(&buf, d, "", "	"); err != nil {
		return fmt.Errorf("format json: %w", err)
	}
	return ioutil.WriteFile(path, buf.Bytes(), 0644)
}

func (r *Repository) Get(path string) (*model.Config, error) {
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	var m model.Config
	err = easyjson.Unmarshal(raw, &m)
	if err != nil {
		return nil, fmt.Errorf("easyjson unmarshal: %w", err)
	}

	return &m, err
}
