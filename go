
avr-gcc -g -Os -mmcu=atmega328p -c  -D F_CPU=8000000 XXcontrol_KR.c
avr-gcc -g -mmcu=atmega328p -o  XXcontrol_KR.elf XXcontrol_KR.o

avr-objdump -h -S XXcontrol_KR.elf > XXcontrol_KR.lst

avr-objcopy -j .text -j .data -O ihex  XXcontrol_KR.elf XXcontrol_KR.hex

echo to program:
echo avrdude -c eet  -p m328p -U flash:w:XXcontrol_KR.hex
