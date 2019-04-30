#define wait(s) atomic { s > 0 -> s-- }  //exclusion mutua con atomic
#define signal(s) s++

#define NUM_PROCS 2

byte s = 1

//verificar si hay exclusion mutua
byte critical = 0

active[NUM_PROCS] proctype P() {
	do
	:: 
	   wait(s)
	   critical++
	   assert critical <= 1 //verificar si hay exclusion mutua
	   critical--
	   signal(s)	
	od
}