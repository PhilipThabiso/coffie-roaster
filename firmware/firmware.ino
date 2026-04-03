#include <Adafruit_MAX31865.h>
#include <Adafruit_GFX.h>
#include <Adafruit_SSD1306.h>
#include <Wire.h>
#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>

// WiFi Settings
const char* ssid = "EWII-20B0-2.4GHz";
const char* password = "GCJMXUCD3DIGV2";
const char* serverUrl = "http://192.168.1.248:8080/log";

// OLED setup
#define SCREEN_WIDTH 128
#define SCREEN_HEIGHT 32
#define OLED_RESET -1
#define OLED_ADDR 0x3C
Adafruit_SSD1306 display(SCREEN_WIDTH, SCREEN_HEIGHT, &Wire, OLED_RESET);

// Thermo setup
Adafruit_MAX31865 thermo(16);
#define TEMP_OFFSET 5.0
#define RREF      427.0
#define RNOMINAL  100.0

void setup() {
  Serial.begin(115200);

  // WiFi Setup
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  Serial.println("\nConnected to WiFi");

  // OLED  Setup
  Wire.begin(4,5);
  display.begin(SSD1306_SWITCHCAPVCC, OLED_ADDR);
  display.clearDisplay();
  display.setTextSize(1);
  display.setTextColor(SSD1306_WHITE);
  
  // Thermo setup
  thermo.begin(MAX31865_2WIRE);
}

void loop() {
  float temp = thermo.temperature(RNOMINAL, RREF) - TEMP_OFFSET;

  // Print to OLED
  display.clearDisplay();
  display.setCursor(0, 0); 
  display.print("Temp: "); 
  display.print(temp);
  display.println(" C");
  display.display();

  // Send to Go Backend
  if (WiFi.status() == WL_CONNECTED) {
    WiFiClient client;
    HTTPClient http;

    http.begin(client, serverUrl);
    http.addHeader("Content-Type", "application/json");

    // Construct simple JSON string
    String httpRequestData = "{\"temperature\":" + String(temp) + "}";
    
    int httpResponseCode = http.POST(httpRequestData);
    
    Serial.print("HTTP Response code: ");
    Serial.println(httpResponseCode);
    
    http.end();
  }

  delay(5000); // Send data every 5 seconds
}
