package controller

import (
	"context"
	"testing"

	"thegame/model"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockUser struct {
	mock.Mock
}

func (u *mockUser) Create(ctx context.Context, input model.NewUser) (*model.User, error) {
	args := u.Called(ctx, input)
	return args.Get(0).(*model.User), args.Error(1)
}

func (u *mockUser) UpdateGameState(ctx context.Context, input *model.UserGameState) (*model.GameState, error) {
	args := u.Called(ctx, input)
	return args.Get(0).(*model.GameState), args.Error(1)
}

func (u *mockUser) AddFriends(ctx context.Context, userID string, friends []string) ([]*model.Friend, error) {
	args := u.Called(ctx, userID, friends)
	return args.Get(0).([]*model.Friend), args.Error(1)
}

func (u *mockUser) GetAll(ctx context.Context) ([]*model.User, error) {
	args := u.Called(ctx)
	return args.Get(0).([]*model.User), args.Error(1)
}

func (u *mockUser) GetGameState(ctx context.Context, userID string) (*model.GameState, error) {
	args := u.Called(ctx, userID)
	return args.Get(0).(*model.GameState), args.Error(1)
}

func (u *mockUser) GetFriends(ctx context.Context, userID string) ([]*model.Friend, error) {
	args := u.Called(ctx, userID)
	return args.Get(0).([]*model.Friend), args.Error(1)
}

func getIntPointer(value int) *int {
	return &value
}

type UserControllerSuite struct {
	suite.Suite
	resolver      *mutationResolver
	queryResolver *queryResolver
	user          *mockUser
}

func (s *UserControllerSuite) SetupTest() {
	s.user = &mockUser{}
	s.resolver = &mutationResolver{&Resolver{UserService: s.user}}
	s.queryResolver = &queryResolver{&Resolver{UserService: s.user}}
}

// TODO: Test using table driven tests

func (s *UserControllerSuite) TestCreateUser() {
	// Prepare test data
	input := model.NewUser{
		Name: "John Doe",
	}
	expectedUser := &model.User{
		ID:   "1",
		Name: "John Doe",
	}

	// Prepare mock
	s.user.On("Create", mock.Anything, input).Return(expectedUser, nil)

	// Execute test
	user, err := s.resolver.CreateUser(context.TODO(), input)

	// Assert result
	s.NoError(err)
	s.Equal(expectedUser, user)
	s.user.AssertExpectations(s.T())
}

func (s *UserControllerSuite) TestUpdateGameState() {
	// Prepare test data
	input := model.UserGameState{
		UserID:      "1",
		GamesPlayed: getIntPointer(10),
		Score:       getIntPointer(10),
	}
	expected := &model.GameState{
		ID:          "1",
		GamesPlayed: getIntPointer(10),
		Score:       getIntPointer(10),
		UserID:      "1",
	}

	// Prepare mock
	s.user.On("UpdateGameState", mock.Anything, &input).Return(expected, nil)

	// Execute test
	actual, err := s.resolver.UpdateGameState(context.TODO(), input)

	// Assert result
	s.NoError(err)
	s.Equal(expected, actual)
	s.user.AssertExpectations(s.T())
}

func (s *UserControllerSuite) TestAddFriends() {
	// Prepare test data
	input1 := "1"
	input2 := []string{"1", "2"}

	f := &model.Friend{
		ID:        "1",
		Name:      "Musashi",
		Highscore: 100,
	}

	expected := []*model.Friend{f}

	// Prepare mock
	s.user.On("AddFriends", mock.Anything, input1, input2).Return(expected, nil)

	// Execute test
	actual, err := s.resolver.AddFriends(context.TODO(), input1, input2)

	// Assert result
	s.NoError(err)
	s.Equal(expected, actual)
	s.user.AssertExpectations(s.T())
}

func (s *UserControllerSuite) TestGetAll() {
	// Prepare test data

	u := &model.User{
		ID:   "1",
		Name: "Miyamoto",
	}

	expected := []*model.User{u}

	// Prepare mock
	s.user.On("GetAll", mock.Anything).Return(expected, nil)

	// Execute test
	actual, err := s.queryResolver.Users(context.TODO())

	// Assert result
	s.NoError(err)
	s.Equal(expected, actual)
	s.user.AssertExpectations(s.T())
}

func (s *UserControllerSuite) TestGameState() {
	// Prepare test data
	input := "1"
	expected := &model.GameState{
		ID:          "1",
		GamesPlayed: getIntPointer(10),
		Score:       getIntPointer(10),
		UserID:      "1",
	}

	// Prepare mock
	s.user.On("GetGameState", mock.Anything, input).Return(expected, nil)

	// Execute test
	actual, err := s.queryResolver.GetGameState(context.TODO(), input)

	// Assert result
	s.NoError(err)
	s.Equal(expected, actual)
	s.user.AssertExpectations(s.T())
}

func (s *UserControllerSuite) TestGetFriends() {
	// Prepare test data
	input := "1"
	f := &model.Friend{
		ID:        "1",
		Name:      "Musashi",
		Highscore: 100,
	}

	expected := []*model.Friend{f}

	// Prepare mock
	s.user.On("GetFriends", mock.Anything, input).Return(expected, nil)

	// Execute test
	actual, err := s.queryResolver.GetFriends(context.TODO(), input)

	// Assert result
	s.NoError(err)
	s.Equal(expected, actual)
	s.user.AssertExpectations(s.T())
}

func TestUserControllerSuite(t *testing.T) {
	suite.Run(t, new(UserControllerSuite))
}
