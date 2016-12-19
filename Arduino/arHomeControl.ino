#include <SPI.h>
#include <Ethernet.h>

/* Network stuff */
EthernetServer server(80);

#define STRING_BUFFER_SIZE 128
char buffer[STRING_BUFFER_SIZE];

/* Relais stuff */
// 0 = OFF, 1 = ON
int lightState = 0;
int buttonState = 0;
int buttonPressed = 0;

uint8_t mac[6] = { 0x00, 0x01, 0x02, 0x03, 0x04, 0x05 };


void setup()
{
	Serial.begin(9600);
	Serial.println("Start Setup");
	// ********** PIN definition ***********
	Serial.println("begin PIN definition");
	// Output Relais
	pinMode(34, OUTPUT);
	pinMode(36, OUTPUT);
	pinMode(38, OUTPUT);
	pinMode(40, OUTPUT);
	pinMode(42, OUTPUT);
	pinMode(44, OUTPUT);
	pinMode(46, OUTPUT);
	pinMode(48, OUTPUT);


	// Input Push Button
	//pinMode(8, INPUT);
	// ******* END PIN definition **********

	// ******* Network definition **********
	Ethernet.begin(mac);
	server.begin();

	Serial.print("localIP: ");
	Serial.println(Ethernet.localIP());
	Serial.print("subnetMask: ");
	Serial.println(Ethernet.subnetMask());
	Serial.print("gatewayIP: ");
	Serial.println(Ethernet.gatewayIP());
	Serial.print("dnsServerIP: ");
	Serial.println(Ethernet.dnsServerIP());

}

void loop()
{
	//checkButtonState();

	checkNetworkState();
}

void checkButtonState(){
	buttonState = digitalRead(8);
	if (buttonState == HIGH) {
		if (buttonPressed == 0) {
			//toggleLight();
			buttonPressed = 1;
		}
	}
	else {
		buttonPressed = 0;
	}
}

void toggleLight(int pin)
{
	Serial.println(pin);
	if (digitalRead(pin) == 0)
	{
		Serial.println("Set to Hight");
		digitalWrite(pin, HIGH);
		delay(200);
	}
	else if (digitalRead(pin) == 1)
	{
		Serial.println("Set to Low");
		digitalWrite(pin, LOW);
		delay(200);
	}
}

void handleCommand(EthernetClient client, char* cmd, char* pinArray) {
	int pin;
	sscanf(pinArray, "%d", &pin);
	
	if (strcmp(cmd, "status") == 0) {
		if (pin == 0) {
			sendStatus(client);
		}
	}
	else if (strcmp(cmd, "toggle") == 0) {
		toggleLight(pin);
		sendStatus(client);
	}

	else {
		Serial.println("404");
		send404(client);
	}
}

void sendStatus(EthernetClient client) {
	// Send a standard http response header
	client.println("HTTP/1.1 200 OK");
	client.println("Content-Type: application/json");
	client.println("Connnection: close");
	client.println();
	client.println("{");
	// Output the value of each analog input pin
	client.print("\"LightStatus");
	client.print("\": ");
	client.print(lightState);
	client.println(",");
	client.println("\n}");
}

void send404(EthernetClient client) {
	client.println("HTTP/1.1 404 OK");
	client.println("Content-Type: text/html");
	client.println("Connnection: close");
	client.println("");
	client.println("<!DOCTYPE HTML>");
	client.println("<html><body>404: Arduino sais NOOOO!</body></html>");
}

void checkNetworkState(){
	//Serial.println("listen for netzwork clients");
	// listen for incoming clients
	EthernetClient client = server.available();
	if (client) {
		// an http request ends with a blank line
		boolean currentLineIsBlank = true;
		while (client.connected()) {
			if (client.available()) {
				char c;
				int bufindex = 0; // reset buffer
				buffer[0] = client.read();
				buffer[1] = client.read();
				bufindex = 2;
				// Read the first line to determin the request page
				while (buffer[bufindex - 2] != '\r' && buffer[bufindex - 1] != '\n') {
					// read full row and save it in buffer
					c = client.read();
					if (bufindex < STRING_BUFFER_SIZE) {
						buffer[bufindex] = c;
					}
					bufindex++;
				}
				// Clean buffer for next row
				bufindex = 0;
				// Parse the query string
				int nSegments = countSegments(buffer);
				char **pathsegments = parse(buffer);
				int i = 0;
				for (i = 0; i < nSegments; i++) {
					// Serial.println(pathsegments[i]);
				}
				if (c == '\n' && currentLineIsBlank) {
					/*Serial.println(pathsegments[0]);
					Serial.println(pathsegments[1]);*/
					handleCommand(client, pathsegments[0], pathsegments[1]);
					break;
				}
				if (c == '\n') {
					currentLineIsBlank = true;
				}
				else if (c != '\r') {
					currentLineIsBlank = false;
				}
			}
		}
		// Give the web browser time to receive the data
		delay(1);
		// Close the connection:
		client.stop();
		Serial.println("Client disonnected");
	}
}

int countSegments(char* str) {
	int p = 0;
	int count = 0;
	while (str[p] != '\0') {
		if (str[p] == '/') {
			count++;
		}
		p++;
	}
	// We don't want to count the / in 'HTTP/1.1'
	count--;
	return count;
}

/**
* Parse the string and return an array which contains all path segments
*/
char** parse(char* str) {
	char ** messages;
	messages = (char**)malloc(sizeof(char *));
	char *p;
	p = strtok(str, " ");
	unsigned int i = 0;
	while (p != NULL) {
		p = strtok(NULL, "/");
		char *sp;
		boolean last = false;
		sp = strchr(p, ' ');
		if (sp != NULL) {
			*sp++ = '\0';
			last = true;
		}
		messages[i] = p;
		i++;
		if (last) {
			break;
		}
		messages = (char**)realloc(messages, sizeof(char *) * i + 1);
	}
	messages[i] = '\0';
	return messages;
}