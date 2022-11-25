file  = open("./input.txt","r")
data = file.read().split("\n")

# put the winning numbers inside an array 
appo= data[0].split(",")
winning_numbers = []
for i,number in enumerate(appo):
    winning_numbers.append(int(number))
print(winning_numbers)
    
data.pop(0) # erase numbers 
data.pop(0) # erase first empty line 

# Create an array of bingo boards 
boards = []
while (len(data) != 1): # and (len(data) != 1 and data[0]=="") :
    matrix = [[[ 0 for _ in range(2)] for _ in range(5)] for _ in range(5)]
    for i,sent in enumerate(data):
        row = sent.split()
        for j,x in enumerate(row):
            matrix[i][j][0] = int(x)
        if (i==4): break 
    # print("\n",matrix)
    boards.append(matrix)
    for i in range(6):
        if (len(data)!=0): data.pop(0)
# print(boards)

# sei un ritardato così guardi tutti i numeri su ogni tabella porco dio 
def first_star():
    # lets mark all the winning numbers inside the tables 
    for i,board in enumerate(boards):
        for j,line in enumerate(board):
            for k,number in enumerate(line):
                for point in winning_numbers:
                    column = 0 
                    row =0  
                    if ( number[0] == point ):
                        boards[i][j][k][1] =  1 # highlights the number 
                        for offset in range(5): # checks rows and columns  
                            column += boards[i][offset][k][1] 
                            row += boards[i][j][offset][1] 
                        if (number[0] == point and (column==5 or row==5)): # returns value if the column is comleted
                            print(i,point)
                            return(i*point)
                        break
    return None

print(first_star())

# print(boards)
