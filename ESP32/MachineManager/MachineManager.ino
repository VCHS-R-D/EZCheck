#include <Keypad.h>
#include <LiquidCrystal.h>
#include <WiFi.h>
#include <HTTPClient.h>

//HTTP vars
const char* ssid = "***REMOVED***";
const char* password = "***REMOVED***";
String machineID = "machine1";
String path = "http://10.124.5.23:8080/auth";

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
    digitalWrite(led,LOW);
    //D -> remove last char
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
        lcd.print("Hello Mr. Huber");
        digitalWrite(led,HIGH);
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

boolean validate(String pass){
  boolean ret = false;
  //connected
  if(WiFi.status() == WL_CONNECTED){
    //confirmation
    lcd.clear();
    lcd.print("Sending pin: "+pass);
    //send post request
    WiFiClient client;
    HTTPClient http;
    http.begin(client, path);
    http.addHeader("Content-Type", "application/json");
    int httpResponseCode = http.POST("{\"code\":\""+pass+"\",\"machineID\":\""+machineID+"\"}");
    //valid request
    if (httpResponseCode>0) {
      ret = true;
    }
    else {
      lcd.clear();
      lcd.print("Error: ");
      lcd.print(httpResponseCode);
    }
    // Free resources
    http.end();
  }
  //disconnected
  else {
    lcd.clear();
    lcd.print("WiFi Disconnected");
  }
  return ret;
}
