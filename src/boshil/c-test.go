package main

/*
#include <stdio.h>
#include <stdlib.h>

void say_hello() {
    printf("Hello world\n");
}
*/

import "C"
import "fmt"

func main() {
	C.say_hello()
}
