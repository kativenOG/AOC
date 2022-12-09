data = open("./input.txt","r").read()
# Checcka tutti i numeri succesivi nella finestra partendo dal primo, esce subito alla prima uguaglianza essendo all() una serie di and
# Se avessi le palle lo riscriveresti con ricorsione
def checker(slice): 
    selected = slice[0]
    if all(selected != val for val in slice[1:]):
        return True 
    else: return False

def first_star(Data):
    for i in range(len(Data)):
        print("\n")
        if i>2:
            if all( checker(data[(i-pos) :i+1]) for pos in [3,2,1]):
                return i+1
    return None

def second_star(Data):
    for i in range(len(Data)):
        if i>12:
            if all( checker(data[(i-pos) :i+1]) for pos in [13,12,11,10,9,8,7,6,5,4,3,2,1]):
                return i+1
    return None

print(f"First Star: {first_star(data)}\nSecond Star: {second_star(data)}")
