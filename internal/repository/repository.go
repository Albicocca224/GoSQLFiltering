package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Albicocca224/Practice5/internal/model"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

var allowedColumns = map[string]bool{
	"id":         true,
	"name":       true,
	"email":      true,
	"gender":     true,
	"birth_date": true,
}

func (r *Repository) GetPaginatedUsers(f model.UserFilter) (model.PaginatedResponse, error) {
	args := []interface{}{}
	conditions := []string{}
	argIdx := 1

	if f.ID != nil {
		conditions = append(conditions, fmt.Sprintf("id = $%d", argIdx))
		args = append(args, *f.ID)
		argIdx++
	}
	if f.Name != nil {
		conditions = append(conditions, fmt.Sprintf("name ILIKE $%d", argIdx))
		args = append(args, "%"+*f.Name+"%")
		argIdx++
	}
	if f.Email != nil {
		conditions = append(conditions, fmt.Sprintf("email ILIKE $%d", argIdx))
		args = append(args, "%"+*f.Email+"%")
		argIdx++
	}
	if f.Gender != nil {
		conditions = append(conditions, fmt.Sprintf("gender = $%d", argIdx))
		args = append(args, *f.Gender)
		argIdx++
	}
	if f.BirthDate != nil {
		conditions = append(conditions, fmt.Sprintf("birth_date = $%d", argIdx))
		args = append(args, *f.BirthDate)
		argIdx++
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	orderCol := "id"
	if f.OrderBy != "" && allowedColumns[f.OrderBy] {
		orderCol = f.OrderBy
	}
	orderDir := "ASC"
	if strings.ToUpper(f.OrderDir) == "DESC" {
		orderDir = "DESC"
	}

	var totalCount int
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM users %s", whereClause)
	err := r.db.QueryRow(countQuery, args...).Scan(&totalCount)
	if err != nil {
		return model.PaginatedResponse{}, fmt.Errorf("count query: %w", err)
	}

	offset := (f.Page - 1) * f.PageSize
	query := fmt.Sprintf(
		`SELECT id, name, email, gender, birth_date FROM users %s ORDER BY %s %s LIMIT $%d OFFSET $%d`,
		whereClause, orderCol, orderDir, argIdx, argIdx+1,
	)
	args = append(args, f.PageSize, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return model.PaginatedResponse{}, fmt.Errorf("fetch query: %w", err)
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Gender, &u.BirthDate); err != nil {
			return model.PaginatedResponse{}, fmt.Errorf("scan: %w", err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return model.PaginatedResponse{}, err
	}

	return model.PaginatedResponse{
		Data:       users,
		TotalCount: totalCount,
		Page:       f.Page,
		PageSize:   f.PageSize,
	}, nil
}

func (r *Repository) GetCommonFriends(userID1, userID2 int) ([]model.User, error) {
	query := `
		SELECT u.id, u.name, u.email, u.gender, u.birth_date
		FROM users u
		JOIN user_friends uf1 ON uf1.friend_id = u.id AND uf1.user_id = $1
		JOIN user_friends uf2 ON uf2.friend_id = u.id AND uf2.user_id = $2
	`
	rows, err := r.db.Query(query, userID1, userID2)
	if err != nil {
		return nil, fmt.Errorf("common friends query: %w", err)
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Gender, &u.BirthDate); err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		users = append(users, u)
	}
	return users, rows.Err()
}
