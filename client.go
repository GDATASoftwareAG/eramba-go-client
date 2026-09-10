package eramba

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/gdatasoftwareag/eramba-go-client/model"
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
	if resp.StatusCode >= http.StatusBadRequest {
		return nil, a.errorHandling(resp, http.MethodGet)
	}
	return resp.Body, nil
}

func (a *Client) postByPath(ctx context.Context, method, path string, data []byte) (io.ReadCloser, error) {
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
		return nil, a.errorHandling(resp, method)
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
		return a.errorHandling(resp, http.MethodDelete)
	}
	return nil
}

func (a *Client) errorHandling(resp *http.Response, method string) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	slog.Error(fmt.Sprintf("failed %s", method), "body", body, "url", resp.Request.URL.String())
	return fmt.Errorf("unexpected status code: %d for method: %s", resp.StatusCode, method)
}

func (a *Client) getDataById[K any](
	ctx context.Context,
	path string,
	id int32,
) (K, error) {
	return a.getData[K](ctx, fmt.Sprintf("%s/%d", path, id))
}

func (a *Client) getData[K any](
	ctx context.Context,
	path string,
) (K, error) {
	res := responseSingle[K]{}
	body, err := a.getByPath(ctx, path)
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

func (a *Client) getAllData[K any](
	ctx context.Context,
	path string,
) ([]K, error) {
	hasNextPage := true
	risks := make([]K, 0)
	for i := 1; hasNextPage; i++ {
		res, err := a.getDataForPagination[K](ctx, path, i)
		if err != nil {
			return risks, err
		}
		hasNextPage = res.Pagination.HasNextPage
		risks = append(risks, res.Data...)
	}
	return risks, nil
}

func (a *Client) getDataForPagination[K any](
	ctx context.Context,
	path string,
	i int,
) (responseList[K], error) {
	res := responseList[K]{}
	body, err := a.getByPath(ctx, fmt.Sprintf("%s?limit=%d&page=%d", path, PageSize, i))
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

func (a *Client) postOrPatchJsonByPath[K any](
	ctx context.Context,
	method, path string,
	data *K,
) (*K, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	body, err := a.postByPath(ctx, method, path, dataBytes)
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

type GetClient[K any] struct {
	client *Client
	path   string
}

func (a *GetClient[K]) Get(ctx context.Context, id int32) (K, error) {
	return a.client.getDataById[K](ctx, a.path, id)
}

func (a *GetClient[K]) GetAll(ctx context.Context) ([]K, error) {
	return a.client.getAllData[K](ctx, fmt.Sprintf("%s/index", a.path))
}

type GetAndPatchClient[K any] struct {
	GetClient[K]
}

func (a *GetAndPatchClient[K]) Post(ctx context.Context, data *K) (*K, error) {
	return a.client.postOrPatchJsonByPath(ctx, http.MethodPost, fmt.Sprintf("%s/add", a.path), data)
}

func (a *GetAndPatchClient[K]) Patch(
	ctx context.Context,
	id int32,
	data *K,
) (*K, error) {
	return a.client.postOrPatchJsonByPath(ctx, http.MethodPatch, fmt.Sprintf("%s/%d", a.path, id), data)
}

func (a *GetAndPatchClient[K]) Delete(ctx context.Context, id int32) error {
	return a.client.deleteById(ctx, a.path, id)
}

type GetAndPatchClientWithComment[K any] struct {
	GetAndPatchClient[K]
}

func (a *GetAndPatchClientWithComment[K]) Comments() *CommentsClient {
	return &CommentsClient{
		client: a.client,
		path:   a.path,
	}
}

type GetAndPatchClientWithCommentAndReview[K any] struct {
	GetAndPatchClientWithComment[K]
	pathReview string
	model      string
}

func (a *GetAndPatchClientWithCommentAndReview[K]) Reviews() *ReviewsClient {
	return &ReviewsClient{
		Model:  a.model,
		client: a.client,
		path:   a.pathReview,
	}
}

type ReviewsClient struct {
	Model string
	GetAndPatchClient[model.Review]
}
