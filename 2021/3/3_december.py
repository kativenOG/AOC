import math ,copy 
file = open("./input.txt","r")
data = file.read().split("\n")
data.pop(len(data)-1)

def first_star():
    gamma_rate = "" 
    epsilon_rate =  ""
    
    trashold= math.floor(len(data)/2) # to check if it's a 1 or a 0 
    l = len(data[0]) 
    counters = [0 for _ in range(l)] # counts the number fo values 
    
    # counting the total number of 1s in the text 
    for digits in data:
        for i,digit in enumerate(digits):
            if (digit == "1"): counters[i]+=1 

    # Calculating gamma and epsilon
    for digit in counters:
        gamma_rate += "1" if (digit > trashold)  else "0"
        epsilon_rate+= "0" if (digit > trashold)  else "1"

    # Casting the string back to base ten from base 2 and then returning their multiplication 
    gamma_rate = int(gamma_rate,2)
    epsilon_rate = int(epsilon_rate,2)
    return (gamma_rate*epsilon_rate) 


def evaluate(mask,data):
    result = 0 
    for digit,target in enumerate(mask):
        for i,value in enumerate(data):
            if(len(data)==1): 
                print(value,len(data))
                result = int(value,2)
                break
            elif int(value[digit]) != target: 
                data.pop(i) 
                i+=1
        if(len(data) == 1): break 
    return result 


def second_star():
    oxigen= 0 
    co2= 0  

    trashold= math.floor(len(data)/2) # to check if it's a 1 or a 0 
    l = len(data[0]) 
    counters = [0 for _ in range(l)] # counts the number fo values 

    # counting the total number of 1s in the text 
    for digits in data:
        for i,digit in enumerate(digits):
            if (digit == "1"): counters[i]+=1 

    # Generating masks 
    oxigen_mask = [0 for _ in range(len(counters))] 
    co2_mask = [0 for _ in range(len(counters))] 
    for i,value in enumerate(counters):
       oxigen_mask[i] = 1 if ((value)>(trashold - 1)) else 0 
       co2_mask[i] = 1 if (value<(trashold + 1)) else 0 
        
    # Finding values by removing value not equal to the mask 
    # Probabilmente non servono deep copy ma sono autistico e voglio essere sicuro 
    o_data = copy.deepcopy(data) 
    c_data = copy.deepcopy(data)
    oxigen = evaluate(oxigen_mask,o_data)
    print("Oxigen ",oxigen) 
    co2 = evaluate(co2_mask,c_data)
    print("CO2",co2) 
    return (oxigen*co2)
    
print("First Star result: ", first_star())
print("Second Star result: ", second_star())
