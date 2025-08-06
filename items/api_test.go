package items

import (
	"context"
	"testing"

	"github.com/tidyrocks/mercado-libre-go-sdk/internal/testenv"
)

func TestGetByID(t *testing.T) {
	item, err := GetByID(context.Background(), testenv.TestItemID, testenv.AccessToken)
	if err != nil {
		t.Fatalf("Error getting item: %v", err)
	}

	if item.ID != testenv.TestItemID {
		t.Errorf("Expected item ID %s, got %s", testenv.TestItemID, item.ID)
	}

	if item.Title == "" {
		t.Error("Item title is empty")
	}

	if item.Price <= 0 {
		t.Error("Item price should be greater than 0")
	}

	t.Logf("Item: %s - %s - $%.2f", item.ID, item.Title, item.Price)
	t.Logf("Status: %s, Condition: %s", item.Status, item.ItemCondition)
	t.Logf("Available: %d", item.AvailableQuantity)
}
