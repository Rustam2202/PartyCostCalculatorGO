package calculation

import (
	"context"
	"party-calc/internal/server/grpc/proto"
)

func (h *CalcHandler) GetCalculatedEvent(ctx context.Context, pb *proto.CalculatedEventGet) (*proto.EventData, error) {
	eventData, err := h.service.CalculateEvent(ctx, pb.EventId, pb.RoundRate)
	if err != nil {
		return nil, err
	}
	result := proto.EventData{
		EventName:    eventData.Name,
		Date:         eventData.Date.Format("2006-01-02"),
		AverageSpent: float32(eventData.AverageSpent.InexactFloat64()),
		TotalSpent:   float32(eventData.TotalSpent.InexactFloat64()),
		PersonsCount: eventData.AllPeronsCount,
		RoundRate:    eventData.RoundRate,
		Debetors:     make(map[string]*proto.Recepients),
	}
	for debetor, recepients := range eventData.Owes {
		r := proto.Recepients{}
		for recepient, debt := range recepients {
			r.Recepients[recepient] = float32(debt.Ceil().InexactFloat64())
		}
		result.Debetors[debetor].Recepients = r.Recepients
	}
	return &result, nil
}
