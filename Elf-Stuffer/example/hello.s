    .arch       armv8-a
    .section    .rodata
    .text
.LC0:
    .string     "Hello World!\n"
    .align      2
    .global      _start
_start:
    mov     x0, #1
    ldr     x1, =.LC0
    mov     x2, #13
    mov     x8, #64
    svc     #0
    mov     x0, #0
    mov     x8, #93
    svc     #0
