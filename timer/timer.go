package timer

import (
	"fmt"
	"os/exec"
	"time"
)

type Pomodoro struct {
	Pause    bool
	IsWork   bool
	CurrTime time.Duration
}

var (
	work  = 25 * time.Minute
	chill = 5 * time.Minute

	Pomo = Pomodoro{
		Pause:    true,
		IsWork:   true,
		CurrTime: work,
	}

	soundPath = "/usr/share/sounds/freedesktop/stereo/complete.oga"
)

func Start() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if !Pomo.Pause {
			if Pomo.CurrTime > 0 {
				Pomo.CurrTime -= time.Second
			}

			if Pomo.CurrTime <= 0 {
				if Pomo.IsWork {
					Pomo.IsWork = false
					Pomo.CurrTime = chill
					fmt.Println("🎉 Работа завершена! Время перерыва")
					notifyUser("🍅 Помодоро", "🎉 Работа завершена! Время перерыва")
				} else {
					Pomo.IsWork = true
					Pomo.CurrTime = work
					Pomo.Pause = true
					fmt.Println("💪 Перерыв окончен! Время работать")
					notifyUser("🍅 Помодоро", "💪 Перерыв окончен! Время работать... Или отдыхаем?")
				}
			}
		}
	}
}

func notifyUser(title, message string) {
	if err := exec.Command("notify-send", title, message).Run(); err != nil {
		fmt.Println("❌ Не удалось отправить уведомление:", err)
	}

	if err := exec.Command("paplay", soundPath).Run(); err != nil {
		fmt.Println("❌ Не удалось воспроизвести звук:", err)
	}
}
