package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"time"
)

func MakeServer() http.Handler {
	//mux := echo.New()
	//mux.HTTPErrorHandler = func(err error, c echo.Context) {
	//	scope.GetLogger().Error("error: ", slog.Any("error", err))
	//	var e echo.HTTPError
	//	if !errors.Is(err, &e) {
	//		c.Error(echo.ErrInternalServerError.WithInternal(err))
	//	}
	//	c.Error(err)
	//}
	//mux.GET("/hello", HandleHello())
	//mux.POST("/chat", HandleChatCompletions())

	mux := http.NewServeMux()

	mux.HandleFunc("/events", eventsHandler)
	mux.Handle("/", http.FileServer(http.Dir("./static")))
	return mux
}

func HandleHello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "hello",
		})
	}
}

func HandleChatCompletions() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")

	text := `
	As the sun dipped below the horizon, the city skyline transformed into a sea of glowing lights, each window a story, each street a vein pulsing with life. The air was thick with the scent of rain-soaked pavement, mingling with the aroma of street food wafting from hidden alleys. In the distance, the hum of distant conversations blended with the rhythmic beat of a street musician’s guitar, creating a symphony that only the night could compose.

	Among the crowd, a lone figure moved with purpose, their footsteps echoing softly against the cobblestones. Clad in a worn leather jacket, their eyes scanned the bustling streets, searching for something—or perhaps someone. The city's energy coursed through them, a silent companion on their nocturnal quest. As they turned a corner, a flash of neon light reflected in their eyes, revealing a hidden smile.

	This was a place where secrets whispered through the cracks in the pavement, where dreams danced in the shadows, waiting to be claimed by those brave enough to chase them. Here, in the heart of the night, anything was possible.
	`

	words := strings.Split(text, " ")
	for i, word := range words {

		fmt.Fprintf(w, "data: %s\n\n", word)
		if i%15 == 0 {
			fmt.Fprintf(w, "\n\n\n")
		}
		time.Sleep(50 * time.Millisecond)
		w.(http.Flusher).Flush()
	}

	doneChan := r.Context().Done()
	<-doneChan

}
