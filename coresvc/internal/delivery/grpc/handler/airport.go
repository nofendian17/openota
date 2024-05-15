package handler

import (
	"context"
	"errors"
	"github.com/bufbuild/protovalidate-go"
	"github.com/google/uuid"
	pb "github.com/nofendian17/openota/coresvc/gen/go/proto/airport/v1"
	"github.com/nofendian17/openota/coresvc/internal/entity"
	"github.com/nofendian17/openota/coresvc/internal/usecase/airport"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"time"
)

type AirportServiceServer struct {
	airportUseCase airport.UseCase
	pb.UnimplementedAirportServiceServer
}

func NewAirportServiceServer(airportUseCase airport.UseCase) *AirportServiceServer {
	return &AirportServiceServer{
		airportUseCase: airportUseCase,
	}
}

func (h *AirportServiceServer) GetByID(ctx context.Context, req *pb.GetByIDRequest) (*pb.GetByIDResponse, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result, err := h.airportUseCase.GetByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetByIDResponse{Airport: &pb.Airport{
		Id:         result.ID.String(),
		Code:       result.Code,
		Name:       result.Name,
		CityId:     result.CityID.String(),
		Latitude:   result.Latitude,
		Longitude:  result.Longitude,
		IsDomestic: *result.IsDomestic,
		Timezone:   result.Timezone,
		IsActive:   *result.IsActive,
		Precedence: result.Precedence,
		CreatedAt:  timestamppb.New(result.CreatedAt),
		UpdatedAt:  timestamppb.New(result.UpdatedAt),
	}}, nil
}

func (h *AirportServiceServer) GetByCode(ctx context.Context, req *pb.GetByCodeRequest) (*pb.GetByCodeResponse, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result, err := h.airportUseCase.GetByCode(ctx, req.GetCode())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetByCodeResponse{Airport: &pb.Airport{
		Id:         result.ID.String(),
		Code:       result.Code,
		Name:       result.Name,
		CityId:     result.CityID.String(),
		Latitude:   result.Latitude,
		Longitude:  result.Longitude,
		IsDomestic: *result.IsDomestic,
		Timezone:   result.Timezone,
		IsActive:   *result.IsActive,
		Precedence: result.Precedence,
		CreatedAt:  timestamppb.New(result.CreatedAt),
		UpdatedAt:  timestamppb.New(result.UpdatedAt),
	}}, nil
}
func (h *AirportServiceServer) GetAll(ctx context.Context, req *emptypb.Empty) (*pb.GetAllResponse, error) {
	results, err := h.airportUseCase.GetAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	airports := make([]*pb.Airport, len(results))

	for i, result := range results {
		airports[i] = &pb.Airport{
			Id:         result.ID.String(),
			Code:       result.Code,
			Name:       result.Name,
			CityId:     result.CityID.String(),
			Latitude:   result.Latitude,
			Longitude:  result.Longitude,
			IsDomestic: *result.IsDomestic,
			Timezone:   result.Timezone,
			IsActive:   *result.IsActive,
			Precedence: result.Precedence,
			CreatedAt:  timestamppb.New(result.CreatedAt),
			UpdatedAt:  timestamppb.New(result.UpdatedAt),
		}
	}

	return &pb.GetAllResponse{
		Airports: airports,
	}, nil
}

func (h *AirportServiceServer) Create(ctx context.Context, req *pb.CreateRequest) (*emptypb.Empty, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = h.airportUseCase.Create(ctx, entity.Airport{
		Code:       req.GetCode(),
		Name:       req.GetName(),
		CityID:     uuid.MustParse(req.CityId),
		Latitude:   req.GetLatitude(),
		Longitude:  req.GetLongitude(),
		IsDomestic: &req.IsDomestic,
		Timezone:   req.GetTimezone(),
		IsActive:   &req.IsActive,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (h *AirportServiceServer) Update(ctx context.Context, req *pb.UpdateRequest) (*emptypb.Empty, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	_, err = h.airportUseCase.GetByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = h.airportUseCase.Update(ctx, req.GetId(), entity.Airport{
		Code:       req.GetCode(),
		Name:       req.GetName(),
		CityID:     uuid.MustParse(req.GetCityId()),
		Latitude:   req.GetLatitude(),
		Longitude:  req.GetLongitude(),
		IsDomestic: &req.IsDomestic,
		Timezone:   req.GetTimezone(),
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

func (h *AirportServiceServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*emptypb.Empty, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	_, err = h.airportUseCase.GetByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = h.airportUseCase.Delete(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
