from copy import deepcopy
from tqdm import tqdm 

class category():
    def __init__(self)->None:
        self.range_list = []
        
    def add_range(self,destination,src,rangee):
        self.range_list.append(tuple((src,src+rangee,destination))) 
            
    def sort_ranges(self)->None:
        self.range_list.sort()

    def check_value(self,seed_value)-> int:
        for rangee in self.range_list:
            src_start, src_end, destination = rangee
            if seed_value < src_start: return seed_value 
            if src_start <= seed_value <= src_end:
                return seed_value - src_start + destination
            
        return seed_value 
        # raise ValueError('A very specific bad thing happened.')

    def free_memory(self)->None:
        self.range_list = []

    def __str__(self)->str:
        return f'Category with ranges:\n {self.range_list}\n'
        


def first_star(input: list)->int:
    seeds = input[0].split(':')[1]
    seeds = list(map(int,seeds.split()))
    del input[0]
    
    categories = []

    for line in input:
        if len(line)<2: continue 
        elif '-' in line: 
            categories.append(category())
        else:
            vals= list(map(int,line.split()))
            categories[-1].add_range(vals[0],vals[1],vals[2])
    
    seeds_locs = seeds
    for cat in categories:
        cat.sort_ranges()
        new_vals = []
        for seed in seeds_locs:
            new_vals.append(cat.check_value(seed))
        seeds_locs = new_vals


    return min(seeds_locs)

def second_star(input: list)->int:
    seeds_line= input[0].split(':')[1]
    categories = []

    for line in input[1:]:
        if len(line)<2: continue 
        elif '-' in line: 
            categories.append(category())
        else:
            vals= list(map(int,line.split()))
            categories[-1].add_range(vals[0],vals[1],vals[2])
    for cat in categories: cat.sort_ranges()
    del input 

    # Divide and Conquer 
    min_seeds = 100000000000000000000000000000
    seeds_line = list(map(int,seeds_line.split()))
    previous = -1 
    for seed in seeds_line: 
        if previous == -1: previous = seed 
        else: 
            for seed in range(previous,previous+seed):
                new_val = 0 
                for cat in categories:
                    new_val = cat.check_value(seed)
                min_seeds = min(min_seeds,new_val)
            previous = -1

    return min_seeds

# input = open('test.txt','r').read().splitlines()
input = open('input.txt','r').read().splitlines()
input_= deepcopy(input)
minn = first_star(input)
print(f'First Star result: {minn}')
minn = second_star(input_)
print(f'Second Star result: {minn}')


