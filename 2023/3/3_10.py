# special_characters = "!@#$%^&*()-+?_=,<>/"
from icecream import ic

class posnum():
    def __init__(self,num: str ,pos: list, directions: list=[]) -> None:
        self.num = int(num)
        self.pos = pos 
        self.directions = directions
    
    def __str__(self) -> str:
        return f'PosNum {self.num} with pos: {self.pos}' 

    def check_sc(self,sc_list)->int:
        found = False
        for pos in self.pos:
            if found: break 
            for dir in self.directions:
                new_pos = (pos[0]+dir[0],pos[1]+dir[1])
                if new_pos in sc_list: 
                    found = True
                    break  
            
        return self.num if found else 0

    def check_gear(self,gear_list):
        # ic(gear_list)
        gears = set()
        for pos in self.pos:
            for dir in self.directions:
                new_pos = (pos[0]+dir[0],pos[1]+dir[1])
                # ic(new_pos)
                if new_pos in gear_list: 
                    gears.add(new_pos)
        return tuple((self.num, gears))

def first_star(input:list)->int:
    # Direction list 
    directions = [[1,0],[0,1],[1,1],[-1,0],[0,-1],[-1,-1],[-1,1],[1,-1]]
    nums_list, sc_list= [],[]
    for i,line in enumerate(input):
        num = ''
        for j,value in enumerate(line):
            if value.isnumeric():
                num+= value
            else:
                if num!='':
                    n_len = len(num)  
                    nums_list.append(posnum(num=num, pos=[(i,j-val-1) for val in range(n_len)],directions=directions)) 
                    num=''
                # if value in special_characters:
                if not value.isalnum() and value!='.': 
                    sc_list.append((i,j))
        # Add the last Number if there  
        if num != '':
            n_len = len(num)  
            line_len = len(line)
            nums_list.append(posnum(num=num, pos=[(i,line_len-val-1) for val in range(n_len)],directions=directions)) 
            num=''

    sum = 0 
    for num in nums_list:
        sum+=num.check_sc(sc_list)
    return sum 
    

def second_star(input:list)->int:
    directions = [[1,0],[0,1],[1,1],[-1,0],[0,-1],[-1,-1],[-1,1],[1,-1]]
    nums_list, gear_list= [],[]
    for i,line in enumerate(input):
        num = ''
        for j,value in enumerate(line):
            if value.isnumeric():
                num+= value
            else:
                if num!='':
                    n_len = len(num)  
                    nums_list.append(posnum(num=num, pos=[(i,j-val-1) for val in range(n_len)],directions=directions)) 
                    num=''
                # if value in special_characters:
                if value=='*': 
                    gear_list.append((i,j))
        # Add the last Number if there  
        if num != '':
            n_len = len(num)  
            line_len = len(line)
            nums_list.append(posnum(num=num, pos=[(i,line_len-val-1) for val in range(n_len)],directions=directions)) 
            num=''

    num_gears_pairs = []
    for num in nums_list:
        num_gears_pairs.append(num.check_gear(gear_list))
    
    sum = 0 
    for i, pair1 in enumerate(num_gears_pairs):
            num1, gears1 = pair1 
            # print(f'ORIGINAL\nNum: {num1} -> Gears: {gears1}')
            for j,pair2 in enumerate(num_gears_pairs): 
                if i==j: continue 
                num2, gears2 = pair2
                # print(f'COMPARISON Num: {num2} -> Gears: {gears2}')
                inter_len= len(gears1.intersection(gears2))
                if inter_len!=0:
                    sum+= num1*num2
         

    return sum//2

# Transform the input into a matrix  
# input = open('test.txt','r').read().splitlines()
input = open('input.txt','r').read().splitlines()
input = [list(line) for line in input]
input_ = input.copy()
# Results:
res = first_star(input)
print(f'First Star: {res}')
res = second_star(input)
print(f'Second Star: {res}')
