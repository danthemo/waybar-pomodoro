package main

import (
	"fmt"
	"net/http"

	"github.com/danthemo/waybar-pomodoro/timer"
)

func main() {
	go timer.Start()

	mux := http.NewServeMux()
	mux.HandleFunc("/toggle", timer.ToggleHandler)
	mux.HandleFunc("/reset", timer.ResetHandler)
	mux.HandleFunc("/status", timer.StatusHandler)

	fmt.Println("Started on http://localhost:8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		fmt.Println("Ошибка выполнения:", err)
	}
}
