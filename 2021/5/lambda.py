"""
PYTHON LAMBDAS:
    
Lambdas in Python are based on lambda calcolus from Alonzo Church, church-turing hypotesis
to convert one of the two computation model into the other
Immagina le lambda come le callback di python
"""
lamb = lambda x: x+1 # Function called lamb
print(lamb(1))

print((lambda x,y: x+y)(1,2)) # Calling the function immediatly 

# Lambdas can take in input other functions, so:
first_lambda = lambda x,fun: x + fun(x)
print(first_lambda(5,(lambda x: x*x)))



