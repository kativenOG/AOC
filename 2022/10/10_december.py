commands = open("./input.txt","r").read().splitlines()

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

print(f"First Star: {cpu(commands)}")
