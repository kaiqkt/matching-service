package repositories

import (
	"database/sql"
	"matching-service/internal/models"
)

type RideRepository struct {
	db *sql.DB
}

func NewRideRepository(db *sql.DB) *RideRepository {
	return &RideRepository{db}
}

func (r *RideRepository) Save(ride *models.Ride) error {
	stm, err := r.db.Prepare("INSERT INTO ride (id, user_id, source_latitude, source_longitude, destination_latitude, destination_longitude, status, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		return err
	}
	defer stm.Close()

	_, err = stm.Exec(ride.ID.String(), ride.UserID, ride.Source.Latitude, ride.Source.Longitude, ride.Destination.Latitude, ride.Destination.Longitude, ride.Status, ride.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *RideRepository) ExistsByUserID(userID *string) (bool, error) {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM ride WHERE user_id = $1 AND status != $2)", userID, models.CANCELED).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// func (r *RideRepository) FindByUserID(userID *string) (*models.Ride, error) {
// 	var ride models.Ride

// 	err := r.db.QueryRow("SELECT * FROM ride WHERE person_id = ?", userID).Scan(&ride.ID, &ride.UserID, &ride.Source.Latitude, &ride.Source.Longitude, &ride.Destination.Latitude, &ride.Destination.Longitude, &ride.Status, &ride.CreatedAt)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, &cerrors.CustomErr{Message: fmt.Sprintf("No ride found for UserID %w", userID), HttpStatus: http.Not}
// 		}
// 		return nil, err
// 	}

// 	return &ride, nil
// }
