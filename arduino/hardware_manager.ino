// arduino json
#include <WiFi.h>
#include <HTTPClient.h>
#include <Keypad.h>
#include <LiquidCrystal.h>

#define BUTTON_PIN 2
#define LED_PIN 3
//Keypad vars
const int ROW_NUM = 4; //four rows
const int COLUMN_NUM = 4; //four columns

char keys[ROW_NUM][COLUMN_NUM] = {
  {'1','2','3', 'A'},
  {'4','5','6', 'B'},
  {'7','8','9', 'C'},
  {'*','0','#', 'D'}
};

byte pin_rows[ROW_NUM] = {36, 34, 32, 30}; //connect to the row pinouts of the keypad
byte pin_column[COLUMN_NUM] = {28, 26, 24, 22}; //connect to the column pinouts of the keypad

Keypad keypad = Keypad( makeKeymap(keys), pin_rows, pin_column, ROW_NUM, COLUMN_NUM );

//LCD vars
const int rs = 12, en = 11, d4 = 5, d5 = 4, d6 = 3, d7 = 2;
LiquidCrystal lcd(rs, en, d4, d5, d6, d7);

const char* PASSWORD = "nosoup4u";
const char* URL = "http://:8000";
const int LENGTH = 6;

String keypad_buffer;

HTTPClient arduino_client;
WiFiClient wifi_client;

void setup() {
    keypad_buffer = "";
    // light protoype with LED
    setMode(LED_PIN, OUTPUT);
    setMode(BUTTON_PIN, INPUT);

    arduino_client.begin()

    Serial.begin(9600); 
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
    char key = keypad.getKey();
    String lcd_input = "";
    if (key) {
        if key == 'D' && keypad_buffer.length() > 0 {
            keypad_buffer.remove(keypad_buffer.length() - 1);
            lcd.print(lcd_input.remove(lcd_input.length() - 1));
        } else {
            keypad_buffer.concat(key);
            lcd.print(lcd_input.concat("*"));

            if(keypad_buffer.length() == 6) {
                WiFiClient client;
                HTTPClient http;

                http.begin(client, URL);
                
                http.addHeader("Content-Type", "application/x-www-form-urlencoded");
                
                String httpRequestDataJSON = "{\"code\": " + keypad_buffer + "\"}";;           

                int httpResponseCode = http.POST(httpRequestDataJSON);
                
                Serial.print("HTTP Response code: ");
                Serial.println(httpResponseCode);
                
                // Free resources
                http.end();

                lcd.clear();
            }
        }
    }

}