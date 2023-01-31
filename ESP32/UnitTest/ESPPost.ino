#include <WiFi.h>
#include <HTTPClient.h>

//pass
const char* ssid = "";
const char* password = "";

//path
String server = "http://192.168.1.208:8080";
String authPath = server+"/auth";

//id
String machineID = "laser_cutter";

void setup() {
  Serial.begin(115200); 

  WiFi.begin(ssid, password);
  Serial.println("Connecting");
  while(WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  Serial.println("");
  Serial.print("Connected to WiFi network with IP Address: ");
  Serial.println(WiFi.localIP());
}

void loop() {
  if (Serial.available()) {
    //Check WiFi connection status
    if(WiFi.status()== WL_CONNECTED){
      //get entered pin
      String pin = Serial.readString();
      //confirmation
      Serial.println("Sending pin: "+pin+" to machine: "+machineID);
      //send post request
      WiFiClient client;
      HTTPClient http;
      http.begin(client, authPath);
      http.addHeader("Content-Type", "application/json");
      int httpResponseCode = http.POST("{\"code\":\""+pin+"\",\"machineID\":\""+machineID+"\"}");
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
    }
    else {
      Serial.println("WiFi Disconnected");
    }
  }
}
