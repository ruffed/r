# ELF Packer

Copyright 2023 (c) Aiden Fox Ivey <aiden@fox-ivey.com>, Alisya K.

## Methodology

Under the hood we use [Construct](https://construct.readthedocs.io/en/latest/) to parse the ELF files. We've written our code so that
there is support for 32 and 64 bit ELF files. As is expected, the parser respects the `e_ident[EI_DATA]` rather than assuming the provided
binary is the host encoding or an arbitrary one.

Thanks must be given to the [elf32.py](https://github.com/construct/construct/blob/master/deprecated_gallery/elf32.py) example provided
in Construct's deprecated gallery, as it demonstrated a very clean way to construct the code.

A version of our current example has been [added to Construct](https://github.com/construct/construct/blob/master/gallery/elf.py).

Also, we use [Capstone Dissassembler](http://www.capstone-engine.org/) for optional dissassembly of code within the files.

## Resources

AARCH64 Syscalls available [here](https://chromium.googlesource.com/chromiumos/docs/+/HEAD/constants/syscalls.md#arm64-64_bit).

ELF resources provided [by tmp.out](https://github.com/tmpout/awesome-elf).

Some inspiration provided [from Faster Than Lime](https://fasterthanli.me/series/making-our-own-executable-packer).
