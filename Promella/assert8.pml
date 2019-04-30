int np = 0
int nq = 0
byte critical = 0

active proctype P() {
	do
	:: true ->
		np = nq + 1
		if
		:: (nq == 0 || np <= nq) ->
			critical++
			assert(critical <= 1)
			critical--
			np = 0
		fi
	od
}

active proctype Q() {
	do
	:: true ->
		nq = np + 1
		if
		:: (np == 0 || nq < np) ->
			critical++
			assert(critical <= 1)
			critical--
			nq = 0
		fi
	od
}

init {
	atomic {
		run P()
		run Q()
	}
}