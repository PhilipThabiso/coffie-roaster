coffie-roaster/
├── backend/
│   ├── main.go          # Entry point and HTTP routes
│   ├── handlers.go      # HTTP logic for receiving data
│   ├── database.go      # SQLite initialization and queries
│   ├── data.db          # Generated SQLite file
│   └── go.mod           # Go module file
├── firmware/
│   └── sensor.ino       # ESP8266 Arduino/C++ code
└── schema.sql           # Initial SQL table definitions

arduino-cli monitor -p /dev/ttyUSB0 -c baudrate=115200

arduino-cli compile --upload -p /dev/ttyUSB0 \
  --fqbn esp8266:esp8266:generic ~/proj/coffie-roaster/firmware

arduino-cli lib install "Adafruit MAX31865 library"
