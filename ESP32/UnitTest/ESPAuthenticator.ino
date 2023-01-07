#include <Keypad.h>

#include <LiquidCrystal.h>

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

Keypad keypad = Keypad( makeKeymap(keys), pin_rows, pin_column, ROW_NUM, COLUMN_NUM);


//LCD vars
const int rs = 33, en = 25, d4 = 26, d5 = 27, d6 = 14, d7 = 12;
LiquidCrystal lcd(rs, en, d4, d5, d6, d7);

//Pass vars
String pass = "";

void setup(){
  Serial.begin(9600);
  lcd.begin(16, 2);
  lcd.print("Welcome!");
}

void loop(){
  //handle keypress
  char key = keypad.getKey();
  if (key){
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
      if (pass.length()==6) lcd.print("Hello Mr. Huber");
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
