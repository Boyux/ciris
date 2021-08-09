//
// Created by Hei on 2021-07-13.
//

#ifndef RUST_ZEBRA_H
#define RUST_ZEBRA_H

#include <stdlib.h>

typedef struct Token {
    uint8_t buf[40];
    size_t len;
} Token;

Token *zebra_gen_token(char *frame);

int zebra_verify_token(char *frame, char *input);

void zebra_free_token(Token *token);

#endif //RUST_ZEBRA_H
