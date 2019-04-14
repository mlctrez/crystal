#include <ESP8266WiFi.h>
#include <WiFiClient.h>
#include <ESP8266WiFi.h>

const char* ssid = "xxxx";
const char* password = "xxxx";

WiFiServer wifiServer(8266);

void setup() {

  pinMode(D4, OUTPUT); // D4 is the led
  digitalWrite(D4, HIGH); // turn on the led

  pinMode(D1, OUTPUT);
  digitalWrite(D1, HIGH); // data
  pinMode(D2, OUTPUT);
  digitalWrite(D2, HIGH); // chip select
  pinMode(D3, OUTPUT);
  digitalWrite(D3, HIGH); // clock

  delay(1000);
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    writeData('.');
  }
  writeData(4);
  writeData(12);
  writeData('O');
  writeData('K');

  wifiServer.begin();

}

void loop() {

  WiFiClient client = wifiServer.available();
  if (client) {
    digitalWrite(D4, LOW); // led on for client connected
    while (client.connected()) {
      while (client.available() > 0) {
        writeData(client.read());
      }
    }
    client.stop();
    digitalWrite(D4, HIGH); // led off for client disconnected
  }
}

void writeData(byte data) {
  int timing = 2;
  digitalWrite(D2, LOW);  delay(timing); // select on
  for ( int b = 7; b >= 0; b--) {
    digitalWrite(D1, (bitRead(data, b)) == 1 ? HIGH : LOW ); delay(timing); // data bit
    digitalWrite(D3, LOW); delay(timing); // clock low
    digitalWrite(D3, HIGH); delay(timing); // clock high
  }
  digitalWrite(D2, HIGH);  delay(timing); // select off
}