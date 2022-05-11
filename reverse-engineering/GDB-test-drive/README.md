chmod +x gdbme
gdb gdbme
layout asm
break *(main+99)
run
jump *(main+104)
