package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/your-username/project/internal/models"
	"github.com/your-username/project/internal/repository"
)

type Handler struct {
	repo *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{repo: repo}
}

// CreateSubscription godoc
// @Summary      Создать подписку
// @Description  Сохраняет новую запись о подписке в базу данных
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        input body models.Subscription true "Данные подписки"
// @Success      201 {object} models.Subscription
// @Failure      400 {object} map[string]string
// @Router       /subscriptions [post]
func (h *Handler) CreateSubscription(c *gin.Context) {
	var input models.Subscription
	
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("[ERROR] Failed to parse request body: %v", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request body"})
		return
	}

	log.Printf("[INFO] Creating subscription for user: %s, service: %s", input.UserID, input.ServiceName)

	if err := h.repo.Create(input); err != nil {
		log.Printf("[ERROR] Database error during creation: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	log.Printf("[INFO] Subscription created successfully for user: %s", input.UserID)
	c.JSON(http.StatusCreated, input)
}

// GetTotalCost godoc
// @Summary      Подсчитать общую стоимость
// @Description  Агрегирует стоимость подписок пользователя с учетом фильтров
// @Tags         subscriptions
// @Produce      json
// @Param        user_id query string true "UUID пользователя"
// @Param        service_name query string false "Название сервиса"
// @Success      200 {object} models.TotalCostResponse
// @Failure      400 {object} map[string]string
// @Router       /stats [get]
func (h *Handler) GetTotalCost(c *gin.Context) {
	userIDStr := c.Query("user_id")
	serviceName := c.Query("service_name")

	if userIDStr == "" {
		log.Println("[WARN] Missing user_id in request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		log.Printf("[WARN] Invalid UUID format: %s", userIDStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id format"})
		return
	}

	log.Printf("[INFO] Calculating total cost for user: %s (service: %s)", userID, serviceName)

	total, err := h.repo.GetTotalCost(userID, serviceName)
	if err != nil {
		log.Printf("[ERROR] Failed to get stats from DB: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, models.TotalCostResponse{TotalCost: total})
}