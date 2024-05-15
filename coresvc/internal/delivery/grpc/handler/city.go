package handler

import (
	"context"
	"errors"
	"github.com/bufbuild/protovalidate-go"
	"github.com/google/uuid"
	pb "github.com/nofendian17/openota/coresvc/gen/go/proto/city/v1"
	"github.com/nofendian17/openota/coresvc/internal/entity"
	"github.com/nofendian17/openota/coresvc/internal/usecase/city"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"time"
)

type CityServiceServer struct {
	cityUseCase city.UseCase
	pb.UnimplementedCityServiceServer
}

func NewCityServiceServer(cityUseCase city.UseCase) *CityServiceServer {
	return &CityServiceServer{
		cityUseCase: cityUseCase,
	}
}

func (h *CityServiceServer) GetByID(ctx context.Context, req *pb.GetByIDRequest) (*pb.GetByIDResponse, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result, err := h.cityUseCase.GetByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetByIDResponse{City: &pb.City{
		Id:         result.ID.String(),
		Name:       result.Name,
		StateId:    result.StateID.String(),
		Latitude:   result.Latitude,
		Longitude:  result.Longitude,
		IsActive:   *result.IsActive,
		Precedence: result.Precedence,
		CreatedAt:  timestamppb.New(result.CreatedAt),
		UpdatedAt:  timestamppb.New(result.UpdatedAt),
	}}, nil
}

func (h *CityServiceServer) GetAll(ctx context.Context, req *emptypb.Empty) (*pb.GetAllResponse, error) {
	results, err := h.cityUseCase.GetAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	cities := make([]*pb.City, len(results))

	for i, result := range results {
		cities[i] = &pb.City{
			Id:         result.ID.String(),
			Name:       result.Name,
			StateId:    result.StateID.String(),
			Latitude:   result.Latitude,
			Longitude:  result.Longitude,
			IsActive:   *result.IsActive,
			Precedence: result.Precedence,
			CreatedAt:  timestamppb.New(result.CreatedAt),
			UpdatedAt:  timestamppb.New(result.UpdatedAt),
		}
	}

	return &pb.GetAllResponse{
		Cities: cities,
	}, nil
}

func (h *CityServiceServer) Create(ctx context.Context, req *pb.CreateRequest) (*emptypb.Empty, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = h.cityUseCase.Create(ctx, entity.City{
		Name:       req.GetName(),
		StateID:    uuid.MustParse(req.GetStateId()),
		Latitude:   req.GetLatitude(),
		Longitude:  req.GetLongitude(),
		IsActive:   &req.IsActive,
		Precedence: 0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (h *CityServiceServer) Update(ctx context.Context, req *pb.UpdateRequest) (*emptypb.Empty, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	_, err = h.cityUseCase.GetByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = h.cityUseCase.Update(ctx, req.GetId(), entity.City{
		Name:       req.GetName(),
		StateID:    uuid.MustParse(req.GetStateId()),
		Latitude:   req.GetLatitude(),
		Longitude:  req.GetLongitude(),
		IsActive:   &req.IsActive,
		Precedence: req.GetPrecedence(),
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Now(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (h *CityServiceServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*emptypb.Empty, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	_, err = h.cityUseCase.GetByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = h.cityUseCase.Delete(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
