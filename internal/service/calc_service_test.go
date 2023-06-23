package service

import (
	"party-calc/internal/database/repository"
	"testing"
)

func TestCalcService_calculateOwes(t *testing.T) {
	type fields struct {
		repo    *repository.PersEventsRepository
		service *PersEventsService
		data    *EventData
	}
	type args struct {
		b []PersonBalance
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CalcService{
				repo:    tt.fields.repo,
				service: tt.fields.service,
			}
			s.calculateOwes(tt.args.b)
		})
	}
}
