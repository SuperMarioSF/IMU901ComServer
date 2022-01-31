# IMU901ComServer

A simple server to convert ATK-IMU901 serial data into WebSocket stream.

## Status

WIP, proof of concept at very early stage.

### Todo list
- [ ] Serial communication logic 
  - [X] Decode for reported measurement data
  - [ ] Data encoding for configuring IMU device
  - [ ] Decode for configuration confirmation data
- [ ] WebSocket Server  
  - [ ] WebSocket data payload format design
  - [ ] Combine WebSocket streaming with event data from IMU  
- [ ] Generic application logic
  - [ ] Able to select serial port by commandline argument
  - [ ] Able to select serial port by environment variable
  - [ ] Able to select serial port by config file

## Hardware

Purchase link (Taobao): https://detail.tmall.com/item.htm?id=623564385801&skuId=4581762262708

## Software

- `Golang 1.17`
- `go.bug.st/serial`

## License

Currently, WIP, no License were specified. Treat as "Propriety". Will change to other open source license later.

## Misc.

Datasheet file is in directory `datasheet`. This file is **NOT** covered in this repository's license. This file is only provided as a reference.

## By the way, what is this for?

This server is used to build one of my another project, which is building a personal AR HUD system.
The sensor data provided by my AR glasses internal IMU were not usable at all, so I have to use a external IMU to track the motion of glasses.
