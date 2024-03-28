#!/bin/env python3
# -*-coding: utf-8 -*-


from construct import (
    ConstructError,
)
import sys
import elfformat
from capstone import Cs, CS_ARCH_ARM64, CS_MODE_ARM


def main() -> None:
    if len(sys.argv) < 2:
        print(f"Usage: {sys.argv[0]} [a.out]")
        exit(1)

    with open(sys.argv[1], "rb") as f:
        bytes = f.read()

        print(f"Parsing {sys.argv[1]}...")
        data = try_parse(elfformat.elf, bytes)

        segments = []

        for segment in data.body.program_table:
            rg = range(segment.virtual_address, segment.virtual_address + segment.size_mem)
            print(f"Memory range is {rg}")
            if (data.body.entry in rg):
                segments += (rg, segment)
            
            if (segment.p_type=="PT_DYNAMIC"):
                print(segment)

        print(f"Disassembling {sys.argv[1]}...")
        CODE = bytes[data.body.entry : segments[1].virtual_address + segments[1].size_mem]

        md = Cs(CS_ARCH_ARM64, CS_MODE_ARM)

        for address, size, mnemonic, op_str in md.disasm_lite(CODE, data.body.entry):
            print("0x%x:\t%s\t%s" % (address, mnemonic, op_str))


def hexdump(b: bytes) -> None:
    import binascii

    i = 0
    while i < len(b):
        for _ in range(4):
            block = b[i : i + 8]
            if block:
                print(binascii.hexlify(block), end=" ")
            i += 8
            print()


def try_parse(format, bytes):
    try:
        data = format.parse(bytes)
    except ConstructError as e:
        print(e)
        exit(1)
    return data


if __name__ == "__main__":
    main()
