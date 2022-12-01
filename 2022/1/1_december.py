file = open("input.txt","r")
data = file.read().split("\n\n")

def first_star(Data):
    result = 0 
    for i,elf in enumerate(Data):
        if i == (len(Data)-1): elf = elf[:len(elf)-1] 
        cal= list(map(int,elf.split("\n")))
        result = max(sum(cal),result)
    return result 

def second_star(Data):
    elfs = [] 
    for i,elf in enumerate(Data):
        if i == (len(Data)-1): elf = elf[:len(elf)-1] 
        cal= list(map(int,elf.split("\n")))
        elfs.append(sum(cal)) 
    elfs.sort()
    return sum(elfs[-3:]) 


#print("First star: ",first_star(data))
print("Second star: ",second_star(data))
