import copy 
from collections import defaultdict


def first_star(input:list,red: int=12, green: int=13,blue: int=14)->int:
    odds = {'red':red, 'blue':blue, 'green':green}
    sum = 0
    for i,game in enumerate(input): 
        valid = True  
        for key, item_list in game.items():
            if max(item_list) > odds[key]: valid = False
        if valid: sum+=i+1

    return sum 

def second_star(input:list)->int:
    sum = 0 
    for game in input:
        mul = 1
        for value in game.values():
            mul*=max(value) 
        sum+=mul
            
    return sum 

input = open('input.txt','r').read().splitlines()
input_ = [] 
for line in input:
    game_dict = defaultdict(list)

    split1 = line.split(':')

    balls_list = split1[1].replace(';',',')
    balls = balls_list.split(',')
    for ball in balls:
        split2 = ball.strip().split()
        game_dict[split2[1]].append(int(split2[0]))

    input_.append(game_dict)        
    

input__ = copy.deepcopy(input_)
sum = first_star(input_)
print(f'First Star: {sum}')
sum = second_star(input__)
print(f'Second Star: {sum}')

