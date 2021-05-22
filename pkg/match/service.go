package match

import (
	"github.com/bagus2x/inamen-go-be/pkg/entity"
	"github.com/bagus2x/inamen-go-be/pkg/model"
)

type service struct {
	repo Repository
}

type Service interface {
	CreateMatch(req *model.CreateMatchRequest) (*model.CreateMatchResponse, error)
	Fetch(id string) (*model.FetchMatchResponse, error)
	Update(id string, req *model.UpdateMatchRequest) (*model.UpdateMatchResponse, error)
	Delete(id string) error
}

func (s *service) CreateMatch(req *model.CreateMatchRequest) (*model.CreateMatchResponse, error) {
	match := entity.Match{
		TournamentID:   req.TournamentID,
		Name:           req.Name,
		Dates:          entity.MatchDates(req.Dates),
		ParticipantIDs: req.ParticipantIDs,
	}

	err := s.repo.Create(&match)
	if err != nil {
		return nil, err
	}

	res := model.CreateMatchResponse{
		ID:             match.ID,
		TournamentID:   match.TournamentID,
		Name:           match.Name,
		Dates:          model.MatchDates(match.Dates),
		ParticipantIDs: match.ParticipantIDs,
		CreatedAt:      match.CreatedAt,
		UpdatedAt:      match.UpdatedAt,
	}

	return &res, nil
}

func (s *service) Fetch(id string) (*model.FetchMatchResponse, error) {
	match, err := s.repo.Read(id)
	if err != nil {
		return nil, err
	}

	res := model.FetchMatchResponse{
		ID:             match.ID,
		TournamentID:   match.TournamentID,
		Name:           match.Name,
		Dates:          model.MatchDates(match.Dates),
		ParticipantIDs: match.ParticipantIDs,
		CreatedAt:      match.CreatedAt,
		UpdatedAt:      match.UpdatedAt,
	}

	return &res, err
}

func (s *service) FetchAllByTourID(id string) ([]*model.FetchMatchResponse, error) {
	matchs, err := s.repo.ReadAllByTournamentID(id)
	if err != nil {
		return nil, err
	}

	var res = make([]*model.FetchMatchResponse, 0)

	for _, match := range matchs {
		s := model.FetchMatchResponse{
			ID:             match.ID,
			TournamentID:   match.TournamentID,
			Name:           match.Name,
			Dates:          model.MatchDates(match.Dates),
			ParticipantIDs: match.ParticipantIDs,
			CreatedAt:      match.CreatedAt,
			UpdatedAt:      match.UpdatedAt,
		}

		res = append(res, &s)
	}

	return res, err
}

func (s *service) Update(id string, req *model.UpdateMatchRequest) (*model.UpdateMatchResponse, error) {
	match := entity.Match{

		Name:        req.Name,
		Description: req.Description,
		Dates:       entity.MatchDates(req.Dates),
	}

	err := s.repo.Update(id, &match)
	if err != nil {
		return nil, err
	}

	res := model.UpdateMatchResponse{
		ID:             match.ID,
		TournamentID:   match.TournamentID,
		Name:           match.Name,
		Dates:          model.MatchDates(match.Dates),
		ParticipantIDs: match.ParticipantIDs,
		CreatedAt:      match.CreatedAt,
		UpdatedAt:      match.UpdatedAt,
	}

	return &res, nil
}

func (s *service) Delete(id string) error {
	return s.repo.Delete(id)
}
