active proctype P() {
	byte dividendo = 100
	byte divisor = 13
	byte cociente;
	byte residuo = dividendo
	
	do
	:: residuo > divisor ->
		cociente++
		residuo = residuo - divisor
	:: else -> break
	od
	
	assert(cociente * divisor + residuo == dividendo)
}