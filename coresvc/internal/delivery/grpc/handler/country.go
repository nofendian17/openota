package handler

import (
	"context"
	"errors"
	"github.com/bufbuild/protovalidate-go"
	pb "github.com/nofendian17/openota/coresvc/gen/go/proto/country/v1"
	"github.com/nofendian17/openota/coresvc/internal/entity"
	"github.com/nofendian17/openota/coresvc/internal/usecase/country"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"time"
)

type CountryServiceServer struct {
	countryUseCase country.UseCase
	pb.UnimplementedCountryServiceServer
}

func NewCountryServiceServer(countryUseCase country.UseCase) *CountryServiceServer {
	return &CountryServiceServer{
		countryUseCase: countryUseCase,
	}
}

func (h *CountryServiceServer) GetByID(ctx context.Context, req *pb.GetByIDRequest) (*pb.GetByIDResponse, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result, err := h.countryUseCase.GetByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetByIDResponse{Country: &pb.Country{
		Id:           result.ID.String(),
		Name:         result.Name,
		Code:         result.Code,
		PhoneCode:    result.PhoneCode,
		Capital:      result.Capital,
		Latitude:     result.Latitude,
		Longitude:    result.Longitude,
		CurrencyCode: result.CurrencyCode,
		IsActive:     *result.IsActive,
		Precedence:   result.Precedence,
		CreatedAt:    timestamppb.New(result.CreatedAt),
		UpdatedAt:    timestamppb.New(result.UpdatedAt),
	}}, nil
}
func (h *CountryServiceServer) GetByCode(ctx context.Context, req *pb.GetByCodeRequest) (*pb.GetByCodeResponse, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result, err := h.countryUseCase.GetByCode(ctx, req.GetCode())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetByCodeResponse{Country: &pb.Country{
		Id:           result.ID.String(),
		Name:         result.Name,
		Code:         result.Code,
		PhoneCode:    result.PhoneCode,
		Capital:      result.Capital,
		Latitude:     result.Latitude,
		Longitude:    result.Longitude,
		CurrencyCode: result.CurrencyCode,
		IsActive:     *result.IsActive,
		Precedence:   result.Precedence,
		CreatedAt:    timestamppb.New(result.CreatedAt),
		UpdatedAt:    timestamppb.New(result.UpdatedAt),
	}}, nil
}
func (h *CountryServiceServer) GetAll(ctx context.Context, req *emptypb.Empty) (*pb.GetAllResponse, error) {
	results, err := h.countryUseCase.GetAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	countries := make([]*pb.Country, len(results))

	for i, result := range results {
		countries[i] = &pb.Country{
			Id:           result.ID.String(),
			Name:         result.Name,
			Code:         result.Code,
			PhoneCode:    result.PhoneCode,
			Capital:      result.Capital,
			Latitude:     result.Latitude,
			Longitude:    result.Longitude,
			CurrencyCode: result.CurrencyCode,
			IsActive:     *result.IsActive,
			Precedence:   result.Precedence,
			CreatedAt:    timestamppb.New(result.CreatedAt),
			UpdatedAt:    timestamppb.New(result.UpdatedAt),
		}
	}

	return &pb.GetAllResponse{
		Countries: countries,
	}, nil
}
func (h *CountryServiceServer) Create(ctx context.Context, req *pb.CreateRequest) (*emptypb.Empty, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = h.countryUseCase.Create(ctx, entity.Country{
		Name:         req.GetName(),
		Code:         req.GetCode(),
		PhoneCode:    req.GetPhoneCode(),
		Capital:      req.GetCapital(),
		Latitude:     req.GetLatitude(),
		Longitude:    req.GetLongitude(),
		CurrencyCode: req.GetCurrencyCode(),
		IsActive:     &req.IsActive,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
func (h *CountryServiceServer) Update(ctx context.Context, req *pb.UpdateRequest) (*emptypb.Empty, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	_, err = h.countryUseCase.GetByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = h.countryUseCase.Update(ctx, req.GetId(), entity.Country{
		Name:         req.GetName(),
		Code:         req.GetCode(),
		PhoneCode:    req.GetPhoneCode(),
		Capital:      req.GetCapital(),
		Latitude:     req.GetLatitude(),
		Longitude:    req.GetLongitude(),
		CurrencyCode: req.GetCurrencyCode(),
		IsActive:     &req.IsActive,
		Precedence:   req.GetPrecedence(),
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
func (h *CountryServiceServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*emptypb.Empty, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	_, err = h.countryUseCase.GetByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = h.countryUseCase.Delete(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
