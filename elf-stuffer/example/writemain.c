#include <stdio.h>
#include <stdlib.h>

int main() {
  int *main_addr = (int *)&main;
  printf("main = %x\n", *main_addr);
  printf("writing to main...\n");
  *main_addr = 0xABCD;
  printf("cons = %x\n", *main_addr);
}
