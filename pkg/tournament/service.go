package tournament

import (
	"github.com/bagus2x/inamen-go-be/pkg/entity"
	"github.com/bagus2x/inamen-go-be/pkg/model"
)

type service struct {
	repo Repository
}

type Service interface {
	Create(hostID string, hostUsername string, req *model.CreateTourRequest) (*model.CreateTourResponse, error)
	Fetch(id string) (*model.FetchTourResponse, error)
	FetchToursByHostID(hostID string) (model.FetchToursResponse, error)
	Update(id, hostID string, req *model.UpdateTourRequest) (*model.UpdateTourResponse, error)
	Delete(id, hostID string) error
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(hostID string, hostUsername string, req *model.CreateTourRequest) (*model.CreateTourResponse, error) {
	tour := entity.Tournament{
		Name:        req.Name,
		Description: req.Description,
		Host: entity.Host{
			ID:       hostID,
			Username: hostUsername,
		},
		Game:              req.Game,
		Mode:              req.Mode,
		Presence:          req.Presence,
		Visibilty:         req.Visibilty,
		TotalParticipants: req.TotalParticipants,
		TotalTeamMembers:  req.TotalTeamMembers,
		Platforms:         req.Platforms,
		Date:              entity.TournamentDate(req.Date),
		Published:         false,
	}

	err := s.repo.Create(&tour)

	res := model.CreateTourResponse{
		ID:                tour.ID,
		Description:       tour.Description,
		Name:              tour.Name,
		Host:              model.TourHost(tour.Host),
		Game:              tour.Game,
		Mode:              tour.Mode,
		Presence:          tour.Presence,
		Visibilty:         tour.Visibilty,
		TotalParticipants: tour.TotalParticipants,
		TotalTeamMembers:  tour.TotalTeamMembers,
		Platforms:         tour.Platforms,
		Date:              model.TourDate(tour.Date),
		Published:         false,
		CreatedAt:         tour.CreatedAt,
		UpdatedAt:         tour.UpdatedAt,
	}

	return &res, err
}

func (s *service) Fetch(id string) (*model.FetchTourResponse, error) {
	tour, err := s.repo.Read(id)
	if err != nil {
		return nil, err
	}

	res := model.FetchTourResponse{
		ID:                tour.ID,
		Name:              tour.Name,
		Host:              model.TourHost(tour.Host),
		Game:              tour.Game,
		Mode:              tour.Mode,
		Presence:          tour.Presence,
		Published:         tour.Published,
		Visibilty:         tour.Visibilty,
		TotalParticipants: tour.TotalParticipants,
		TotalTeamMembers:  tour.TotalTeamMembers,
		Platforms:         tour.Platforms,
		Date:              model.TourDate(tour.Date),
		CreatedAt:         tour.CreatedAt,
		UpdatedAt:         tour.UpdatedAt,
	}

	return &res, nil
}

func (s *service) FetchToursByHostID(hostID string) (model.FetchToursResponse, error) {
	tours, err := s.repo.ReadAllByHost(hostID)
	if err != nil {
		return nil, err
	}

	var res model.FetchToursResponse

	for _, tour := range tours {
		t := model.FetchTourResponse{
			ID:                tour.ID,
			Name:              tour.Name,
			Host:              model.TourHost(tour.Host),
			Game:              tour.Game,
			Mode:              tour.Mode,
			Presence:          tour.Presence,
			Published:         tour.Published,
			Visibilty:         tour.Visibilty,
			TotalParticipants: tour.TotalParticipants,
			TotalTeamMembers:  tour.TotalTeamMembers,
			Platforms:         tour.Platforms,
			Date:              model.TourDate(tour.Date),
			CreatedAt:         tour.CreatedAt,
			UpdatedAt:         tour.UpdatedAt,
		}

		res = append(res, &t)
	}

	return res, nil
}

func (s *service) Update(id, hostID string, req *model.UpdateTourRequest) (*model.UpdateTourResponse, error) {
	tour := entity.Tournament{
		Name:        req.Name,
		Description: req.Name,
		Published:   req.Published,
		Visibilty:   req.Visibilty,
		Date:        entity.TournamentDate(req.Date),
	}

	err := s.repo.Update(id, hostID, &tour)

	res := model.UpdateTourResponse{
		ID:                tour.ID,
		Description:       tour.Description,
		Name:              tour.Name,
		Host:              model.TourHost(tour.Host),
		Game:              tour.Game,
		Mode:              tour.Mode,
		Presence:          tour.Presence,
		Visibilty:         tour.Visibilty,
		TotalParticipants: tour.TotalParticipants,
		TotalTeamMembers:  tour.TotalTeamMembers,
		Platforms:         tour.Platforms,
		Date:              model.TourDate(tour.Date),
		Published:         false,
		CreatedAt:         tour.CreatedAt,
		UpdatedAt:         tour.UpdatedAt,
	}

	return &res, err
}

func (s *service) Delete(id, hostID string) error {
	return s.repo.Delete(id, hostID)
}
