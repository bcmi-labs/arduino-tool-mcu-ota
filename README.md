# WiFi Link MCU OTA
Microcontroller OTA upload tool for ESP8266 Arduino based boards running WiFi Link


## Usage
### Required Arguments:
* *-i*&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;IP address of the destination board.
* *-f*&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Firmware file to upload onto the mcu

### Optional Arguments:
* *-p*&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Connection port [default 80]
* *-l*&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Number of lines of the specified files to upload for each request
* *-h*&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Help


#### examples:

```shell
$> ./arduino_mcuota -f /tmp/Blink.ino.hex -i 192.168.0.110 -p 80

Sending /tmp/Blink.ino.hex to host 192.168.0.110
[1 / 15] Done...
[2 / 15] Done...
[3 / 15] Done...
[4 / 15] Done...
[5 / 15] Done...
[6 / 15] Done...
[7 / 15] Done...
[8 / 15] Done...
[9 / 15] Done...
[10 / 15] Done...
[11 / 15] Done...
[12 / 15] Done...
[13 / 15] Done...
[14 / 15] Done...
[15 / 15] Done...
Upload Completed
```

## Build
You can build this tool the same way in Mac, Windows and Linux platforms

### Prerequisites
* golang [installation](https://golang.org/doc/install)
* [resty libraries](https://github.com/go-resty/resty)


### Build
To build run this command:
```shell
$> go build arduino_mcuota.go
```
It creates a `build` and `dist` folder. Destination binary is inside `dist` folder

That's all.
