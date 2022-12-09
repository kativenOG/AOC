from tqdm import tqdm 
# Lenght of the list after 80 days
data = list(map(int,open("./input.txt","r").read().split(",")))

def exp_fishes(Data,days):
    for day in tqdm(range(days)): #Cycle through days
        n_newFish = 0 
        for i,fish in enumerate(Data):
            if fish == 0:
                Data[i] = 6
                n_newFish+=1
            else:
                Data[i] -= 1 

        # Generate New fishes
        if n_newFish>0:
            for _ in range(n_newFish):
                Data.append(8)

    return len(Data)

# print(f"First Star: {exp_fishes(data,80)}")
print(f"Second Star: {exp_fishes(data,256)}")

