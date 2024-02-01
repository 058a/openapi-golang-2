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

func TestPutOk(t *testing.T) {

	// Setup
	client := &http.Client{}
	bforeName := uuid.NewString()
	afterName := uuid.NewString()

	// Given
	postReqBody := &oapicodegen.PostStockItemJSONRequestBody{
		Name: bforeName,
	}
	postRes, err := PostHelper(client, postReqBody)
	if err != nil {
		t.Fatal(err)
	}
	defer postRes.Body.Close()

	if postRes.StatusCode != http.StatusCreated {
		t.Fatal(err)
	}

	postResBodyByte, _ := io.ReadAll(postRes.Body)
	postResBody := &oapicodegen.Created{}
	json.Unmarshal(postResBodyByte, &postResBody)
	if postResBody.Id == uuid.Nil {
		t.Fatalf("expected not empty, actual empty")
	}

	// When
	putReqBody := &oapicodegen.PutStockItemJSONRequestBody{
		Name: afterName,
	}
	putRes, err := PutHelper(client, postResBody.Id, putReqBody)
	if err != nil {
		t.Fatal(err)
	}
	defer putRes.Body.Close()

	putResBodyByte, _ := io.ReadAll(putRes.Body)
	putResBody := &oapicodegen.Created{}
	json.Unmarshal(putResBodyByte, &putResBody)

	// Then
	if putRes.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, putRes.StatusCode)
	}
}

func TestPutNotFound(t *testing.T) {
	// Setup
	client := &http.Client{}
	name := uuid.NewString()

	putReqBody := &oapicodegen.PutStockItemJSONRequestBody{
		Name: name,
	}
	putRes, err := PutHelper(client, uuid.New(), putReqBody)
	if err != nil {
		t.Fatal(err)
	}
	defer putRes.Body.Close()

	putResBodyByte, _ := io.ReadAll(putRes.Body)
	putResBody := &oapicodegen.Created{}
	json.Unmarshal(putResBodyByte, &putResBody)

	// Then
	if putRes.StatusCode != http.StatusNotFound {
		t.Errorf("want %d, got %d", http.StatusNotFound, putRes.StatusCode)
	}
}

func TestPutBadRequest(t *testing.T) {
	// Setup
	client := &http.Client{}
	name := uuid.NewString()
	nameLenZero := ""
	nameLenOver := strings.Repeat("a", 101)

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
		t.Fatal(err)
	}

	postResBodyByte, _ := io.ReadAll(postRes.Body)
	postResBody := &oapicodegen.Created{}
	json.Unmarshal(postResBodyByte, &postResBody)
	if postResBody.Id == uuid.Nil {
		t.Errorf("expected not empty, actual empty")
	}

	// When
	putResLenZero, err := PutHelper(
		client,
		postResBody.Id,
		&oapicodegen.PutStockItemJSONRequestBody{
			Name: nameLenZero,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer putResLenZero.Body.Close()

	putResLenOver, err := PutHelper(
		client,
		postResBody.Id,
		&oapicodegen.PutStockItemJSONRequestBody{
			Name: nameLenOver,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer putResLenOver.Body.Close()

	// Then
	if putResLenZero.StatusCode != http.StatusBadRequest {
		t.Errorf("want %d, got %d", http.StatusBadRequest, putResLenZero.StatusCode)
	}

	if putResLenOver.StatusCode != http.StatusBadRequest {
		t.Errorf("want %d, got %d", http.StatusBadRequest, putResLenOver.StatusCode)
	}
}
