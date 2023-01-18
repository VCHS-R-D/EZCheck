#include <Keypad.h>

const int ROW_NUM = 4; //four rows
const int COLUMN_NUM = 4; //four columns

char keys[ROW_NUM][COLUMN_NUM] = {
  {'1','2','3', 'A'},
  {'4','5','6', 'B'},
  {'7','8','9', 'C'},
  {'*','0','#', 'D'}
};

byte pin_rows[ROW_NUM] = {15, 2, 4, 16}; //connect to the row pinouts of the keypad
byte pin_column[COLUMN_NUM] = {17, 5, 18, 19}; //connect to the column pinouts of the keypad

Keypad keypad = Keypad( makeKeymap(keys), pin_rows, pin_column, ROW_NUM, COLUMN_NUM );

const int len = 6;
String pass = "";

void setup(){
  Serial.begin(9600);
}

void loop(){
  char key = keypad.getKey();
  if (key){
    if (key=='D'){
      Serial.println("passcode was reset");
      pass="";
    } else {
      pass.concat(key);
      if (pass.length()==6) {
        Serial.println("The passcode you entered was: "+pass);
        pass="";
      } else {
        Serial.println(pass);
      }
    }
  }
}
