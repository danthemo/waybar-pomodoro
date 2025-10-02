package timer

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ToggleHandler(w http.ResponseWriter, r *http.Request) {
	Pomo.Pause = !Pomo.Pause
}

func ResetHandler(w http.ResponseWriter, r *http.Request) {
	Pomo.Pause = true
	Pomo.IsWork = true
	Pomo.CurrTime = work
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	minutes := int(Pomo.CurrTime.Minutes())
	seconds := int(Pomo.CurrTime.Seconds()) % 60

	response := map[string]string{
		"text": fmt.Sprintf("🍅 %02d:%02d", minutes, seconds),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
