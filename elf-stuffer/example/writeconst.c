#include <stdio.h>
#include <stdlib.h>

const int cons = 29;

int main() {
  int *cons_addr = (int *)&cons;
  printf("cons = %d\n", *cons_addr);
  printf("writing to cons...\n");
  // note: `cons = 31` doesn't compile
  *cons_addr = 31;
  printf("cons = %d\n", *cons_addr);
}
