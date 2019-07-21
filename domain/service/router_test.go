package service_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/rcdmk/shortest-flight-path/domain"

	"github.com/rcdmk/shortest-flight-path/data/datamock"
	"github.com/rcdmk/shortest-flight-path/domain/entity"
	"github.com/rcdmk/shortest-flight-path/domain/service"
)

func setupMockAirportRepo(mockDM *datamock.DataManager) {
	// mock airports
	var knownAirports = map[string]bool{
		"GRU": true,
		"LIM": true,
		"MIA": true,
		"PUN": true,
	}

	mockAirportRepo := mockDM.Airports().(*datamock.AirportRepo)

	// Unknown airports return not found error
	mockAirportRepo.On("GetByCode",
		mock.MatchedBy(func(iata3 string) bool {
			return !knownAirports[iata3]
		})).Return(entity.Airport{}, domain.ErrNotFound)

	// Known airports return not found error
	mockAirportRepo.On("GetByCode",
		mock.MatchedBy(func(iata3 string) bool {
			return knownAirports[iata3]
		})).Return(entity.Airport{}, nil)

	return
}

func setupMockRouteRepo(mockDM *datamock.DataManager) {
	// mock routes
	mockRouteRepo := mockDM.Routes().(*datamock.RouteRepo)

	// Known airports
	routes := []entity.Route{
		entity.Route{
			Origin:      "GRU",
			Destination: "LIM",
			AirlineCode: "LT",
		},
		entity.Route{
			Origin:      "GRU",
			Destination: "MIA",
			AirlineCode: "AA",
		},
	}

	mockRouteRepo.On("GetAllDepartingFromAirport", "GRU").Return(routes, nil)

	routes = []entity.Route{
		entity.Route{
			Origin:      "LIM",
			Destination: "GRU",
			AirlineCode: "LT",
		},
		entity.Route{
			Origin:      "LIM",
			Destination: "PUN",
			AirlineCode: "LT",
		},
	}

	mockRouteRepo.On("GetAllDepartingFromAirport", "LIM").Return(routes, nil)

	routes = []entity.Route{
		entity.Route{
			Origin:      "PUN",
			Destination: "LIM",
			AirlineCode: "LT",
		},
	}

	mockRouteRepo.On("GetAllDepartingFromAirport", "PUN").Return(routes, nil)

	return
}

func setupMockDataManager() *datamock.DataManager {
	mockDM := datamock.New()

	setupMockAirportRepo(mockDM)
	setupMockRouteRepo(mockDM)

	return mockDM
}

func Test_router_GetShortestRoute(t *testing.T) {
	mockDM := setupMockDataManager()

	var r = service.NewRouter(mockDM)

	tests := []struct {
		name        string
		source      string
		destination string
		wantStops   []entity.Route
		wantErr     error
	}{
		{
			name:        "Should return not found error when an non existing source is provided",
			source:      "XXX",
			destination: "GRU",
			wantStops:   nil,
			wantErr:     domain.ErrInvalidRouteOrigin,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotStops, gotErr := r.GetShortestRoute(test.source, test.destination)
			if (test.wantErr == nil && gotErr != nil) ||
				(test.wantErr != nil && gotErr == nil) ||
				(test.wantErr.Error() != gotErr.Error()) {
				t.Errorf("error = %v, want err = %v", gotErr, test.wantErr)
				return
			}

			if !reflect.DeepEqual(gotStops, test.wantStops) {
				t.Errorf("got = %v, want = %v", gotStops, test.wantStops)
			}
		})
	}
}
