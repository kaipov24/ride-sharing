package repository

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type inmemRepository struct {
	trips     map[string]*domain.TripModel
	rideFares map[string]*domain.RideFareModel
}

func NewInMemRepository() *inmemRepository {
	return &inmemRepository{
		trips:     make(map[string]*domain.TripModel),
		rideFares: make(map[string]*domain.RideFareModel),
	}
}

func (r *inmemRepository) CreateTrip(ctx context.Context, trip domain.TripModel) (*domain.TripModel, error) {
	if trip.ID.IsZero() {
		trip.ID = primitive.NewObjectID()
	}

	r.trips[trip.ID.Hex()] = &trip
	return &trip, nil
}
