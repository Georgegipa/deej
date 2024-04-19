#include <Arduino.h>

// deej slider configuration
const int NUM_SLIDERS = 5;
const int analogInputs[NUM_SLIDERS] = {A0, A1, A2, A3, A6}; // A4 A5 are used for i2c
int analogSliderValues[NUM_SLIDERS];

#include <Keypad.h>
// for button configuration a keypad is used
#define ROWS 1
#define COLS 5
// define the symbols on the buttons of the keypads
char hexaKeys[ROWS][COLS] = {
    {'1', '2', '3', '4', '5'}};

byte rowPins[ROWS] = {7};             // connect to the row pinouts of the keypad
byte colPins[COLS] = {6, 5, 4, 3, 2}; // connect to the column pinouts of the keypad
#define DEBUNCE_TIME 50               // time between each button press
// keypad configuration
Keypad customKeypad = Keypad(makeKeymap(hexaKeys), rowPins, colPins, ROWS, COLS);
// last pressed button
uint8_t lastButton = 0;
long timeLastButtonPressed = 0;

void updateSliderValues()
{
  for (int i = 0; i < NUM_SLIDERS; i++)
  {
    analogSliderValues[i] = analogRead(analogInputs[i]);
  }
}

void updateButtonValues()
{
  char customKey = customKeypad.getKey();
  if (customKey)
  {
    lastButton = customKey - '0';
    timeLastButtonPressed = millis();
  }
}

void sendSliderValues()
{
  for (int i = 0; i < NUM_SLIDERS; i++)
  {
    Serial.print("s");
    Serial.print(analogSliderValues[i]);
    if (i < NUM_SLIDERS - 1)
    {
      Serial.print("|");
    }
  }
  Serial.print("|");
  // the following code is used by the Miodec/deej
  for (int i = 0; i < COLS; i++)
  {
    Serial.print("b");
    // check if this was the last button pressed
    if (lastButton - 1 == i && millis() - timeLastButtonPressed < DEBUNCE_TIME)
    {
      Serial.print("1");
    }
    else
    {
      Serial.print("0");
    }
    if (i < COLS - 1)
    {
      Serial.print("|");
    }
  }
  Serial.println();
}

void setup()
{
  Serial.begin(9600);
}

void loop()
{
  // communicate with the deej
  updateSliderValues();
  updateButtonValues();
  sendSliderValues();
}
