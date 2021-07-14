package v1

import (
	"context"

	"github.com/wxbsocial/idgenerator/models"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

type IdGeneratorServer struct {
	generator *models.Generator
}

func NewIdGeneratorServer(generator *models.Generator) *IdGeneratorServer {
	return &IdGeneratorServer{
		generator: generator,
	}
}

func (gen *IdGeneratorServer) Get(
	ctx context.Context,
	empty *emptypb.Empty,

) (*wrapperspb.Int64Value, error) {

	value, err := gen.generator.Get()

	if err != nil {
		return nil, err
	}

	return &wrapperspb.Int64Value{
		Value: value,
	}, nil
}
