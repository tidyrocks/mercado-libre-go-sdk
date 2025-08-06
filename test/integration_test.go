package test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidyrocks/mercado-libre-go-sdk/categories"
	"github.com/tidyrocks/mercado-libre-go-sdk/internal/logger"
	"github.com/tidyrocks/mercado-libre-go-sdk/internal/metrics"
	"github.com/tidyrocks/mercado-libre-go-sdk/internal/testenv"
	"github.com/tidyrocks/mercado-libre-go-sdk/internal/validation"
	"github.com/tidyrocks/mercado-libre-go-sdk/items"
)

func TestIntegration_FullWorkflow(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	ctx := context.Background()

	// Test con logging estructurado
	logger.LogTestStep(ctx, "TestIntegration_FullWorkflow", "start", map[string]interface{}{
		"test_item_id": testenv.TestItemID,
	})

	t.Run("validate_inputs", func(t *testing.T) {
		// Validar access token
		err := validation.ValidateAccessToken(testenv.AccessToken)
		require.NoError(t, err)

		// Validar item ID
		err = validation.ValidateItemID(testenv.TestItemID)
		require.NoError(t, err)
	})

	t.Run("get_item_details", func(t *testing.T) {
		start := time.Now()

		item, err := items.GetByID(ctx, testenv.TestItemID, testenv.AccessToken)
		duration := time.Since(start)

		require.NoError(t, err)
		require.NotNil(t, item)

		// Validar campos críticos
		assert.Equal(t, testenv.TestItemID, item.ID)
		assert.NotEmpty(t, item.Title)
		assert.Greater(t, item.Price, 0.0)
		assert.NotEmpty(t, item.CategoryID)

		// Log métricas
		metrics.RecordHTTPCall(ctx, "GET", "/items/"+testenv.TestItemID, 200, duration)

		logger.LogTestStep(ctx, "TestIntegration_FullWorkflow", "item_retrieved", map[string]interface{}{
			"item_id":     item.ID,
			"title":       item.Title,
			"price":       item.Price,
			"category":    item.CategoryID,
			"duration_ms": duration.Milliseconds(),
		})
	})

	t.Run("get_category_details", func(t *testing.T) {
		// Primero obtener el ítem para sacar su categoría
		item, err := items.GetByID(ctx, testenv.TestItemID, testenv.AccessToken)
		require.NoError(t, err)

		// Obtener detalles de la categoría
		start := time.Now()
		category, err := categories.GetByID(item.CategoryID, testenv.AccessToken)
		duration := time.Since(start)

		require.NoError(t, err)
		require.NotNil(t, category)

		assert.Equal(t, item.CategoryID, category.ID)
		assert.NotEmpty(t, category.Name)

		logger.LogTestStep(ctx, "TestIntegration_FullWorkflow", "category_retrieved", map[string]interface{}{
			"category_id": category.ID,
			"name":        category.Name,
			"duration_ms": duration.Milliseconds(),
		})
	})

	t.Run("validate_workflow_metrics", func(t *testing.T) {
		// Obtener métricas del collector in-memory
		collector := metrics.GetCollector()

		if inMemory, ok := collector.(*metrics.InMemoryMetrics); ok {
			stats := inMemory.GetStats()

			// Verificar que se registraron llamadas HTTP
			counters := stats["counters"].(map[string]int64)
			assert.Greater(t, counters["http_requests_total"], int64(0))

			t.Logf("Integration test metrics: %+v", stats)
		}
	})
}

func TestIntegration_ErrorHandling(t *testing.T) {
	ctx := context.Background()

	t.Run("invalid_access_token", func(t *testing.T) {
		invalidToken := "invalid-token"

		_, err := items.GetByID(ctx, testenv.TestItemID, invalidToken)
		assert.Error(t, err)

		// Verificar que la validación catch el error
		validationErr := validation.ValidateAccessToken(invalidToken)
		assert.Error(t, validationErr)
	})

	t.Run("invalid_item_id", func(t *testing.T) {
		invalidItemID := "INVALID123"

		// Validación debe fallar antes de hacer el request
		err := validation.ValidateItemID(invalidItemID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "must match pattern")
	})

	t.Run("nonexistent_item", func(t *testing.T) {
		nonExistentID := "MLM999999999999"

		_, err := items.GetByID(ctx, nonExistentID, testenv.AccessToken)
		assert.Error(t, err)
		// Esperamos un 404 de la API
	})
}

func TestIntegration_MockServer(t *testing.T) {
	// Crear mock server para tests unitarios
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/items/MLM123456":
			response := map[string]interface{}{
				"id":                 "MLM123456",
				"title":              "Mock Item",
				"price":              100.50,
				"category_id":        "MLM1051",
				"status":             "active",
				"condition":          "new",
				"available_quantity": 10,
				"currency_id":        "MXN",
				"buying_mode":        "buy_it_now",
				"listing_type_id":    "gold_pro",
				"date_created":       time.Now().Format(time.RFC3339),
				"shipping": map[string]interface{}{
					"mode":          "me2",
					"logistic_type": "fulfillment",
				},
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		default:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, `{"error": "not_found", "message": "Resource not found"}`)
		}
	}))
	defer server.Close()

	t.Run("mock_item_retrieval", func(t *testing.T) {
		// Aquí normalmente inyectarías la URL del mock server
		// En una implementación real, tendrías configuración para override de URLs
		t.Log("Mock server running at:", server.URL)
		t.Log("In production code, you would inject this URL into the HTTP client")

		// Para este ejemplo, solo verificamos que el mock funciona
		resp, err := http.Get(server.URL + "/items/MLM123456")
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestIntegration_PerformanceBaseline(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping performance test in short mode")
	}

	ctx := context.Background()
	const maxDuration = 2 * time.Second // SLA baseline

	t.Run("item_retrieval_performance", func(t *testing.T) {
		start := time.Now()

		_, err := items.GetByID(ctx, testenv.TestItemID, testenv.AccessToken)
		duration := time.Since(start)

		require.NoError(t, err)
		assert.Less(t, duration, maxDuration, "Item retrieval exceeded SLA of %v", maxDuration)

		t.Logf("Item retrieval took: %v (SLA: %v)", duration, maxDuration)
	})
}
