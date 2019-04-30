int turn = 1;
int critical = 0;

proctype P() {
	byte i = 0
	do
	:: i < 100 ->
		printf("P%d: Non critical section\n",_pid);
		printf("P%d: Non critical section\n",_pid);
		if
		:: turn == 1 ->
				critical++
				assert(critical <= 1)
				printf("P%d: Critical section\n",_pid)
				printf("P%d: Critical section\n",_pid)
				critical++
				turn = 2
		fi
		i++
	od
}

proctype Q() {
	byte i = 0
	do
	:: i < 100 ->
		printf("P%d: Non critical section\n",_pid);
		printf("P%d: Non critical section\n",_pid);
		if
		:: turn == 2 ->
				critical++
				assert(critical <= 1)
				printf("P%d: Critical section\n",_pid)
				printf("P%d: Critical section\n",_pid)
				critical++
				turn = 1
		fi
		i++
	od
}

init {
	atomic {
		run P()
		run Q()
	}
}