package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Albicocca224/Practice5/internal/model"
	"github.com/Albicocca224/Practice5/internal/repository"
)

type Handler struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Handler {
	return &Handler{repo: repo}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

// GET /users?page=1&page_size=10&order_by=name&order_dir=asc&name=alice&gender=male
func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	page, _ := strconv.Atoi(q.Get("page"))
	if page < 1 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(q.Get("page_size"))
	if pageSize < 1 {
		pageSize = 10
	}

	filter := model.UserFilter{
		Page:     page,
		PageSize: pageSize,
		OrderBy:  q.Get("order_by"),
		OrderDir: q.Get("order_dir"),
	}

	if v := q.Get("id"); v != "" {
		id, err := strconv.Atoi(v)
		if err != nil {
			writeError(w, http.StatusBadRequest, "invalid id")
			return
		}
		filter.ID = &id
	}
	if v := q.Get("name"); v != "" {
		filter.Name = &v
	}
	if v := q.Get("email"); v != "" {
		filter.Email = &v
	}
	if v := q.Get("gender"); v != "" {
		filter.Gender = &v
	}
	if v := q.Get("birth_date"); v != "" {
		t, err := time.Parse("2006-01-02", v)
		if err != nil {
			writeError(w, http.StatusBadRequest, "birth_date must be YYYY-MM-DD")
			return
		}
		filter.BirthDate = &t
	}

	resp, err := h.repo.GetPaginatedUsers(filter)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, resp)
}

// GET /users/common-friends?user1=1&user2=2
func (h *Handler) GetCommonFriends(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	user1, err := strconv.Atoi(q.Get("user1"))
	if err != nil || user1 < 1 {
		writeError(w, http.StatusBadRequest, "user1 must be a positive integer")
		return
	}
	user2, err := strconv.Atoi(q.Get("user2"))
	if err != nil || user2 < 1 {
		writeError(w, http.StatusBadRequest, "user2 must be a positive integer")
		return
	}
	if user1 == user2 {
		writeError(w, http.StatusBadRequest, "user1 and user2 must be different")
		return
	}

	friends, err := h.repo.GetCommonFriends(user1, user2)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"user1":          user1,
		"user2":          user2,
		"common_friends": friends,
		"count":          len(friends),
	})
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/users", h.GetUsers)
	mux.HandleFunc("/users/common-friends", h.GetCommonFriends)
}
