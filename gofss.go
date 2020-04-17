package gofss

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client interface {
	GeneratePDF(string) ([]byte, error)
}

type defaultClient struct {
	serviceURL string
}

func (c *defaultClient) GeneratePDF(html string) ([]byte, error) {
	type msg struct {
		Value string `json:"value"`
	}

	reqBuff := &bytes.Buffer{}
	req := msg{Value: html}
	err := json.NewEncoder(reqBuff).Encode(req)
	if err != nil {
		return nil, fmt.Errorf("could not encode request: %w", err)
	}

	resp, err := http.Post(c.serviceURL, "application/json", reqBuff)
	if err != nil {
		return nil, fmt.Errorf("could not generate pdf: %w", err)
	}

	defer closeQuietly(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid pdf service response code: %d", resp.StatusCode)
	}

	jsonResp := msg{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, fmt.Errorf("could not decode response from pdf service: %w", err)
	}

	pdf, err := base64.StdEncoding.DecodeString(jsonResp.Value)
	if err != nil {
		return nil, fmt.Errorf("could not decode base64 pdf: %w", err)
	}

	return pdf, nil
}

func closeQuietly(c io.Closer) {
	_ = c.Close()
}

func New(serviceURL string) Client {
	return &defaultClient{serviceURL: serviceURL}
}
