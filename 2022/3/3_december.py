import math 
data = open("./input-txt", "r").read().splitlines()

def first_star(Data):
    result = 0 
    for line in Data:
        halfLine = math.floor((len(line))/2)
        first,second = line[:halfLine],line[halfLine:]
        for val in first: # Cycle in first Half
            if second.find(val) != -1:
                    result += (ord(val) - 96) if(val.islower()) else (ord(val) - 38)
                    break
    return result

def second_star(Data): # Priorities badges 
    result = 0 
    for n_group in range(math.floor(len(Data)/3)):
        for val in Data[n_group*3]:
            if Data[(n_group*3 + 1)].find(val) != -1 and Data[(n_group*3 + 2)].find(val) != -1:
                    result += (ord(val) - 96) if(val.islower()) else (ord(val) - 38)
                    break

    return result

print(f"First Star: {first_star(data)}\nSecond Star: {second_star(data)}")
