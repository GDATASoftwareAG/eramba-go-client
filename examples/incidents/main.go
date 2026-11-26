package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	erambaClient "github.com/gdatasoftwareag/eramba-go-client"
	"github.com/gdatasoftwareag/eramba-go-client/examples/utils"
	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func main() {
	utils.LoadEnvs()
	client := utils.CreateClientFromEnv()
	ctx := context.Background()
	if err := run(ctx, &client); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, client *erambaClient.Client) (err error) {
	incidents := client.SecurityIncidents()

	all, err := incidents.GetAll(ctx)
	if err != nil {
		return fmt.Errorf("getting security incidents: %w", err)
	}
	if len(all) == 0 {
		return errors.New("no security incidents found, create one to run this example")
	}
	for i := range all {
		printIncident("Existing", &all[i])
	}

	// owners and contacts are required, reuse them from an existing incident
	template := all[0]
	created, err := incidents.Post(ctx, &model.SecurityIncident{
		Title:       "eramba-go-client example incident",
		Description: "created by examples/incidents",
		Type:        model.SecurityIncidentTypeEvent,
		Status:      model.SecurityIncidentStatusOpen,
		Owners:      template.Owners,
		Contacts:    template.Contacts,
		OpenDate:    model.RefErambaDate(time.Now()),
		Reporter:    "eramba-go-client",
	})
	if err != nil {
		return fmt.Errorf("creating security incident: %w", err)
	}
	incidentId := created.Id
	printIncident("Created", created)

	deleted := false
	defer func() {
		if deleted {
			return
		}
		if cleanupErr := incidents.Delete(ctx, incidentId); cleanupErr != nil {
			err = errors.Join(err, fmt.Errorf("cleanup of security incident %d: %w", incidentId, cleanupErr))
			return
		}
		fmt.Printf("Cleaned up security incident %d\n", incidentId)
	}()

	fetched, err := incidents.Get(ctx, incidentId)
	if err != nil {
		return fmt.Errorf("getting security incident %d: %w", incidentId, err)
	}
	printIncident("Fetched", &fetched)

	fetched.Title += " (updated)"
	fetched.Type = model.SecurityIncidentTypeIncident
	fetched.Status = model.SecurityIncidentStatusClosed
	fetched.ClosureDate = model.RefErambaDate(time.Now())
	updated, err := incidents.Put(ctx, incidentId, &fetched)
	if err != nil {
		return fmt.Errorf("updating security incident %d: %w", incidentId, err)
	}
	printIncident("Updated", updated)
	if updated.Title != fetched.Title || updated.Status != model.SecurityIncidentStatusClosed {
		return fmt.Errorf("update not applied: title=%q status=%d", updated.Title, updated.Status)
	}

	if err := testRiskLinking(ctx, client, incidentId); err != nil {
		return err
	}

	if err := testComments(ctx, incidents.Comments(), incidentId); err != nil {
		return err
	}

	if err := incidents.Delete(ctx, incidentId); err != nil {
		return fmt.Errorf("deleting security incident %d: %w", incidentId, err)
	}
	deleted = true
	fmt.Printf("Deleted security incident %d\n", incidentId)

	// expected to fail, the client logs the API error response
	if _, err := incidents.Get(ctx, incidentId); err == nil {
		return fmt.Errorf("security incident %d still exists after delete", incidentId)
	}
	fmt.Printf("Verified security incident %d is gone\n", incidentId)
	return nil
}

// printIncident only logs marshal errors so the deferred cleanup in run still deletes the example incident.
func printIncident(prefix string, incident *model.SecurityIncident) {
	data, err := json.Marshal(incident)
	if err != nil {
		fmt.Printf("%s security incident %d: marshal failed: %v\n", prefix, incident.Id, err)
		return
	}
	fmt.Printf("%s security incident %d: %s\n", prefix, incident.Id, data)
}

// testRiskLinking links an existing risk to the incident and removes it again.
func testRiskLinking(ctx context.Context, client *erambaClient.Client, incidentId int32) error {
	risks, err := client.Risks().GetAll(ctx)
	if err != nil {
		return fmt.Errorf("getting risks: %w", err)
	}
	if len(risks) == 0 {
		fmt.Println("No risks found, skipping risk linking")
		return nil
	}
	risk := &risks[0]
	incidents := client.SecurityIncidents()

	incident, err := incidents.Get(ctx, incidentId)
	if err != nil {
		return fmt.Errorf("getting security incident %d: %w", incidentId, err)
	}
	incident.Risks = model.Risks{risk}
	if _, err := incidents.Put(ctx, incidentId, &incident); err != nil {
		return fmt.Errorf("linking risk %d to security incident %d: %w", risk.Id, incidentId, err)
	}
	linked, err := incidents.Get(ctx, incidentId)
	if err != nil {
		return fmt.Errorf("getting security incident %d: %w", incidentId, err)
	}
	printIncident("Linked risk", &linked)
	if len(linked.Risks) != 1 || linked.Risks[0].Id != risk.Id {
		return fmt.Errorf("risk %d not linked, got %d risks", risk.Id, len(linked.Risks))
	}
	fmt.Printf("Linked risk %d to security incident %d\n", risk.Id, incidentId)

	linked.Risks = model.Risks{}
	if _, err := incidents.Put(ctx, incidentId, &linked); err != nil {
		return fmt.Errorf("unlinking risk %d from security incident %d: %w", risk.Id, incidentId, err)
	}
	unlinked, err := incidents.Get(ctx, incidentId)
	if err != nil {
		return fmt.Errorf("getting security incident %d: %w", incidentId, err)
	}
	printIncident("Unlinked risk", &unlinked)
	if len(unlinked.Risks) != 0 {
		return fmt.Errorf("risk %d still linked, got %d risks", risk.Id, len(unlinked.Risks))
	}
	fmt.Printf("Unlinked risk %d from security incident %d\n", risk.Id, incidentId)
	return nil
}

func testComments(ctx context.Context, comments *erambaClient.CommentsClient, incidentId int32) error {
	created, err := comments.PostComment(ctx, incidentId, &model.Comment{Message: "eramba-go-client example comment"})
	if err != nil {
		return fmt.Errorf("posting comment: %w", err)
	}
	fmt.Printf("Created comment %d: %q\n", created.Id, created.Message)

	fetched, err := comments.GetComment(ctx, incidentId, created.Id)
	if err != nil {
		return fmt.Errorf("getting comment %d: %w", created.Id, err)
	}
	fmt.Printf("Fetched comment %d: %q\n", fetched.Id, fetched.Message)

	all, err := comments.GetComments(ctx, incidentId)
	if err != nil {
		return fmt.Errorf("listing comments: %w", err)
	}
	for _, comment := range all {
		fmt.Printf("Listed comment %d: %q\n", comment.Id, comment.Message)
	}

	if err := comments.DeleteComment(ctx, incidentId, created.Id); err != nil {
		return fmt.Errorf("deleting comment %d: %w", created.Id, err)
	}
	fmt.Printf("Deleted comment %d\n", created.Id)
	return nil
}
