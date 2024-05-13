package handler

import (
	"context"
	"github.com/bufbuild/protovalidate-go"
	pb "github.com/nofendian17/openota/coresvc/gen/go/proto/airline/v1"
	"github.com/nofendian17/openota/coresvc/internal/entity"
	"github.com/nofendian17/openota/coresvc/internal/usecase/airline"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
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
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result, err := h.airlineUseCase.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetByIDResponse{Airline: &pb.Airline{
		Id:         result.ID.String(),
		Name:       result.Name,
		Code:       result.Code,
		Logo:       result.Logo,
		IsActive:   *result.IsActive,
		Precedence: result.Precedence,
		CreatedAt:  timestamppb.New(result.CreatedAt),
		UpdatedAt:  timestamppb.New(result.UpdatedAt),
	}}, nil
}
func (h *AirlineServiceServer) GetByCode(ctx context.Context, req *pb.GetByCodeRequest) (*pb.GetByCodeResponse, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result, err := h.airlineUseCase.GetByCode(ctx, req.GetCode())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetByCodeResponse{Airline: &pb.Airline{
		Id:         result.ID.String(),
		Name:       result.Name,
		Code:       result.Code,
		Logo:       result.Logo,
		IsActive:   *result.IsActive,
		Precedence: result.Precedence,
		CreatedAt:  timestamppb.New(result.CreatedAt),
		UpdatedAt:  timestamppb.New(result.UpdatedAt),
	}}, nil
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
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = h.airlineUseCase.Create(ctx, entity.Airline{
		Name:     req.GetName(),
		Code:     req.GetCode(),
		Logo:     req.GetLogo(),
		IsActive: &req.IsActive,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
func (h *AirlineServiceServer) Update(ctx context.Context, req *pb.UpdateRequest) (*emptypb.Empty, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = h.airlineUseCase.Update(ctx, req.GetId(), entity.Airline{
		Name:       req.GetName(),
		Code:       req.GetCode(),
		Logo:       req.GetLogo(),
		IsActive:   &req.IsActive,
		Precedence: req.GetPrecedence(),
		UpdatedAt:  time.Now(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
func (h *AirlineServiceServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*emptypb.Empty, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	err = h.airlineUseCase.Delete(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
