#define wait(s) atomic { s > 0 -> s-- }  //exclusion mutua con atomic
#define signal(s) s++

byte room = 4
byte fork[5] = {1,1,1,1,1}

active proctype P() {
	byte i = 0
	do
	:: printf("P%d THINK 1\n", _pid)
	   wait(room)
	   wait(fork[i])
	   wait(fork[i+1])
	   printf("P%d EAT 1\n", _pid)
	   i++
	   
	   signal(fork[i])
	   signal(fork[i+1])
	   signal(room)
	   i >= 4 -> i = 0 
	od
}