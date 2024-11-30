package score

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kaz/pprotein/internal/storage"
	"github.com/labstack/echo/v4"
)

type handler struct {
	store storage.Storage
}

func NewHandler(store storage.Storage) *handler {
	return &handler{
		store: store,
	}
}

type score struct {
	GroupID string `json:"groupId"`
	Score   int64  `json:"score"`
}

func (h *handler) GetByID(c echo.Context) error {
	id := c.Param("groupid")
	b, err := h.store.Get("score", id)
	if err != nil {
		return fmt.Errorf("failed get score from store: %w", err)
	}
	n, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return fmt.Errorf("failed Parse score: %w", err)
	}
	return c.JSON(http.StatusOK, score{
		GroupID: id,
		Score:   n,
	})
}

func (h *handler) GetAll(c echo.Context) error {
	bm, err := h.store.GetAllMap("score")
	if err != nil {
		return fmt.Errorf("failed get scores from store: %w", err)
	}
	scores := make([]score, 0, len(bm))
	for k, v := range bm {
		n, _ := strconv.ParseInt(string(v), 10, 64)
		scores = append(scores, score{GroupID: k, Score: n})
	}
	return c.JSON(http.StatusOK, scores)
}

func (h *handler) Put(c echo.Context) error {
	id := c.Param("groupid")

	var score score
	if err := c.Bind(&score); err != nil {
		return fmt.Errorf("failed Bind: %w", err)
	}

	if err := h.store.Put("score", id, []byte(fmt.Sprint(score.Score))); err != nil {
		return fmt.Errorf("failed Put to store: %w", err)
	}
	return c.NoContent(http.StatusOK)
}
