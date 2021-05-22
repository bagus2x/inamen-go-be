package model

import (
	"errors"
	"fmt"
)

var (
	ErrInternalServer            = errors.New("Internal Server Error")
	ErrNotFound                  = errors.New("Not Found")
	ErrConflict                  = errors.New("Conflict")
	ErrBadRequest                = errors.New("Bad request")
	ErrUnauthorized              = errors.New("Unauthorized")
	ErrInvalidIDFormat           = fmt.Errorf("%w: Invalid ID format", ErrBadRequest)
	ErrFailedToCreateTournament  = fmt.Errorf("%w: Failed to create tournament", ErrInternalServer)
	ErrFailedToReadTournament    = fmt.Errorf("%w: Failed to Read tournament", ErrInternalServer)
	ErrFailedToUpdateTournament  = fmt.Errorf("%w: Failed to Update tournament", ErrInternalServer)
	ErrFailedToDeleteTournament  = fmt.Errorf("%w: Failed to Delete tournament", ErrInternalServer)
	ErrTournamentNotfound        = fmt.Errorf("%w: Tournament does not exists", ErrNotFound)
	ErrFailedToCreateStanding    = fmt.Errorf("%w: Failed to create standing", ErrInternalServer)
	ErrFailedToReadStanding      = fmt.Errorf("%w: Failed to Read standing", ErrInternalServer)
	ErrFailedToUpdateStanding    = fmt.Errorf("%w: Failed to Update standing", ErrInternalServer)
	ErrFailedToDeleteStanding    = fmt.Errorf("%w: Failed to Delete standing", ErrInternalServer)
	ErrStandingNotfound          = fmt.Errorf("%w: Standing does not exist", ErrNotFound)
	ErrFailedToCreateParticipant = fmt.Errorf("%w: Failed to create participant", ErrInternalServer)
	ErrFailedToReadParticipant   = fmt.Errorf("%w: Failed to Read participant", ErrInternalServer)
	ErrFailedToUpdateParticipant = fmt.Errorf("%w: Failed to Update participant", ErrInternalServer)
	ErrFailedToDeleteParticipant = fmt.Errorf("%w: Failed to Delete participant", ErrInternalServer)
	ErrParticipantNotfound       = fmt.Errorf("%w: Participant does not exists", ErrNotFound)
	ErrFailedToCreateMatch       = fmt.Errorf("%w: Failed to create participant", ErrInternalServer)
	ErrFailedToReadMatch         = fmt.Errorf("%w: Failed to Read match", ErrInternalServer)
	ErrFailedToUpdateMatch       = fmt.Errorf("%w: Failed to Update match", ErrInternalServer)
	ErrFailedToDeleteMatch       = fmt.Errorf("%w: Failed to Delete match", ErrInternalServer)
	ErrMatchNotfound             = fmt.Errorf("%w: Match does not exist", ErrNotFound)
	ErrInvalidAccessToken        = fmt.Errorf("%w: Invalid Acccess Token", ErrUnauthorized)
	ErrTokenExpired              = fmt.Errorf("%w: Token expired", ErrUnauthorized)
	ErrvalidationFailed          = func(err error) error { return fmt.Errorf("%w: %s", ErrBadRequest, err) }
)
