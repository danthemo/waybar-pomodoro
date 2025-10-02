package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Pomodoro struct {
	Pause    bool          `json:"pause"`
	IsWork   bool          `json:"is_work"`
	CurrTime time.Duration `json:"curr_time"`
}

var work time.Duration = 25 * time.Minute
var chill time.Duration = 5 * time.Minute

var pomo = Pomodoro{
	Pause:    true,
	IsWork:   true,
	CurrTime: work,
}

func ToggleHandler(w http.ResponseWriter, r *http.Request) {
	pomo.Pause = !pomo.Pause
}

func ResetHandler(w http.ResponseWriter, r *http.Request) {
	pomo.Pause = true
	pomo.IsWork = true
	pomo.CurrTime = work
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	minutes := int(pomo.CurrTime.Minutes())
	seconds := int(pomo.CurrTime.Seconds()) % 60

	response := map[string]string{
		"text": fmt.Sprintf("🍅 %02d:%02d", minutes, seconds),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func StartTimer() {
	ticker := time.NewTicker(1 * time.Second) // Исправлено: тикер вместо sleep
	defer ticker.Stop()

	for range ticker.C {
		if !pomo.Pause {
			if pomo.CurrTime > 0 {
				pomo.CurrTime -= time.Second
			}

			// Таймер достиг нуля
			if pomo.CurrTime <= 0 {
				if pomo.IsWork {
					// Работа → Перерыв
					pomo.IsWork = false
					pomo.CurrTime = chill
					pomo.Pause = true // Пауза после работы
					fmt.Println("🎉 Работа завершена! Время перерыва")
				} else {
					// Перерыв → Работа
					pomo.IsWork = true
					pomo.CurrTime = work
					pomo.Pause = true // Пауза после перерыва
					fmt.Println("💪 Перерыв окончен! Время работать")
				}
			}
		}
	}
}

func main() {
	go StartTimer()

	mux := http.NewServeMux()
	mux.HandleFunc("/toggle", ToggleHandler)
	mux.HandleFunc("/reset", ResetHandler)
	mux.HandleFunc("/status", StatusHandler)

	fmt.Println("Started on http://localhost:8081")
	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		fmt.Println("Ошибка выполнения")
	}
}
