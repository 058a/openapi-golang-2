package items_test

import (
	oapicodegen "openapi/internal/infrastructure/oapicodegen/stock"
	"strings"
	"testing"

	_ "github.com/lib/pq"

	"github.com/google/uuid"

	"encoding/json"
	"io"
	"net/http"
)

func TestPostCreated(t *testing.T) {
	// Setup
	client := &http.Client{}
	name := uuid.NewString()

	// When
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

	// Then
	if postRes.StatusCode != http.StatusCreated {
		t.Errorf("want %d, got %d", http.StatusCreated, postRes.StatusCode)
	}

	postResBodyByte, _ := io.ReadAll(postRes.Body)
	postResBody := &oapicodegen.Created{}
	json.Unmarshal(postResBodyByte, &postResBody)
	if postResBody.Id == uuid.Nil {
		t.Errorf("expected not empty, actual empty")
	}

}

func TestPostBadRequest(t *testing.T) {
	// Setup
	client := &http.Client{}
	nameLenZero := ""
	nameLenOver := strings.Repeat("a", 101)

	// When
	postResLenZero, err := PostHelper(
		client,
		&oapicodegen.PostStockItemJSONRequestBody{
			Name: nameLenZero,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer postResLenZero.Body.Close()

	postResLenOver, err := PostHelper(
		client,
		&oapicodegen.PostStockItemJSONRequestBody{
			Name: nameLenOver,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer postResLenOver.Body.Close()

	// Then
	if postResLenZero.StatusCode != http.StatusBadRequest {
		t.Errorf("want %d, got %d", http.StatusBadRequest, postResLenZero.StatusCode)
	}
	if postResLenOver.StatusCode != http.StatusBadRequest {
		t.Errorf("want %d, got %d", http.StatusBadRequest, postResLenOver.StatusCode)
	}
}
