gcc -static main.cpp Network.hpp Network.cpp OblivionisAES.hpp OblivionisAES.cpp base64.hpp base64.cpp debug.hpp debug.cpp KeyExchange.hpp KeyExchange.cpp command.hpp command.cpp -o blank.exe -lws2_32 -lstdc++
go run GenerateBeacon.go
.\beacon.exe