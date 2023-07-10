package service

import (
	"context"
	"party-calc/internal/database"
	"party-calc/internal/repository"
	"testing"
	"time"

	"github.com/pashagolub/pgxmock/v2"
	"github.com/stretchr/testify/assert"
)

func TestCalcEvent(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(ctx)

	//repoPer := repository.NewPersonRepository(&database.DataBase{DBPGX: mock})
	repoEv := repository.NewEventRepository(&database.DataBase{DBPGX: mock})
	repoPerEv := repository.NewPersonsEventsRepository(&database.DataBase{DBPGX: mock})

	// Get from events
	mock.ExpectQuery("SELECT (.+) FROM events").
		WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name", "Date"}).
			AddRow(int64(1), "New Year", time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local)))
	mock.ExpectQuery("SELECT (.+) FROM events").
		WithArgs(int64(2)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name", "Date"}).
			AddRow(int64(1), "Old New Year", time.Date(2022, 01, 14, 23, 59, 59, 0, time.Local)))

	mock.ExpectQuery("SELECT person_id FROM persons_events").
		WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRow(int64(1)).AddRow(int64(2)).AddRow(int64(3)))
	mock.ExpectQuery("SELECT person_id FROM persons_events").
		WithArgs(int64(2)).
		WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRow(int64(1)).AddRow(int64(3)).AddRow(int64(4)))

	mock.ExpectQuery("SELECT id, name FROM persons").
		WithArgs([]int64{1}).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name"}).
			AddRows([]any{int64(1), "John Doe"}))

	// Get PersonsEvents

	servCalc := NewCalcService(NewEventService(repoEv), NewPersonsEventsService(repoPerEv))

	result, err := servCalc.CalcEvent(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, nil, result.Balances)
}

func TestEventData_fillAndSortBalances(t *testing.T) {
	tests := []struct {
		name string
		ed   *EventData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ed.fillAndSortBalances()
		})
	}
}
