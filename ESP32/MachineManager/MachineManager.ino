#include <ArduinoJson.h>
#include <Keypad.h>
#include <LiquidCrystal.h>
#include <WiFi.h>
#include <HTTPClient.h>
//HTTP vars
const char* ssid = "***REMOVED***";
const char* password = "***REMOVED***";

const String machineID = "machine1";
const String k = "http://10.124.6.136:8080";
const String path = k+"/auth";
const String signout = k+"/signout";
bool signedIn = false;
String curr_name = "";

//Keypad vars
const int ROW_NUM = 4; //four rows
const int COLUMN_NUM = 4; //four columns

char keys[ROW_NUM][COLUMN_NUM] = {
  {'1','2','3', 'A'},
  {'4','5','6', 'B'},
  {'7','8','9', 'C'},
  {'*','0','#', 'D'}
};

byte pin_rows[ROW_NUM] = {17, 5, 18, 19}; //connect to the row pinouts of the keypad {R1, R2, R3, R4}
byte pin_column[COLUMN_NUM] = {15, 2, 4, 16}; //connect to the column pinouts of the keypad {C1, C2, C3, C4}

Keypad keypad = Keypad( makeKeymap(keys), pin_rows, pin_column, ROW_NUM, COLUMN_NUM);

//LCD vars
const int rs = 33, en = 25, d4 = 26, d5 = 27, d6 = 14, d7 = 32;
LiquidCrystal lcd(rs, en, d4, d5, d6, d7);

//LED vars
const int led  =  21; 

//Pass vars
String pass = "";

void setup(){
  //activate
  Serial.begin(9600);
  WiFi.begin(ssid, password);
  lcd.begin(16, 2);
  pinMode(led,OUTPUT); 
  
  //setup wifi
  lcd.print("Connecting");
  while(WiFi.status() != WL_CONNECTED) {
    delay(500);
    lcd.print(".");
  }
  lcd.clear();
  lcd.print("Connected");
  delay(1000);
  
  //init state
  lcd.clear();
  lcd.print("Welcome!");
  digitalWrite(led,LOW);
}

void loop(){
  //handle keypress
  char key = keypad.getKey();
  if (key){
    //# -> sign out
    if (signedIn){
      if (key=='#') {
        if (machineSignout()) {
          digitalWrite(led, LOW);
          lcd.clear();
          lcd.print("Signed Out");
          signedIn = !signedIn;
        }
      }
    }
    else {
      //D -> cut pass
      if (key=='D') {
        pass.remove(pass.length()-1);
        //display
        lcd.clear();
        lcd.print(pass);
      }
      //* -> submit pass
      else if (key=='*') {
        lcd.clear();
        //make request
        boolean res = validate(pass);
        // correct pass
        if (res) {
          lcd.clear();
          String message = "Hello ";
          message.concat(curr_name);
          lcd.print(message);
          digitalWrite(led,HIGH);
          signedIn = true;
        }
        //wrong pass
        else lcd.print("Denied Access");
        pass="";
      }
      //else -> add char to pass
      else {
        pass.concat(key);
        //display
        lcd.clear();
        lcd.print(pass);
        Serial.println(pass);
      }
    }
  }
}

boolean validate(String pass){
  //connected 
  if(WiFi.status() == WL_CONNECTED){
    //confirmation
    lcd.clear();
    lcd.print("Sending pin");
    //send post request
    WiFiClient client;
    HTTPClient http;
    http.begin(client, path);
    http.addHeader("Content-Type", "application/json");
    int httpResponseCode = http.POST("{\"code\":\"" + pass + "\",\"machineID\":\"" + machineID+"\"}");
    //valid request
    String response = http.getString();
    DynamicJsonDocument response_data(1024);
    deserializeJson(response_data, response);
    bool auth = response_data["authorized"];
    
    if (auth) {
      String temp_name = response_data["name"];
      curr_name = temp_name;
      http.end();
      pass = "";
      return true;
    }
    else {
      lcd.clear();
      lcd.print("Error: ");
      lcd.print(httpResponseCode);
      Serial.println(httpResponseCode);
      curr_name = "";
      http.end();
      pass = "";
      return false;
    }
    // Free resources
  }
  return false;
}

boolean machineSignout(){
  //connected 
  if(WiFi.status() == WL_CONNECTED){
    //send post request
    WiFiClient client;
    HTTPClient http;
    http.begin(client, signout);
    http.addHeader("Content-Type", "application/json");
    int httpResponseCode = http.POST("{\"name\":\"" + curr_name + "\",\"machineID\":\"" + machineID+"\"}");
    
    if (httpResponseCode == 200) {
      return true;
    }
    else {
      lcd.clear();
      lcd.print("Error Signing Out: ");
      lcd.print(httpResponseCode);
      Serial.println(httpResponseCode);
      curr_name = "";
      http.end();
      pass = "";
      return false;
    }
    // Free resources
  }
  return false;
}
