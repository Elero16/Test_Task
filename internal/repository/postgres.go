package repository

import (
    "database/sql"
    "github.com/google/uuid"
    "github.com/your-username/project/internal/models"
    _ "github.com/lib/pq"
)

type Repository struct {
    db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
    return &Repository{db: db}
}

func (r *Repository) Create(sub models.Subscription) error {
    query := `INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date) 
              VALUES ($1, $2, $3, $4, $5)`
    _, err := r.db.Exec(query, sub.ServiceName, sub.Price, sub.UserID, sub.StartDate, sub.EndDate)
    return err
}

func (r *Repository) List() ([]models.Subscription, error) {
    rows, err := r.db.Query("SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var subs []models.Subscription
    for rows.Next() {
        var s models.Subscription
        if err := rows.Scan(&s.ID, &s.ServiceName, &s.Price, &s.UserID, &s.StartDate, &s.EndDate); err != nil {
            return nil, err
        }
        subs = append(subs, s)
    }
    return subs, nil
}

func (r *Repository) Delete(id string) error {
    _, err := r.db.Exec("DELETE FROM subscriptions WHERE id = $1", id)
    return err
}

func (r *Repository) GetTotalCost(userID uuid.UUID, serviceName string) (int, error) {
    var total int
    query := `SELECT COALESCE(SUM(price), 0) FROM subscriptions WHERE user_id = $1`
    args := []interface{}{userID}
    if serviceName != "" {
        query += " AND service_name = $2"
        args = append(args, serviceName)
    }
    err := r.db.QueryRow(query, args...).Scan(&total)
    return total, err
}