import copy 
from collections import deque

# Standard Digit Encoding 
digits = {"one": '1', 'two':'2','three':'3','four':'4','five':'5','six':'6','seven':'7','eight':'8','nine':'9'}
# Add reverse Items 
iterable = list(digits.items())
for key,value in iterable:
    digits[key[::-1]] = value 


def first_star(input: list)->int:
    sum = 0
    for line in input:
        number = ''
        for char in line:
            if char.isdigit(): 
                number+=char
                break
        for char in line[::-1]:
            if char.isdigit():
                number+=char
                break
        sum+=int(number)
    return sum 



def CheckForDigit(line: str)-> str:
    number = ''
    buffer = deque(maxlen=5)
    for char in line:
        buffer.append(char)
        current_buffer = ''.join(buffer)
        if char.isdigit(): 
            number+=char
            break
        elif digits.get(current_buffer,0) != 0:
            number+= digits[current_buffer]
            break  
        elif digits.get(current_buffer[1:],0) != 0:
            number+= digits[current_buffer[1:]]
            break   
        elif digits.get(current_buffer[2:],0) != 0:
            number+= digits[current_buffer[2:]]
            break   
    return number  

def second_star(input: list)->int:
    sum = 0 
    for line in input:
        number = ''
        number += CheckForDigit(line) 
        number += CheckForDigit(line[::-1]) 
        sum+= int(number)
    return sum  

# Get the Input 
input = open('input.txt','r').read().splitlines()
input2 = copy.deepcopy(input)

sum = first_star(input)
print(f'First Star Result: {sum}')
sum = second_star(input2)
print(f'Second Star Result: {sum}')
