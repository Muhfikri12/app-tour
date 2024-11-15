package repository

import (
	"database/sql"
	"time"
	"tour_destination/model"

	"go.uber.org/zap"
)

type EventRepoDB struct {
	db *sql.DB
	Log *zap.Logger
}

func NewEventRepo(db *sql.DB, log *zap.Logger) *EventRepoDB {
	return &EventRepoDB{db, log}
}

func (d *EventRepoDB) GetEvent(page, limit int, date string) (*[]model.Events, int, error) {
	offset := (page - 1) * limit

	query := `
		SELECT e.price, TO_CHAR(e.date, 'yyyy-mm-dd'), d.name, d.description, d.image,
			ROUND(COALESCE((
					SELECT AVG(p.rating)::numeric
					FROM previews p 
					JOIN transactions t ON p.transaction_id = t.id
					JOIN events ev ON t.event_id = ev.id
					WHERE ev.destination_id = d.id 
					AND t.status = true
				), 0), 1) as rating,
			CAST(
				(SELECT COUNT(t.id)
				FROM transactions t
				JOIN events ev ON t.event_id = ev.id
				WHERE ev.destination_id = d.id 
				AND t.status = true) 
			AS INTEGER) AS visitor		
		FROM events e
		JOIN destinations d ON e.destination_id = d.id
		WHERE e.deleted_at IS NULL AND ($1 = '' OR e.date = $1::date)
		LIMIT $2 OFFSET $3;`

	rows, err := d.db.Query(query, date, limit, offset)
	if err != nil {
		d.Log.Error("event repository: failed to fatch events", zap.Error(err))
		return nil,0, err
	}
	defer rows.Close()

	events := []model.Events{}
	for rows.Next() {
		event := model.Events{}
		if err := rows.Scan(&event.Price, &event.Date, &event.DestinationID.Name, &event.DestinationID.Description, &event.DestinationID.Image, &event.Rating, &event.DestinationID.Visitor); err != nil {
			d.Log.Error("event repository:", zap.Error(err))
			return nil,0, err
		}
		events = append(events, event)
	}

	var totalData int

	countQuery := ` SELECT COUNT(*) 
				FROM events e
				JOIN destinations d ON e.destination_id = d.id
				WHERE e.deleted_at IS NULL AND ($1 = '' OR e.date = $1::date);`

	err = d.db.QueryRow(countQuery, date).Scan(&totalData)
	if err != nil {
		d.Log.Error("event repository: failed to fetch total count", zap.Error(err))
		return nil, 0, err
	}

	return &events, totalData, nil
}


func (d *EventRepoDB) GetEventByID(id int) (*model.Events, error) {

	event := model.Events{
		DestinationID: &model.Destination{},
	}

	query := `SELECT 
				e.id, e.price, TO_CHAR(e.date, 'yyyy-mm-dd'), 
				d.id, d.name, d.description, d.image, 
				ROUND(COALESCE((
					SELECT AVG(p.rating)::numeric
					FROM previews p 
					JOIN transactions t ON p.transaction_id = t.id
					JOIN events ev ON t.event_id = ev.id
					WHERE ev.destination_id = d.id 
					AND t.status = true
				), 0), 1) as rating,
				CAST(
					(SELECT COUNT(t.id)
					FROM transactions t
					JOIN events ev ON t.event_id = ev.id
					WHERE ev.destination_id = d.id 
					AND t.status = true) 
				AS INTEGER) as visitor
			FROM events e
			JOIN destinations d ON e.destination_id = d.id
			LEFT JOIN transactions t ON t.event_id = e.id
			WHERE e.id = $1;`

	if err := d.db.QueryRow(query, id).Scan(
			&event.ID, &event.Price, &event.Date,
			&event.DestinationID.ID, &event.DestinationID.Name,
			&event.DestinationID.Description, &event.DestinationID.Image,
			&event.Rating, &event.DestinationID.Visitor); err != nil {
		d.Log.Error("event repository GetEventByID: failed to fetch events", zap.Error(err))
		return nil, err
	}

	photos, err := d.GetPhotosByDestinationID(event.DestinationID.ID)
	if err != nil {
		d.Log.Error("event repository GetEventByID: ", zap.Error(err))
		return nil, err
	}

	event.DestinationID.Photos = photos
	
	return &event, nil
}

func (d *EventRepoDB) GetPhotosByDestinationID(id int) (*[]model.Image, error) {
	images := []model.Image{}

	queryImage := `SELECT image_url, description FROM images WHERE destination_id=$1 AND deleted_at IS NULL`

	rows, err := d.db.Query(queryImage, id)
	if err != nil {
		d.Log.Error("event repository GetPhotosByDestinationID: failed to fatch images", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		image := model.Image{}
		if err := rows.Scan(&image.Image, &image.Description); err != nil {
			d.Log.Error("event repository GetPhotosByDestinationID: ", zap.Error(err))
			return nil, err
		}

		images = append(images, image)
	}

	

	
	return &images, nil
}


func (d *EventRepoDB) CreateTransaction(trx *model.Transaction) error {
	query := `INSERT INTO transactions (name, event_id, email, confirm_email, phone, message, created_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7) 
			  RETURNING id`

	if err := d.db.QueryRow(query, trx.Name, trx.EventID ,trx.Email, trx.EmailConfirm, trx.Phone, trx.Message, trx.CreatedAt).Scan(&trx.ID); err != nil {
		d.Log.Error("event repository: failed to insert transaction", zap.Error(err))
		return err
	}

	today := time.Now()

	trx.CreatedAt = today

	return nil
}


func (d *EventRepoDB) GetEventPlanById(id int) (*[]model.EventPlan, error) {
	
	query := `SELECT ep.first_plan, ep.second_plan, ep.third_plan, ep.fourth_plan, ep.fifth_plan 
			FROM events e
			LEFT JOIN event_plans ep ON ep.event_id = e.id
			WHERE e.id=$1`

	rows, err := d.db.Query(query, id)
	if err != nil {
		d.Log.Error("event repository GetEventPlanById: ", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	tours := []model.EventPlan{}

	for rows.Next() {
		plan := model.EventPlan{}
		if err := rows.Scan(&plan.FirstPlan, &plan.SecondPlan, &plan.ThirdPlan, &plan.FourthPlan, &plan.FifthPlan); err != nil {
			d.Log.Error("event repository GetPhotosByDestinationID: ", zap.Error(err))
			return nil, err
		}

		tours = append(tours, plan)
	}

	return &tours, nil
}

func (d *EventRepoDB) GetEventLocationById(id int) (*[]model.Location, error) {
	
	query := `SELECT d.id, l.first_description, l.coordinate, l.second_description 
			FROM destinations d
			LEFT JOIN locations l ON l.destination_id = d.id
			LEFT JOIN events e ON d.id = e.destination_id 
			WHERE e.id=$1`

	rows, err := d.db.Query(query, id)
	if err != nil {
		d.Log.Error("event repository: ", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	Locations := []model.Location{}

	for rows.Next() {
		location := model.Location{
			Destination: &model.Destination{},
		}
		if err := rows.Scan(&location.Destination.ID, &location.FirstDescription, &location.Coordinate, &location.SecondDescription); err != nil {
			d.Log.Error("event repository: ", zap.Error(err))
			return nil, err
		}

		Locations = append(Locations, location)
	}

	return &Locations, nil
}