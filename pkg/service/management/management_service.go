package management_service

import (
	grpc_client "github.com/Sensory-Cloud/go-sdk/pkg/client"
	"github.com/Sensory-Cloud/go-sdk/pkg/config"
	management_api_v1 "github.com/Sensory-Cloud/go-sdk/pkg/generated/v1/management"
	token_manager "github.com/Sensory-Cloud/go-sdk/pkg/token"
)

type ManagementService struct {
	config           *config.ClientConfig
	tokenManager     token_manager.IAuthorizationMetadataService
	enrollmentClient management_api_v1.EnrollmentServiceClient
}

// NewManagementService creates a service to handle audio requests to Sensory Cloud
func NewManagementService(config *config.ClientConfig, tokenManager token_manager.IAuthorizationMetadataService) (*ManagementService, error) {
	client, err := config.GetClient()
	if err != nil {
		return nil, err
	}

	enrollmentClient := management_api_v1.NewEnrollmentServiceClient(client)
	return &ManagementService{config, tokenManager, enrollmentClient}, nil
}

// Obtains all of the active enrollments given the userId
func (s *ManagementService) GetEnrollments(userId string) (*management_api_v1.GetEnrollmentsResponse, error) {
	ctx, cancel := grpc_client.GetDefaultContext()
	defer cancel()

	ctx, err := s.tokenManager.SetAuthorizationMetadata(ctx)
	if err != nil {
		return nil, err
	}

	return s.enrollmentClient.GetEnrollments(ctx, &management_api_v1.GetEnrollmentsRequest{
		UserId: userId,
	})
}

// Obtains all of the active enrollment groups registered by this userId
func (s *ManagementService) GetEnrollmentGroups(userId string) (*management_api_v1.GetEnrollmentGroupsResponse, error) {
	ctx, cancel := grpc_client.GetDefaultContext()
	defer cancel()

	ctx, err := s.tokenManager.SetAuthorizationMetadata(ctx)
	if err != nil {
		return nil, err
	}

	return s.enrollmentClient.GetEnrollmentGroups(ctx, &management_api_v1.GetEnrollmentsRequest{
		UserId: userId,
	})
}

// Register a new enrollment group. Enrollment groups can contain up to 10 enrollments, and they enable multiple users to be recognized with the same request.
func (s *ManagementService) CreateEnrollmentGroups(group *management_api_v1.CreateEnrollmentGroupRequest) (*management_api_v1.EnrollmentGroupResponse, error) {
	ctx, cancel := grpc_client.GetDefaultContext()
	defer cancel()

	ctx, err := s.tokenManager.SetAuthorizationMetadata(ctx)
	if err != nil {
		return nil, err
	}

	return s.enrollmentClient.CreateEnrollmentGroup(ctx, group)
}

// Add a new enrollment to an enrollment group.
func (s *ManagementService) AppendEnrollmentGroup(request *management_api_v1.AppendEnrollmentGroupRequest) (*management_api_v1.EnrollmentGroupResponse, error) {
	ctx, cancel := grpc_client.GetDefaultContext()
	defer cancel()

	ctx, err := s.tokenManager.SetAuthorizationMetadata(ctx)
	if err != nil {
		return nil, err
	}

	return s.enrollmentClient.AppendEnrollmentGroup(ctx, request)
}

// Removes an enrollment from the system
func (s *ManagementService) DeleteEnrollment(id string) (*management_api_v1.EnrollmentResponse, error) {
	ctx, cancel := grpc_client.GetDefaultContext()
	defer cancel()

	ctx, err := s.tokenManager.SetAuthorizationMetadata(ctx)
	if err != nil {
		return nil, err
	}

	return s.enrollmentClient.DeleteEnrollment(ctx, &management_api_v1.DeleteEnrollmentRequest{Id: id})
}

// Removes an enrollment group from the system
func (s *ManagementService) DeleteEnrollmentGroup(id string) (*management_api_v1.EnrollmentGroupResponse, error) {
	ctx, cancel := grpc_client.GetDefaultContext()
	defer cancel()

	ctx, err := s.tokenManager.SetAuthorizationMetadata(ctx)
	if err != nil {
		return nil, err
	}

	return s.enrollmentClient.DeleteEnrollmentGroup(ctx, &management_api_v1.DeleteEnrollmentGroupRequest{Id: id})
}
