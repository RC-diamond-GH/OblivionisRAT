#ifndef Base64
#define Base64
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
typedef unsigned int uint32_t;

const char base64_chars[] = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";

char *base64_encode(const unsigned char *input, size_t length);

unsigned char *base64_decode(const char *input, size_t *output_length);

#endif