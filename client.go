package eramba

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

const (
	PageSize              = 200
	DefaultTimeoutSeconds = 30
)

type Pagination struct {
	HasNextPage bool  `json:"has_next_page"`
	CurrentPage int32 `json:"current_page"`
}

type responseList[K any] struct {
	Data       []K        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type responseSingle[K any] struct {
	Data       K          `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type Client struct {
	username   string
	password   string
	url        string
	httpClient *http.Client
}

func New(erambaUrl, username, password string) Client {
	return Client{
		username:   username,
		password:   password,
		url:        erambaUrl,
		httpClient: &http.Client{Timeout: DefaultTimeoutSeconds * time.Second},
	}
}

func (a *Client) BaseUrl() string {
	return strings.ReplaceAll(a.url, "/api", "")
}

func (a *Client) getByPath(ctx context.Context, path string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/%s", a.url, path), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(a.username, a.password)
	resp, err := a.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (a *Client) postOrPatchJsonByPath(ctx context.Context, method, path string, data []byte) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", a.url, path), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(a.username, a.password)
	resp, err := a.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= http.StatusBadRequest {
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		slog.Error(fmt.Sprintf("Failed %s", method), "body", body)
		return nil, fmt.Errorf("failed to %s", method)
	}
	return resp.Body, nil
}

func (a *Client) deleteById(ctx context.Context, path string, id int32) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, fmt.Sprintf("%s/%s/%d", a.url, path, id), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(a.username, a.password)
	resp, err := a.httpClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode >= http.StatusBadRequest {
		return errors.New("failed to delete")
	}
	return nil
}

func getDataById[K any](
	ctx context.Context,
	path string,
	id int32,
	getByPath func(ctx context.Context, path string) (io.ReadCloser, error),
) (K, error) {
	return getData[K](ctx, fmt.Sprintf("%s/%d", path, id), getByPath)
}

func getData[K any](
	ctx context.Context,
	path string,
	getByPath func(ctx context.Context, path string) (io.ReadCloser, error),
) (K, error) {
	res := responseSingle[K]{}
	body, err := getByPath(ctx, path)
	if err != nil {
		return res.Data, err
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&res)
	if err != nil {
		return res.Data, err
	}
	return res.Data, nil
}

func getAllData[K any](
	ctx context.Context,
	path string,
	getByPath func(ctx context.Context, path string) (io.ReadCloser, error),
) ([]K, error) {
	hasNextPage := true
	risks := make([]K, 0)
	for i := 1; hasNextPage; i++ {
		res, err := getDataForPagination[K](ctx, path, i, getByPath)
		if err != nil {
			return risks, err
		}
		hasNextPage = res.Pagination.HasNextPage
		risks = append(risks, res.Data...)
	}
	return risks, nil
}

func getDataForPagination[K any](
	ctx context.Context,
	path string,
	i int,
	getByPath func(ctx context.Context, path string) (io.ReadCloser, error),
) (responseList[K], error) {
	res := responseList[K]{}
	body, err := getByPath(ctx, fmt.Sprintf("%s?limit=%d&page=%d", path, PageSize, i))
	if err != nil {
		return res, err
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func postOrPatchJsonByPath[K any](
	ctx context.Context,
	method, path string,
	data *K,
	postByPath func(ctx context.Context, method, path string, data []byte) (io.ReadCloser, error),
) (*K, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	body, err := postByPath(ctx, method, path, dataBytes)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	res := responseSingle[K]{}

	err = json.NewDecoder(body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
