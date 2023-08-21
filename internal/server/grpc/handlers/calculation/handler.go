package calculation

import (
	pb "party-calc/internal/server/grpc/proto"
	"party-calc/internal/service"
)

type CalcHandler struct {
	pb.CalculationServer
	service *service.CalcService
}

func NewCalcHandler(s *service.CalcService) *CalcHandler {
	return &CalcHandler{
		service: s,
	}
}
