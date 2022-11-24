file = open("./input.txt","r")
data = file.read().split("\n")
data.pop(len(data)-1)
# print(data)

def first_star():
    horizzontal = 0 
    depth = 0 
    for line in data:
        command = line.split()
        horizzontal += int(command[1]) if (command[0]=="forward") else 0  
        depth += int(command[1]) if (command[0]=="down") else -int(command[1]) if (command[0] == "up") else 0 
        
    return(horizzontal*depth) 


def second_star():
    horizzontal = 0 
    depth = 0 
    aim = 0 
    for line in data:
        command = line.split()
        aim += int(command[1]) if (command[0] == "down") else -int(command[1]) if(command[0] == "up") else 0 
        if command[0] == "forward":
            horizzontal+= int(command[1])
            depth += int(command[1])*aim
    return (horizzontal*depth)


print("First Star result: ", first_star())
print("Second Star result: ", second_star())

