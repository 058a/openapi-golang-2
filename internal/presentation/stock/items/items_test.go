package items

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"openapi/internal/infrastructure/env"
	oapicodegen "openapi/internal/infrastructure/oapicodegen/stock"
	"testing"

	"github.com/google/uuid"
)

func TestOapiCodegenEcho(t *testing.T) {
	// Setup
	client := http.Client{}
	name := uuid.NewString()

	putReqBody := &oapicodegen.PutStockItemJSONRequestBody{
		Name: name,
	}
	putReqBodyJson, _ := json.Marshal(putReqBody)
	putReq, err := http.NewRequest(
		http.MethodPut,
		env.GetServiceUrl()+"/stock/items/"+uuid.NewString(),
		bytes.NewBuffer(putReqBodyJson))
	if err != nil {
		t.Fatal(err)
	}
	putReq.Header.Set("Content-Type", "application/json")
	putRes, err := client.Do(putReq)
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
