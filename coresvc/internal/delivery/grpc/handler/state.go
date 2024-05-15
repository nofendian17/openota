package handler

import (
	"context"
	"errors"
	"github.com/bufbuild/protovalidate-go"
	"github.com/google/uuid"
	pb "github.com/nofendian17/openota/coresvc/gen/go/proto/state/v1"
	"github.com/nofendian17/openota/coresvc/internal/entity"
	"github.com/nofendian17/openota/coresvc/internal/usecase/state"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"time"
)

type StateServiceServer struct {
	stateUseCase state.UseCase
	pb.UnimplementedStateServiceServer
}

func NewStateServiceServer(stateUseCase state.UseCase) *StateServiceServer {
	return &StateServiceServer{
		stateUseCase: stateUseCase,
	}
}

func (h *StateServiceServer) GetByID(ctx context.Context, req *pb.GetByIDRequest) (*pb.GetByIDResponse, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result, err := h.stateUseCase.GetByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetByIDResponse{State: &pb.State{
		Id:         result.ID.String(),
		Name:       result.Name,
		CountryId:  result.CountryID.String(),
		Latitude:   result.Latitude,
		Longitude:  result.Longitude,
		IsActive:   *result.IsActive,
		Precedence: result.Precedence,
		CreatedAt:  timestamppb.New(result.CreatedAt),
		UpdatedAt:  timestamppb.New(result.UpdatedAt),
	}}, nil
}

func (h *StateServiceServer) GetAll(ctx context.Context, req *emptypb.Empty) (*pb.GetAllResponse, error) {
	results, err := h.stateUseCase.GetAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	states := make([]*pb.State, len(results))

	for i, result := range results {
		states[i] = &pb.State{
			Id:         result.ID.String(),
			Name:       result.Name,
			CountryId:  result.CountryID.String(),
			Latitude:   result.Latitude,
			Longitude:  result.Longitude,
			IsActive:   *result.IsActive,
			Precedence: result.Precedence,
			CreatedAt:  timestamppb.New(result.CreatedAt),
			UpdatedAt:  timestamppb.New(result.UpdatedAt),
		}
	}

	return &pb.GetAllResponse{
		States: states,
	}, nil
}

func (h *StateServiceServer) Create(ctx context.Context, req *pb.CreateRequest) (*emptypb.Empty, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = h.stateUseCase.Create(ctx, entity.State{
		Name:       req.GetName(),
		CountryID:  uuid.MustParse(req.GetCountryId()),
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

func (h *StateServiceServer) Update(ctx context.Context, req *pb.UpdateRequest) (*emptypb.Empty, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	_, err = h.stateUseCase.GetByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = h.stateUseCase.Update(ctx, req.GetId(), entity.State{
		Name:       req.GetName(),
		CountryID:  uuid.MustParse(req.GetCountryId()),
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

func (h *StateServiceServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*emptypb.Empty, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	_, err = h.stateUseCase.GetByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = h.stateUseCase.Delete(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
