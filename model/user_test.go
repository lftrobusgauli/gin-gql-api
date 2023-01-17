package model

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	mw "thegame/middleware"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserTestSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	db   *sql.DB
}

func (s *UserTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	s.mock = mock
	s.db = db
	s.NoError(err)

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})
	mw.GetDbFromContext = func(ctx context.Context) *gorm.DB {
		return gormDB
	}
}

func (s *UserTestSuite) TearDownTest() {
	s.db.Close()
}

func (s *UserTestSuite) TestUser_Create() {
	// Arrange
	newUser := NewUser{Name: "John Doe"}
	expectedUser := &User{
		ID:         "uuid",
		Name:       "John Doe",
		GameState:  nil,
		Friends:    "null",
		FriendsArr: []string(nil),
	}

	row1 := sqlmock.NewRows([]string{"id", "name", "friends"}).
		AddRow("uuid", "John Doe", "null")
	row2 := sqlmock.NewRows([]string{"games_played", "score", "user_id", "id"}).
		AddRow(0, 0, "uuid1", "uuid2")
	query1 := `INSERT INTO "users" ("name","friends","id") VALUES ($1,$2,$3) RETURNING "id"`
	query2 := `INSERT INTO "game_states" ("games_played","score","user_id","id") VALUES ($1,$2,$3,$4) RETURNING "id"`

	// Expectations
	s.mock.ExpectBegin()

	s.mock.ExpectQuery(regexp.QuoteMeta(query1)).
		WithArgs("John Doe", "null", sqlmock.AnyArg()).
		WillReturnRows(row1)

	s.mock.ExpectQuery(regexp.QuoteMeta(query2)).
		WithArgs(0, 0, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(row2)
	s.mock.ExpectCommit()

	// Act
	user, err := (&User{}).Create(context.Background(), newUser)
	s.NoError(err)

	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
	s.Equal(expectedUser, user)
}

func (s *UserTestSuite) TestUser_GetAll() {
	row1 := sqlmock.NewRows([]string{"id", "name", "friends"}).
		AddRow("uuid", "John Doe", `["uuid1","uuid2","uuid3"]`)
	row2 := sqlmock.NewRows([]string{"id", "games_played", "score", "user_id"}).
		AddRow("uuid0", 12, 15, "uuid")

	query1 := `SELECT * FROM "users"`
	query2 := `SELECT * FROM "game_states" WHERE "game_states"."user_id" = $1`

	s.mock.ExpectQuery(regexp.QuoteMeta(query1)).
		WillReturnRows(row1)
	s.mock.ExpectQuery(regexp.QuoteMeta(query2)).
		WillReturnRows(row2)

	users, err := (&User{}).GetAll(context.Background())
	s.NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
	s.GreaterOrEqual(len(users), 1)
}

func (s *UserTestSuite) TestUser_UpdateGameState() {
	input := &UserGameState{
		UserID:      "uuid",
		GamesPlayed: GetIntPointer(12),
		Score:       GetIntPointer(25),
	}
	query1 := `SELECT * FROM "game_states" WHERE user_id = $1 ORDER BY "game_states"."id" LIMIT 1`
	query2 := `UPDATE "game_states" SET "games_played"=$1,"score"=$2,"user_id"=$3 WHERE "id" = $4`
	row1 := sqlmock.NewRows([]string{"id", "games_played", "score", "user_id"}).
		AddRow("uuid0", 12, 15, "uuid")

	s.mock.ExpectQuery(regexp.QuoteMeta(query1)).
		WillReturnRows(row1)
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(query2)).
		WithArgs(12, 25, "uuid", "uuid0").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	state, err := (&User{}).UpdateGameState(context.Background(), input)
	s.NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
	s.NotEqual(state, nil)
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
