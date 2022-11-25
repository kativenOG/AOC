# La maschera deve essere calcolata ogni volta che cambio digit, e non solo all'inizio
import math ,copy 
# import pdb 

file = open("./input.txt","r")
data = file.read().split("\n")
data.pop(len(data)-1)

def first_star():
    gamma_rate = "" 
    epsilon_rate =  ""
    
    treshold= math.floor(len(data)/2) # to check if it's a 1 or a 0 
    l = len(data[0]) 
    counters = [0 for _ in range(l)] # counts the number fo values 
    
    # counting the total number of 1s in the text 
    for digits in data:
        for i,digit in enumerate(digits):
            if (digit == "1"): counters[i]+=1 

    # Calculating gamma and epsilon
    for digit in counters:
        gamma_rate += "1" if (digit > treshold)  else "0"
        epsilon_rate+= "0" if (digit > treshold)  else "1"

    # Casting the string back to base ten from base 2 and then returning their multiplication 
    gamma_rate = int(gamma_rate,2)
    epsilon_rate = int(epsilon_rate,2)
    return (gamma_rate*epsilon_rate) 


def generate_mask(current_data,digit,mode):

    return_mask = 0
    treshold= math.floor(len(current_data)/2)  

    for value in current_data:
        if(int(value[digit])==1): return_mask+=1

    if(mode==0): return_mask= 1 if ((return_mask)>(treshold - 1)) else 0 
    else: return_mask= 1 if (return_mask<(treshold + 1)) else 0 
    print("Return mask:",return_mask)
    # pdb.set_trace()
    return return_mask 

def evaluate(edata,mode):
    result = 0 
    for digit in range(12):

        target = generate_mask(edata,digit,mode)

        for i,value in enumerate(edata):
            if(len(edata)==1): 
                print(value,len(edata))
                break
            elif int(value[digit]) != target: 
                edata.pop(i) 
                i-=1

        if(len(data) == 1): 
            result = int(edata[0],2)
            break 

    return result 

def second_star():
    oxigen= 0 
    co2= 0  
       
    # Finding values by removing value not equal to the mask 
    # Probabilmente non servono deep copy ma sono autistico e voglio essere sicuro 
    o_data = copy.deepcopy(data) 
    c_data = copy.deepcopy(data)

    oxigen = evaluate(o_data,mode=0)
    print("Oxigen ",oxigen) 
    co2 = evaluate(c_data,mode=1)
    print("CO2",co2) 

    return (oxigen*co2)
    
# print("First Star result: ", first_star())
print("Second Star result: ", second_star())
