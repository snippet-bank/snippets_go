Computers operate fundamentally on fixed-size numbers called "words".

Modern personal computer CPUs use a word size of 64 bits.

CPU architecture is related to the size of the word.

For example, 1 word is the number of bytes the address bus can handle. 

A 64-bit bus is literally 64 wires that, in a given instant, transmit in parallel a high or low voltage -- 64 ones or zeroes.

I.e., 8 bytes of data can flow thru the bus each time the CPU clock ticks.

Since the entire computer architecture is organized around data sized at 1 "word", your data needs to be in chunks that line up in units of words.

Therefore, as it compiles your code, Go adds padding to align your data along word boundaries for optimal fetching.

This snippet demonstrates that optimization.




