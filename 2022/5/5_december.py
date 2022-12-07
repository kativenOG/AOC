file = open("./input.txt").read().splitlines()

# Getting the initial configuration
n  = max(list(map(int,file[8].split()))) # Number of Colums 
cargo_ship = [[] for _ in range(n)]
for line in file[:10]: # Carico i Carichi 
    for i,char in enumerate(line):
        if char.isalpha(): cargo_ship[round(i/4)].append(char)

# Inverto le liste per avere come ultimo elemento quello che è da poppare  (ultimo == top of the stack)
for i in range(n): 
    cargo_ship[i] = cargo_ship[i][::-1]

# Lista dei Comandi! 
commands = file[10:] 

def first_star(commandss,cargo):
    # Re-arranging Crates: CrateMover 9000
    for command in commandss:
        command = command.split()
        for _ in range(int(command[1])):
            gru = cargo[int(command[3])-1].pop()
            cargo[int(command[5])-1].append(gru)

    # Calculating the Result 
    result = []
    for i in range(9):
        result.append(cargo[i][len(cargo[i])-1])
    return result

def second_star(commandss,cargo):
    # Re-arranging Crates: CrateMover 9001
    for command in commandss:
        command = command.split()
        gru = []
        for _ in range(int(command[1])):
            appo = cargo[int(command[3])-1].pop()
            gru.append(appo) 
        for appo in gru[::-1]:
            cargo[int(command[5])-1].append(appo)
    print(cargo)
    # Calculating the Result 
    result = []
    for i in range(9):
        result.append(cargo[i][len(cargo[i])-1])
    return result

# print(f"First Star: {first_star(commands,cargo_ship)}\n")
print(f"Second Star: {second_star(commands,cargo_ship)}")
