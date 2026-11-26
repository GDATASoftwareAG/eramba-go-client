package eramba_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	eramba "github.com/gdatasoftwareag/eramba-go-client"
	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func newRecordingServer(t *testing.T) (srv *httptest.Server, requests *[]string) {
	t.Helper()
	collected := []string{}
	requests = &collected
	srv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		collected = append(collected, r.Method+" "+r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/comments") {
			_, _ = w.Write([]byte(`{"data":[{"id":1,"message":"hi"}],"pagination":{"has_next_page":false}}`))
			return
		}
		_, _ = w.Write([]byte(`{"data":{"id":1,"message":"hi"}}`))
	}))
	t.Cleanup(srv.Close)
	return srv, requests
}

func exerciseComments(t *testing.T, comments *eramba.CommentsClient) {
	t.Helper()
	ctx := context.Background()
	if _, err := comments.GetComments(ctx, 7); err != nil {
		t.Fatal(err)
	}
	if _, err := comments.GetComment(ctx, 7, 1); err != nil {
		t.Fatal(err)
	}
	if _, err := comments.PostComment(ctx, 7, &model.Comment{Message: "hi"}); err != nil {
		t.Fatal(err)
	}
	if err := comments.DeleteComment(ctx, 7, 1); err != nil {
		t.Fatal(err)
	}
}

func assertRequests(t *testing.T, got, want []string) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("requests = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("request %d = %q, want %q", i, got[i], want[i])
		}
	}
}

func Test_SecurityIncidentComments_Paths(t *testing.T) {
	srv, requests := newRecordingServer(t)
	client := eramba.New(srv.URL+"/api", "u", "p")
	exerciseComments(t, client.SecurityIncidents().Comments())
	assertRequests(t, *requests, []string{
		"GET /api/v2/security-incidents/7/comments",
		"GET /api/v2/security-incidents/7/comments/1",
		"POST /api/v2/security-incidents/7/comments",
		"DELETE /api/v2/security-incidents/7/comments/1",
	})
}

func Test_RiskComments_Paths(t *testing.T) {
	srv, requests := newRecordingServer(t)
	client := eramba.New(srv.URL+"/api", "u", "p")
	exerciseComments(t, client.Risks().Comments())
	assertRequests(t, *requests, []string{
		"GET /api/risks/7/comments",
		"GET /api/risks/7/comments/1",
		"POST /api/risks/7/comments/add",
		"DELETE /api/risks/7/comments/1",
	})
}
