from copy import deepcopy

def first_star(input:list)->int:
    mull = 1 
    times = list(map(int,input[0].split(':')[1].split()))
    distances = list(map(int,input[1].split(':')[1].split()))
    
    for time,distance in zip(times,distances):
        sum = 0 
        for current_time in range(time):
            if current_time*(time-current_time)>=distance: sum+=1
        mull*=sum

    return mull

def second_star(input:list)->int:
    sum = 0    
    time = int(''.join(input[0].split(':')[1].split()))
    distance = int(''.join(input[1].split(':')[1].split()))

    for current_time in range(time):
        if current_time*(time-current_time)>=distance: sum+=1

    return sum


input = open('input.txt','r').read().splitlines()
input_ = deepcopy(input)
res = first_star(input)
print(f'First Star result: {res}')
res = second_star(input_)
print(f'Second Star result: {res}')
