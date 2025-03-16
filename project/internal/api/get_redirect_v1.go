package api

import (
	"context"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/pkg/rules_v1"
)

func (s *Service) GetRedirectV1(ctx context.Context, req *rules_v1.GetRedirectV1Request) (*rules_v1.GetRedirectV1Response, error) {
	result, err := s.rules.FindRedirect(ctx,
		linkFromProto(req.GetLink()),
		requestFromProto(req.GetRequest()),
	)
	if err != nil {
		return nil, err
	}

	return &rules_v1.GetRedirectV1Response{
		Link: linkToProto(result),
	}, nil
}

func requestFromProto(request *rules_v1.RedirectRequestV1) *models.Request {
	return &models.Request{
		Method: request.GetMethod(),
		Header: request.GetHeaders(),
		Body:   request.GetBody(),
	}
}
