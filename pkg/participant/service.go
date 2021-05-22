package participant

import (
	"github.com/bagus2x/inamen-go-be/pkg/entity"
	"github.com/bagus2x/inamen-go-be/pkg/model"
	"github.com/bagus2x/inamen-go-be/utils"
)

type service struct {
	repo Repository
}

type Service interface {
	CreateParticipant(req *model.CreateParticipantRequest) (*model.CreateParticipantResponse, error)
	Fetch(id string) (*model.FetchParticipantResponse, error)
	Update(id string, req *model.UpdateParticipantRequest) (*model.UpdateParticipantResponse, error)
	Delete(id string) error
}

func (s *service) CreateParticipant(req *model.CreateParticipantRequest) (*model.CreateParticipantResponse, error) {
	if err := utils.ValidateStruct(req); err != nil {
		return nil, model.ErrvalidationFailed(err)
	}

	players := make([]entity.Player, 0)
	for _, player := range req.Players {
		players = append(players, entity.Player(player))
	}

	participant := entity.Participant{
		TournamentID: req.TournamentID,
		TeamName:     req.TeamName,
		Description:  req.Description,
		Players:      players,
	}

	err := s.repo.Create(&participant)
	if err != nil {
		return nil, err
	}

	res := model.CreateParticipantResponse{
		ID:           participant.TournamentID,
		TournamentID: participant.TournamentID,
		TeamName:     participant.TeamName,
		Description:  participant.Description,
		Players:      req.Players,
		CreatedAt:    participant.CreatedAt,
		UpdatedAt:    participant.UpdatedAt,
	}

	return &res, nil
}

func (s *service) Fetch(id string) (*model.FetchParticipantResponse, error) {
	participant, err := s.repo.Read(id)
	if err != nil {
		return nil, err
	}

	players := make([]model.Player, 0)

	for _, player := range participant.Players {
		players = append(players, model.Player(player))
	}

	res := model.FetchParticipantResponse{
		ID:           participant.ID,
		TournamentID: participant.TournamentID,
		TeamName:     participant.TeamName,
		Description:  participant.Description,
		Players:      players,
		CreatedAt:    participant.CreatedAt,
		UpdatedAt:    participant.UpdatedAt,
	}

	return &res, err
}

func (s *service) FetchAllByTourID(id string) ([]*model.FetchParticipantResponse, error) {
	participants, err := s.repo.ReadAllByTournamentID(id)
	if err != nil {
		return nil, err
	}

	var res = make([]*model.FetchParticipantResponse, 0)

	for _, participant := range participants {
		players := make([]model.Player, 0)

		for _, player := range participant.Players {
			players = append(players, model.Player(player))
		}

		s := model.FetchParticipantResponse{
			ID:           participant.ID,
			TournamentID: participant.TournamentID,
			TeamName:     participant.TeamName,
			Description:  participant.Description,
			Players:      players,
			CreatedAt:    participant.CreatedAt,
			UpdatedAt:    participant.UpdatedAt,
		}

		res = append(res, &s)
	}

	return res, err
}

func (s *service) Update(id string, req *model.UpdateParticipantRequest) (*model.UpdateParticipantResponse, error) {
	if err := utils.ValidateStruct(req); err != nil {
		return nil, model.ErrvalidationFailed(err)
	}

	participant := entity.Participant{
		TeamName:    req.Description,
		Description: req.Description,
	}

	err := s.repo.Update(id, &participant)
	if err != nil {
		return nil, err
	}

	players := make([]model.Player, 0)

	for _, player := range participant.Players {
		players = append(players, model.Player(player))
	}

	res := model.UpdateParticipantResponse{
		ID:           participant.ID,
		TournamentID: participant.TournamentID,
		TeamName:     participant.TeamName,
		Description:  participant.Description,
		Players:      players,
		CreatedAt:    participant.CreatedAt,
		UpdatedAt:    participant.UpdatedAt,
	}

	return &res, nil
}

func (s *service) Delete(id string) error {
	return s.repo.Delete(id)
}
