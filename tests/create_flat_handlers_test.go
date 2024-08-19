package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"real-estate-service/api/generated"
	"real-estate-service/api/handlers"
)

type MockFlatRepository struct {
	mock.Mock
}

type MockHouseRepository struct {
	mock.Mock
}

func (m *MockHouseRepository) CreateHouse(house *generated.House) error {
	args := m.Called(house)
	return args.Error(0)
}

func (m *MockFlatRepository) UpdateFlat(flat *generated.Flat) error {
	args := m.Called(flat)
	return args.Error(0)
}

func (m *MockFlatRepository) GetFlatId(flatId generated.FlatId) (*generated.Flat, error) {
	args := m.Called(flatId)
	return args.Get(0).(*generated.Flat), args.Error(1)
}

func (m *MockFlatRepository) GetFlatsByHouseId(houseId generated.HouseId) ([]generated.Flat, error) {
	args := m.Called(houseId)
	return args.Get(0).([]generated.Flat), args.Error(1)
}

func (m *MockFlatRepository) GetApprovedFlatsByHouseId(houseId generated.HouseId) ([]generated.Flat, error) {
	args := m.Called(houseId)
	return args.Get(0).([]generated.Flat), args.Error(1)
}

func (m *MockFlatRepository) CreateFlat(flat *generated.Flat) error {
	args := m.Called(flat)
	return args.Error(0)
}

func TestPostFlatCreate(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	mockRepo := new(MockFlatRepository)
	server := &handlers.MyServer{
		FlatRepositoryInterface: mockRepo,
		Logger:                  logger,
	}

	flat := &generated.Flat{
		HouseId: 1,
		Price:   100000,
		Rooms:   3,
	}
	body, err := json.Marshal(flat)
	if err != nil {
		t.Fatalf("Failed to marshal flat: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/flat/create", bytes.NewBuffer(body))
	req = req.WithContext(context.WithValue(req.Context(), generated.BearerAuthScopes, "moderator"))
	rec := httptest.NewRecorder()

	mockRepo.On("CreateFlat", flat).Return(nil)

	server.PostFlatCreate(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	err = json.NewDecoder(rec.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	assert.Equal(t, "Успешно создана квартира", response["message"])

	data, ok := response["data"].(map[string]interface{})
	if !ok {
		t.Fatal("Response data is not of expected type")
	}
	assert.Equal(t, 1, int(data["house_id"].(float64)))
	assert.Equal(t, 100000, int(data["price"].(float64)))
	assert.Equal(t, 3, int(data["rooms"].(float64)))

	mockRepo.AssertExpectations(t)
}

func TestGetFlatsByHouseId(t *testing.T) {
	mockRepo := new(MockFlatRepository)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	server := &handlers.MyServer{
		Logger:                  logger,
		FlatRepositoryInterface: mockRepo,
	}

	houseId := generated.HouseId(1)
	flats := []generated.Flat{
		{Id: 1, HouseId: houseId, Price: 100000, Rooms: 3, Status: "approved"},
		{Id: 2, HouseId: houseId, Price: 150000, Rooms: 4, Status: "created"},
	}

	// Настройка ожиданий
	mockRepo.On("GetApprovedFlatsByHouseId", houseId).Return(flats[:1], nil) // Клиент должен получать только одобренные квартиры

	req := httptest.NewRequest(http.MethodGet, "/house/"+fmt.Sprintf("%d", houseId), nil)
	req = req.WithContext(context.WithValue(req.Context(), generated.BearerAuthScopes, "client"))
	rec := httptest.NewRecorder()

	server.GetHouseId(rec, req, houseId)

	assert.Equal(t, http.StatusOK, rec.Code)
	var response []generated.Flat
	err := json.NewDecoder(rec.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(response)) // Только одобренные квартиры
	assert.Equal(t, flats[0], response[0])

	mockRepo.AssertExpectations(t)
}

func TestGetFlatsByHouseIdAsModerator(t *testing.T) {
	mockRepo := new(MockFlatRepository)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	server := &handlers.MyServer{
		Logger:                  logger,
		FlatRepositoryInterface: mockRepo,
	}

	houseId := generated.HouseId(1)
	flats := []generated.Flat{
		{Id: 1, HouseId: houseId, Price: 100000, Rooms: 3, Status: "approved"},
		{Id: 2, HouseId: houseId, Price: 150000, Rooms: 4, Status: "created"},
	}

	// Настройка ожиданий
	mockRepo.On("GetFlatsByHouseId", houseId).Return(flats, nil)
	req := httptest.NewRequest(http.MethodGet, "/house/"+fmt.Sprintf("%d", houseId), nil)
	req = req.WithContext(context.WithValue(req.Context(), generated.BearerAuthScopes, "moderator"))
	rec := httptest.NewRecorder()

	server.GetHouseId(rec, req, houseId)

	assert.Equal(t, http.StatusOK, rec.Code)
	var response []generated.Flat
	err := json.NewDecoder(rec.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(response))
	assert.ElementsMatch(t, flats, response)

	mockRepo.AssertExpectations(t)
}
