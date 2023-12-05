from copy import deepcopy
import numpy as np 

def first_star(input: list)->int:
    prize = 0
    for line in input:
        single_prize = 0 
        values = line.split(':')[1]
        played,winning = values.split('|')
        played, winning= list(map(int,played.split())), list(map(int,winning.split())) 
        for num in played:
            if num in winning: 
                single_prize = 1 if single_prize==0 else single_prize*2
        prize+=single_prize

    return prize 

def second_star(input: list)->int:         
    sc_prize_list = []
    for line in input:
        single_prize = 0 
        values = line.split(':')[1]
        played,winning = values.split('|')
        played, winning= list(map(int,played.split())), list(map(int,winning.split())) 
        for num in played:
            if num in winning: 
                single_prize += 1 
        sc_prize_list.append(single_prize)
    
    len_prizes = len(sc_prize_list)
    cards = np.array([1 for _ in range(len_prizes)])
    for i,value in enumerate(sc_prize_list):
        if value==0: continue  
        i_pos = i+1
        cards[i_pos:i_pos+value] = cards[i_pos:i_pos+value] + 1*cards[i]

    return sum(cards)


# input = open('test.txt','r').read().splitlines()
input = open('input.txt','r').read().splitlines()
input_= deepcopy(input)
prize = first_star(input)
print(f'First Star Result: {prize}')
prize = second_star(input_)
print(f'Second Star Result: {prize}')


