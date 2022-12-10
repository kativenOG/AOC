from time import sleep
from copy import deepcopy 
commands = open("./input.txt","r").read().splitlines()
crt_commands = deepcopy(commands)

def cpu(Commands):
    signal= 0 
    register = 1 
    clock = 0 
    diff_action = False
    for command in Commands:
        command = command.split()
        clock += 1 # FIRST CYCLE:
        if not diff_action:
            if ((clock-20)%40) == 0:
                signal += clock*register
                diff_action = True 
        else: diff_action = False
        clock += 1 if command[0] == "addx" else 0 # SECOND OPTIONAL CYCLE:
        if not diff_action:
            if ((clock-20)%40) == 0:
                signal += clock*register
                diff_action = True 
        else: diff_action = False
        register += int(command[1]) if command[0] == "addx" else 0 # AT THE END OF THE SECOND CYCLE CHANGE THE REGISTER:

    return signal 

def crt(Commands):
    crt = [[] for _ in range(6)]
    crt_pos = 0 
    crt_line= 0 

    register = 1 
    clock = 0 
    for command in Commands:
        clock+=1
        crt_pos+=1
        if command=="noop":
            append_value = 1 if (register==crt_pos or int(register+1)==crt_pos or int(register-1)==crt_pos) else 0
            crt[crt_line].append(append_value)
        else:
            sprite = [int(register-1),register,int(register+1)]
            command = command.split()
            append_value = 1 if (clock in sprite) else 0
            crt[crt_line].append(append_value)

            crt_pos+=1
            clock+=1
            append_value = 1 if (clock in sprite) else 0
            crt[crt_line].append(append_value)
            register += int(command[1]) 
        if(crt_pos == 39):
            crt_line +=1  
            crt_pos = 0 
    return crt


print(f"First Star: {cpu(commands)}")
def crt_drawing(crt):
    for line in crt:
        print("\n")
        for value in line:
            if value ==1: print("\u2588",end=" ")
            else: print(" ",end=" ")
            # sleep(0.2)
    return 
crt_output = crt(crt_commands)
print("Second Star: ")
crt_drawing(crt_output)

