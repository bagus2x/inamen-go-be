package standing

import (
	"github.com/bagus2x/inamen-go-be/pkg/entity"
	"github.com/bagus2x/inamen-go-be/pkg/model"
)

type service struct {
	repo Repository
}

type Service interface {
	CreateStanding(req *model.CreateStandingRequest) (*model.CreateStandingResponse, error)
	Fetch(id string) (*model.FetchStandingResponse, error)
	Update(id string, req *model.UpdateStandingRequest) (*model.UpdateStandingResponse, error)
	Delete(id string) error
}

func (s *service) CreateStanding(req *model.CreateStandingRequest) (*model.CreateStandingResponse, error) {
	standing := entity.Standing{
		TournamentID: req.TournamentID,
		Name:         req.Name,
		Spotlight:    req.Spotlight,
		Schema:       entity.StandingSchema(req.Schema),
	}

	err := s.repo.Create(&standing)
	if err != nil {
		return nil, err
	}

	res := model.CreateStandingResponse{
		ID:           standing.ID,
		TournamentID: standing.TournamentID,
		Name:         standing.Name,
		Spotlight:    standing.Spotlight,
		Schema:       model.StandingSchema(standing.Schema),
		CreatedAt:    standing.CreatedAt,
		UpdatedAt:    standing.UpdatedAt,
	}

	return &res, nil
}

func (s *service) Fetch(id string) (*model.FetchStandingResponse, error) {
	standing, err := s.repo.Read(id)
	if err != nil {
		return nil, err
	}

	res := model.FetchStandingResponse{
		ID:           standing.ID,
		TournamentID: standing.TournamentID,
		Name:         standing.Name,
		Spotlight:    standing.Spotlight,
		Schema:       model.StandingSchema(standing.Schema),
		CreatedAt:    standing.CreatedAt,
		UpdatedAt:    standing.UpdatedAt,
	}

	return &res, err
}

func (s *service) FetchAllByTourID(id string) ([]*model.FetchStandingResponse, error) {
	standings, err := s.repo.ReadAllByTournamentID(id)
	if err != nil {
		return nil, err
	}

	var res = make([]*model.FetchStandingResponse, 0)

	for _, standing := range standings {
		s := model.FetchStandingResponse{
			ID:           standing.ID,
			TournamentID: standing.TournamentID,
			Name:         standing.Name,
			Spotlight:    standing.Spotlight,
			Schema:       model.StandingSchema(standing.Schema),
			CreatedAt:    standing.CreatedAt,
			UpdatedAt:    standing.UpdatedAt,
		}

		res = append(res, &s)
	}

	return res, err
}

func (s *service) Update(id string, req *model.UpdateStandingRequest) (*model.UpdateStandingResponse, error) {
	standing := entity.Standing{
		Name:      req.Name,
		Spotlight: req.Spotlight,
		Schema:    entity.StandingSchema(req.Schema),
	}

	err := s.repo.Update(id, &standing)
	if err != nil {
		return nil, err
	}

	res := model.UpdateStandingResponse{
		ID:           standing.ID,
		TournamentID: standing.TournamentID,
		Name:         standing.Name,
		Spotlight:    standing.Spotlight,
		Schema:       model.StandingSchema(standing.Schema),
		CreatedAt:    standing.CreatedAt,
		UpdatedAt:    standing.UpdatedAt,
	}

	return &res, nil
}

func (s *service) Delete(id string) error {
	return s.repo.Delete(id)
}
