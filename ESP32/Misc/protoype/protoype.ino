#include <WiFi.h>
#include <HTTPClient.h>
#include <Keypad.h>
#include <LiquidCrystal.h>

const char* ssid = "***REMOVED***";
const char* pass = "***REMOVED***";

//Keypad vars
const int ROW_NUM = 4; //four rows
const int COLUMN_NUM = 4; //four columns

char keys[ROW_NUM][COLUMN_NUM] = {
  {'1','2','3', 'A'},
  {'4','5','6', 'B'},
  {'7','8','9', 'C'},
  {'*','0','#', 'D'}
};

byte pin_rows[ROW_NUM] = {17, 5, 18, 19}; //connect to the row pinouts of the keypad
byte pin_column[COLUMN_NUM] = {15, 2, 4, 16}; //connect to the column pinouts of the keypad

Keypad keypad = Keypad( makeKeymap(keys), pin_rows, pin_column, ROW_NUM, COLUMN_NUM );

//LCD vars
const int rs = 33, en = 25, d4 = 26, d5 = 27, d6 =14, d7 = 32;
LiquidCrystal lcd(rs, en, d4, d5, d6, d7);
const int LENGTH = 6;

String keypad_buffer;

//server endpoints
String path = "http://10.124.5.23:8080/auth";
String machineID = "machine1"; // this needs to be hardcoded into each EZCheck module

void setup() {
    keypad_buffer = "";
    // light protoype with LED
    Serial.begin(115200);
    WiFi.begin(ssid, pass);
    Serial.println("Connecting");

    while(WiFi.status() != WL_CONNECTED) {}

    Serial.println("");
    Serial.print("Connected to WiFi network with IP Address: ");
    Serial.println(WiFi.localIP());
}

void loop() {
    char key = keypad.getKey();
    String lcd_input = "";

    if (key == 'D' && keypad_buffer.length() > 0) {
      keypad_buffer.remove(keypad_buffer.length() - 1);
    } else if (key) {  
        
      keypad_buffer.concat(key);
      lcd.print(keypad_buffer);

      if (keypad_buffer.length() == 6) {
        WiFiClient client;
        HTTPClient http;  
        http.begin(client, path);
        http.addHeader("Content-Type", "application/json");
  
        int httpResponseCode = http.POST("{\"code\":\""+keypad_buffer+"\",\"machineID\":\""+machineID+"\"}");
        //handle response
        if (httpResponseCode>0) {
            Serial.print("HTTP Response code: ");
            Serial.println(httpResponseCode);
            String payload = http.getString();
            Serial.println(payload);
        }
        else {
            Serial.print("Error code: ");
            Serial.println(httpResponseCode);
        }
        // Free resources
        http.end(); 
        keypad_buffer = "";
      }
    }
}
