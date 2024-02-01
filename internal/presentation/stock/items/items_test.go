package items_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"openapi/internal/infrastructure/env"
	oapicodegen "openapi/internal/infrastructure/oapicodegen/stock"

	"github.com/google/uuid"
)

func PostHelper(client *http.Client, reqBody *oapicodegen.PostStockItemJSONRequestBody) (*http.Response, error) {
	reqBodyJson, _ := json.Marshal(reqBody)
	req, err := http.NewRequest(
		http.MethodPost,
		env.GetServiceUrl()+"/stock/items",
		bytes.NewBuffer(reqBodyJson),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func PutHelper(client *http.Client, stockitemId uuid.UUID, reqBody *oapicodegen.PutStockItemJSONRequestBody) (*http.Response, error) {
	reqBodyJson, _ := json.Marshal(reqBody)
	req, err := http.NewRequest(
		http.MethodPut,
		env.GetServiceUrl()+"/stock/items/"+stockitemId.String(),
		bytes.NewBuffer(reqBodyJson),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteHelper(client *http.Client, stockitemId uuid.UUID) (*http.Response, error) {
	req, err := http.NewRequest(
		http.MethodDelete,
		env.GetServiceUrl()+"/stock/items/"+stockitemId.String(),
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
