file = open("input.txt","r")
data = file.read().split("\n\n")

def snack_provider(Data):
    elfs = [] 
    for i,elf in enumerate(Data):
        if i == (len(Data)-1): elf = elf[:len(elf)-1] 
        cal= list(map(int,elf.split("\n")))
        elfs.append(sum(cal)) 
    elfs.sort()
    return elfs[-1:],sum(elfs[-3:]) 

firstStar,secondStar = snack_provider(data)
print(f"First star result:{firstStar[0]}\nSecond star result: {secondStar}")
