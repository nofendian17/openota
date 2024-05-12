package handler

import (
	"context"
	"github.com/bufbuild/protovalidate-go"
	pb "github.com/nofendian17/openota/coresvc/gen/go/proto/airline/v1"
	"github.com/nofendian17/openota/coresvc/internal/usecase/airline"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AirlineServiceServer struct {
	airlineUseCase airline.UseCase
	pb.UnimplementedAirlineServiceServer
}

func NewAirlineServiceServer(airlineUseCase airline.UseCase) *AirlineServiceServer {
	return &AirlineServiceServer{
		airlineUseCase: airlineUseCase,
	}
}

func (h *AirlineServiceServer) GetByID(ctx context.Context, req *pb.GetByIDRequest) (*pb.GetByIDResponse, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (h *AirlineServiceServer) GetByCode(ctx context.Context, req *pb.GetByCodeRequest) (*pb.GetByCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByCode not implemented")
}
func (h *AirlineServiceServer) GetAll(ctx context.Context, req *emptypb.Empty) (*pb.GetAllResponse, error) {
	results, err := h.airlineUseCase.GetAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	airlines := make([]*pb.Airline, len(results))

	for i, result := range results {
		airlines[i] = &pb.Airline{
			Id:         result.ID.String(),
			Name:       result.Name,
			Code:       result.Code,
			Logo:       result.Logo,
			IsActive:   *result.IsActive,
			Precedence: result.Precedence,
			CreatedAt:  timestamppb.New(result.CreatedAt),
			UpdatedAt:  timestamppb.New(result.UpdatedAt),
		}
	}

	return &pb.GetAllResponse{
		Airlines: airlines,
	}, nil
}
func (h *AirlineServiceServer) Create(ctx context.Context, req *pb.CreateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (h *AirlineServiceServer) Update(ctx context.Context, req *pb.UpdateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (h *AirlineServiceServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
