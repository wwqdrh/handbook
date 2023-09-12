package http_client

import (
	"encoding/json"
	"net/http"
)

// Storage Interface Supports Get and Put
// of a single value
type Storage interface {
	Get() string
	Put(string)
}

// MemStorage implements Storage
type MemStorage struct {
	value string
}

// Get our in-memory value
func (m *MemStorage) Get() string {
	return m.value
}

// Put our in-memory value
func (m *MemStorage) Put(s string) {
	m.value = s
}

// Controller passes state to our handlers
type Controller struct {
	storage Storage
}

// New is a Controller 'constructor'
func New(storage Storage) *Controller {
	c := Controller{
		storage: storage,
	}
	return &c
}

// Payload is our common response
type Payload struct {
	Value string `json:"value"`
}

////////////////////
// get
////////////////////

// GetValue is a closure that wraps a HandlerFunc, if UseDefault
// is true value will always be "default" else it'll be whatever
// is stored in storage
func (c *Controller) GetValue(UseDefault bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		value := "default"
		if !UseDefault {
			value = c.storage.Get()
		}
		p := Payload{Value: value}
		w.WriteHeader(http.StatusOK)
		if payload, err := json.Marshal(p); err == nil {
			w.Write(payload)
		}
	}
}

// SetValue modifies the underlying storage of the controller object
func (c *Controller) SetValue(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	value := r.FormValue("value")
	c.storage.Put(value)
	w.WriteHeader(http.StatusOK)
	p := Payload{Value: value}
	if payload, err := json.Marshal(p); err == nil {
		w.Write(payload)
	}

}
