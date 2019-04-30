#define wait(s) atomic { s > 0 -> -- }
#define signal(s) s++

byte s = 1
byte array[4]

active proctype P() {
	byte index = 0;
	byte random[4]
	random[0] = 5
	random[1] = 3
	random[2] = 7
	random[3] = 1
	
	  do
	  :: index < 4 -> array[index] = random[index]
		 printf("unordered array: %d\n", array[index])
		 index++	
		 
	  :: else -> break
	  od
}
