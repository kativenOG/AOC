file = open("./input.txt","r")
data = file.read().split("\n")
print(data.pop(len(data)-1))

def second_star():
    sum1,sum2,sum3 = [0,0,0]
    solution = [] 
    sum_counter = 0 
    for i,_ in enumerate(data):
        if i == 2:
            sum1 = float(data[i]) + float(data[i-1]) + float(data[i-2])
            solution.append(f"{sum1} (N/A - no previous sum)")
        if i>2:
            sum1 = float(data[i-1]) + float(data[i-2]) + float(data[i-3])
            sum2 = float(data[i]) + float(data[i-1]) + float(data[i-2])
            result = "(no change)" if (sum1==sum2) else "(increased)" if (sum2>sum1) else "(decreased)" 
            solution.append(f"{sum1} {result}")
            if sum2>sum1: sum_counter+=1
    return (sum_counter)

print(second_star())
