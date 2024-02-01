package items_test

import (
	oapicodegen "openapi/internal/infrastructure/oapicodegen/stock"
	"testing"

	_ "github.com/lib/pq"

	"github.com/google/uuid"

	"encoding/json"
	"io"
	"net/http"
)

func TestDeleteOk(t *testing.T) {

	// Setup
	client := &http.Client{}
	name := uuid.NewString()

	// Given
	postRes, err := PostHelper(
		client,
		&oapicodegen.PostStockItemJSONRequestBody{
			Name: name,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer postRes.Body.Close()

	if postRes.StatusCode != http.StatusCreated {
		t.Fatalf("expected not empty, actual empty")
	}

	postResBodyByte, _ := io.ReadAll(postRes.Body)
	postResBody := &oapicodegen.Created{}
	json.Unmarshal(postResBodyByte, &postResBody)
	if postResBody.Id == uuid.Nil {
		t.Fatal(err)
	}

	// When
	deleteRes, err := DeleteHelper(
		client,
		postResBody.Id,
	)
	if err != nil {
		t.Fatal(err)
	}
	defer deleteRes.Body.Close()

	deleteResBodyByte, _ := io.ReadAll(deleteRes.Body)
	deleteResBody := &oapicodegen.Created{}
	json.Unmarshal(deleteResBodyByte, &deleteResBody)

	// Then
	if deleteRes.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, deleteRes.StatusCode)
	}
}

func TestDeleteNotFound(t *testing.T) {
	// Setup
	client := &http.Client{}

	deleteRes, err := DeleteHelper(
		client,
		uuid.New(),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer deleteRes.Body.Close()

	deleteResBodyByte, _ := io.ReadAll(deleteRes.Body)
	deleteResBody := &oapicodegen.Created{}
	json.Unmarshal(deleteResBodyByte, &deleteResBody)

	// Then
	if deleteRes.StatusCode != http.StatusNotFound {
		t.Errorf("want %d, got %d", http.StatusNotFound, deleteRes.StatusCode)
	}
}
