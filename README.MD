# waybar-pomodoro

**waybar-pomodoro** — это простой таймер Pomodoro, реализованный на Go, который запускается как HTTP-сервер и позволяет управлять таймером (старт/пауза, сброс) и получать статус в формате JSON. Идеально подходит для интеграции с [waybar](https://github.com/Alexays/Waybar) в окружении Hyprland.

---

## Особенности

- Таймер Помодоро с классическими интервалами: 25 минут работы и 5 минут перерыва.
- Управление через HTTP-запросы:
  - `/toggle` — старт/пауза таймера.
  - `/reset` — сброс таймера.
  - `/status` — получение текущего времени таймера в JSON-формате.
- Отправка уведомлений через `notify-send` и звуковое оповещение.
- Легкая интеграция с waybar и другими статус-барами.

---

## Требования

- Go 1.18+
- Linux с установленными `notify-send` и `paplay` (для уведомлений и звуков)
- Waybar
- Hyprland (рекомендуется, но не обязательно)

---

## Установка

1. Клонируйте репозиторий:

   ```bash
   git clone https://github.com/yourusername/waybar-pomodoro.git
   cd waybar-pomodoro
   ```

2. Соберите проект или скачайте из раздела Releases:

    ```bash
    go build -o waybar-pomodoro
    ```

3. Запустите процесс из консоли, или в фоновом режиме (демон):
    
    ```bash
    ./waybar-pomodoro
    ```

## Использование с waybar

Добавьте в конфиг waybar блок с командой, которая будет запрашивать статус таймера:
    `~/.config/waybar/config.jsonc`


    ```bash
    "custom/pomodoro": {
        "exec": "curl http://localhost:8081/status",
        "interval": 1,
        "on-click": "curl http://localhost:8081/toggle",
        "on-click-right": "curl http://localhost:8081/reset",
        "return-type": "json",
    }
    ```

- Нажатие левой кнопкой мыши - старт/стоп таймера

- Нажатие правой кнопкой мыши - сброс таймера
