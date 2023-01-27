#include <LiquidCrystal.h>

const int rs = 33, en = 25, d4 = 26, d5 = 27, d6 = 14, d7 = 12;
LiquidCrystal lcd(rs, en, d4, d5, d6, d7);

void setup() {
  // put your setup code here, to run once:
  lcd.begin(16, 2);
  lcd.print("Testing");
}

void loop() {
  // put your main code here, to run repeatedly:
  
}
