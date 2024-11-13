package repository

import (
	"database/sql"
	"tour_destination/model"
)

type EventRepoDB struct {
	db *sql.DB
}

func NewEventRepo(db *sql.DB) *EventRepoDB {
	return &EventRepoDB{db}
}

func (d *EventRepoDB) GetEvent(page, limit int, date, sort string) (*[]model.Events, error) {
	offset := (page - 1) * limit

	query := `
		SELECT e.id, e.price, e.date, d.name, d.description, d.image,
			(SELECT COUNT(t.id) FROM transactions t WHERE t.event_id = e.id) as customer_count,
			ROUND(CAST((SELECT AVG(p.rating) FROM previews p 
					JOIN transactions t ON p.transaction_id = t.id
					JOIN events ev ON t.event_id = ev.id
					WHERE ev.destination_id = d.id) AS numeric), 1) as average_rating
		FROM events e
		JOIN destinations d ON e.destination_id = d.id
		WHERE e.deleted_at IS NULL AND ($1 = '' OR e.date = $1::date)
		ORDER BY e.price ` + sort + `, e.id ASC
		LIMIT $2 OFFSET $3;`

	rows, err := d.db.Query(query, date, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []model.Events{}
	for rows.Next() {
		event := model.Events{}
		if err := rows.Scan(&event.ID, &event.Price, &event.Date, &event.DestinationID.Name, &event.DestinationID.Description, &event.DestinationID.Image, &event.Traveler, &event.Rating); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return &events, nil
}

