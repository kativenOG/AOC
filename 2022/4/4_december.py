data = open("input.txt","r").read().splitlines()

def first_star(Data):
    result = 0
    for line in Data:
        ranges = line.split(",")
        first,second= list(map(int,ranges[0].split("-"))) , list(map(int,ranges[1].split("-")))
        if (second[0] >= first[0] and first[1] <= second[1]) or (first[0] >= second[0] and second[1] <= first[1]):
            result+=1
    return result 





print(f"First Star: {first_star(data)}")
