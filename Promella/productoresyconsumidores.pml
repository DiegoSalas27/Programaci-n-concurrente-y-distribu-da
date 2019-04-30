//definimos semaforos // little book of semaphores
#define wait(s) atomic { s > 0 -> s-- }
#define signal(s) s++

#define NUM_PROD 4 //productor
#define NUM_CONS 3 // consumidor
#define MAX 10

int buffer[MAX]
int n = 0 //contador de buffer

byte sbuff = 1
byte sprod = MAX
byte scons = 0

byte critical = 0

active[NUM_PROD] proctype producer() {
	int event
	int cont = 0
	do
	:: event = _pid * 1000 + cont;
	   wait(sprod)
	   wait(sbuff)
	   critical++
	   buffer[n] = event
	   n++
	   assert critical <= 1
	   critical--
	   signal(sbuff)
	   signal(scons)
	   cont++
	od
}

active[NUM_CONS] proctype consumer() {
	int event
	do
	:: 
	   wait(scons)
	   wait(sbuff)
	   critical++
	   n--
	   event = buffer[n]
	   assert critical <= 1
	   critical--
	   signal(sbuff)
	   signal(sprod) 
	   printf("Consumidor %d  usa %d\n", _pid, event)
	od
}