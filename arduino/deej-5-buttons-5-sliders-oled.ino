#include <Arduino.h>
#define BUTTON_SUPPORT     // Uncomment in order to enable button support using a keypad (keypad library is required)
#define BURN_IN_PROTECTION // Uncomment in order to enable burn-in protection for OLED displays

#define INVERT_PERIOD_MS 20000  // Invert colors every 20 seconds
#define INVERT_DURATION_MS 3000 // Invert colors for 5 seconds
#define SERIAL_TIMEOUT_MS 70000 // Timeout for serial communication
unsigned long lastInvertTime = 0;
unsigned long lastSerialRxTime = 0;
bool isInverted = false;
bool isInInvertableScreen = false;

// deej slider configuration
const int NUM_SLIDERS = 5;
const int analogInputs[NUM_SLIDERS] = {A0, A1, A2, A3, A6}; // A4 A5 are used for i2c
int analogSliderValues[NUM_SLIDERS];

#ifdef BUTTON_SUPPORT
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
#endif

// a ssd1306 128x32 oled display is used
// oled configuration
#include <Wire.h>
#include <Adafruit_GFX.h>
#include <Adafruit_SSD1306.h>
#define SCREEN_WIDTH 128
#define SCREEN_HEIGHT 32
#define I2COLED 0x3C
#define OLED_RESET -1
Adafruit_SSD1306 display(SCREEN_WIDTH, SCREEN_HEIGHT, &Wire, OLED_RESET);
// oled buffer sizes
#define WORD_BUFFER 22 + 1 // save 22 bytes for the word + 1 for the null terminator
#define VOLUME_BUFFER 4    // save 4 bytes for the volume (100 is the max volume + 1 for the null terminator)
char incomingData[WORD_BUFFER];

void updateSliderValues()
{
  for (int i = 0; i < NUM_SLIDERS; i++)
  {
    analogSliderValues[i] = analogRead(analogInputs[i]);
  }
}

void updateButtonValues()
{
#ifdef BUTTON_SUPPORT
  char customKey = customKeypad.getKey();
  if (customKey)
  {
    lastButton = customKey - '0';
    timeLastButtonPressed = millis();
  }
#endif
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
#ifdef BUTTON_SUPPORT
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
#endif

  Serial.println();
}

void showClockScreen()
{
  display.clearDisplay();
  display.setTextSize(4);
  display.setTextColor(SSD1306_WHITE);
  display.setCursor(8, 0);
  display.println(incomingData);
  display.display();
}

void burnInProtection()
{
#ifdef BURN_IN_PROTECTION
  // if the invert period has passed, turn off the screen to protect the OLED
  if (millis() - lastSerialRxTime > SERIAL_TIMEOUT_MS)
  {
    // turn off screen
    display.clearDisplay();
    display.display();
    isInInvertableScreen = false;
  }
  // every 30 seconds invert the colors for 1 second
  if (millis() - lastInvertTime > INVERT_PERIOD_MS)
  {
    lastInvertTime = millis();
    isInverted = true;
  }
  if (isInverted && isInInvertableScreen && millis() - lastInvertTime < INVERT_DURATION_MS)
  {
    display.invertDisplay(true);
  }
  else
  {
    display.invertDisplay(false);
  }
#endif
}

void showAppNameScreen(int volume)
{
  display.clearDisplay();

  // Draw the text
  display.setTextSize(2);
  display.setTextColor(SSD1306_WHITE);
  display.setCursor(0, 0);
  display.println(incomingData);

  int volumeWidth = map(volume, 0, 100, 0, SCREEN_WIDTH); // Adjusting width according to percentage
  const int topx = 0;
  const int topy = 22;
  const int diff_end = 30;
  display.drawRect(topx, topy, SCREEN_WIDTH - diff_end, 10, SSD1306_WHITE);
  display.fillRect(topx + 2, topy + 2, volumeWidth - diff_end - 4, 6, SSD1306_WHITE);

  // display the volume percentage next to the volume bar
  display.setTextSize(1);
  display.setCursor(SCREEN_WIDTH - diff_end, topy + 1);
  if (volume < 10)
    display.print(" ");
  if (volume < 100)
    display.print(" ");
  display.print(volume);
  display.print("%");
  display.display();
}

void setup()
{
  Serial.begin(9600);
  if (!display.begin(SSD1306_SWITCHCAPVCC, I2COLED))
  {
    Serial.println(F("SSD1306 allocation failed"));
  }

  display.setTextWrap(false);
  display.clearDisplay();
  display.display();
}

void loop()
{
  if (Serial.available() > 0)
  {
    char firstByte = Serial.read();
    lastSerialRxTime = millis();
    Serial.println(firstByte);
    switch (firstByte)
    {
    case '>':
    {
      Serial.println("ShowAppScreen");
      Serial.readBytesUntil(':', incomingData, VOLUME_BUFFER);
      incomingData[VOLUME_BUFFER] = '\0';
      uint8_t volume = atoi(incomingData);
      uint8_t readBytes = Serial.readBytesUntil('\n', incomingData, WORD_BUFFER);
      incomingData[readBytes] = '\0';
      Serial.print("App:");
      Serial.println(incomingData);
      Serial.print("Volume:");
      Serial.println(volume);
      showAppNameScreen(volume);
      isInInvertableScreen = false;
      break;
    }
    case '<':
    {
      uint8_t readBytes = Serial.readBytesUntil('\n', incomingData, WORD_BUFFER);
      incomingData[readBytes] = '\0';
      showClockScreen();
      isInInvertableScreen = true;
      break;
    }
    default:
      break;
    }
  }

  // communicate with the deej
  updateSliderValues();
  updateButtonValues();
  sendSliderValues();
  burnInProtection();
}
