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
    crt = []
    register = 1 
    clock = 0 
    for command in Commands:
        if clock==40: clock=0
        if command=="noop":
            sprite = [int(register-1),register,int(register+1)]
            append_value =  "\u2588" if (clock in sprite) else " " 
            # print(f"Sprite: {[*sprite]}, Clock: {clock}, CRT: {append_value}")
            crt.append(append_value)
            clock+=1
        else:
            command = command.split()
            sprite = [int(register-1),register,int(register+1)]
            append_value = "\u2588" if (clock in sprite) else " " 
            # print(f"Sprite: {[*sprite]}, Clock: {clock}, CRT: {append_value}")
            crt.append(append_value)
            clock+=1
            if clock==40: clock=0
            append_value =  "\u2588" if (clock in sprite) else " " 
            # print(f"Sprite: {[*sprite]}, Clock: {clock}, CRT: {append_value}")
            crt.append(append_value)
            register += int(command[1]) 
            clock+=1
    return crt

print(f"First Star: {cpu(commands)}")
crt_output = crt(crt_commands)
print("Second Star: ")
previous_line = 0 
for line in range(0,280,40):
    for i in range(previous_line,line):
        print(crt_output[i],end='')
        sleep(0.01)
    print()
    previous_line = line
print()

