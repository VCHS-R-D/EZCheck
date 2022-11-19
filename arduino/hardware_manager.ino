// arduino json
#include <WiFi.h>
#include <HTTPClient.h>
#include <ArduinoJson.h>

#define BUTTON_PIN 2
#define LED_PIN 3

const char* SSID = "REPLACE_WITH_YOUR_SSID";
const char* PASSWORD = "REPLACE_WITH_YOUR_PASSWORD";
const char* URL = "http://192.168.1.106:8000";

HTTPClient arduino_client;
WiFiClient wifi_client;

void setup() {
    // light protoype with LED
    setMode(LED_PIN, OUTPUT);
    setMode(BUTTON_PIN, INPUT);
    arduino_client.begin()

    Serial.begin(115200); 
    try {
        WiFi.begin(ssid, password);
        Serial.println("Connecting");
        if (Wifi.status == WL_CONNECT_FAILED) throw("Connect Failed");
    } catch (string error) {
        Serial.println("Connection issue", error);
        while(1);
    }

    Serial.print("Connected to WiFi network with IP Address: ");
    Serial.println(WiFi.localIP());
    
    Serial.println("Timer set to 5 seconds (timerDelay variable), it will take 5 seconds before publishing the first reading.");
}

void loop() {
    if ((millis() - lastTime) > timerDelay) {

      WiFiClient client;
      HTTPClient http;
    
      http.begin(client, URL);
      
      http.addHeader("Content-Type", "application/x-www-form-urlencoded");
      
      char* httpRequestData = "{/"code/": /"000000/"}";           

      int httpResponseCode = http.POST(httpRequestData);
      
      Serial.print("HTTP Response code: ");
      Serial.println(httpResponseCode);
        
      // Free resources
      http.end();
    }

}