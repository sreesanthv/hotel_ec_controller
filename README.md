# Hotel Management Service

## About the application
* A base Skelton for start implementing service which watches for Motions in a hotel/infrastructure, and manages devices to save power consumption.
* Consider a hotel with n number of Floors, m number of Main Corridors per floor, and s number of Sub corridors per floor.
* Each Main & Sub Corridor contains one independently controllable AC and Bulb.
* AC consumes 10 units of power and Bulb consumes 5 units when they are turned on.
* By default, all ACs should be ON. And same goes for the Main corridor Bulb. But Sub corridor Bulb should be OFF.
* When a motion is detected in a Sub Corridor, turn on it's Bulb.
* The total power consumption of all the ACs and lights combined should not exceed
```(m * 15) + (s * 10) units of per floor```.
Sub corridor AC could be switched OFF to ensure that the power consumption is not more than the specified maximum value
* When there is no motion for more than a minute the sub corridor lights should be switched OFF and AC needs to be switched ON.



## Run the application

### Config file
* Open ```config.json```
* Set No of floors, no of main corridors per floor, and no of sub corridors.
* Config sample: ```{"no_of_floors":2,"main_corridors_per_floor":1,"sub_corridors_per_floor":2}```

### Start service
Run ```go run main.go```

### Input motion to service
* Open ```motion_ip.json```
* Input motion details in this file.
* Input Sample: ```{"floor":1,"sub":1}```
* Save the file. The service will automatically read and process input.

## Unit tests
Run ```go test ./...```