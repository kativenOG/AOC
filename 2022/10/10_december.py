# from time import sleep
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
    crt = []
    register = 1 
    clock = 1 
    for command in Commands:
        clock+=1
        if clock==41: clock=0
        if command=="noop":
            sprite = [int(register-1),register,int(register+1)]
            append_value =  "\u2588" if (clock in sprite) else " " 
            crt.append(append_value)
        else:
            command = command.split()
            sprite = [int(register-1),register,int(register+1)]
            append_value = "\u2588" if (clock in sprite) else " " 
            crt.append(append_value)

            clock+=1
            if clock==41: clock=0
            append_value =  "\u2588" if (clock in sprite) else " " 
            crt.append(append_value)
            register += int(command[1]) 
    return crt

# print(f"First Star: {cpu(commands)}")
crt_output = crt(crt_commands)
print("Second Star: ")
for line in range(6):
    for i in range(0,40):
        pos = i+ ( (40*(line)) -1) if line!=0 else i
        print(crt_output[pos],end='')
    print("\n")

