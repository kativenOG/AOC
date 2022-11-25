file  = open("./input.txt","r")
data = file.read().split("\n")

# put the winning numbers inside an array 
appo= data[0].split(",")
winning_numbers = []
for i,number in enumerate(appo):
    winning_numbers.append(int(number))
    
data.pop(0) # erase numbers 
data.pop(0) # erase first empty line 

# Create an array of bingo boards 
Boards = []
while (len(data) != 1): # and (len(data) != 1 and data[0]=="") :
    matrix = [[[ 0 for _ in range(2)] for _ in range(5)] for _ in range(5)]
    for i,sent in enumerate(data):
        row = sent.split()
        for j,x in enumerate(row):
            matrix[i][j][0] = int(x)
        if (i==4): break 
    Boards.append(matrix)
    for i in range(6):
        if (len(data)!=0): data.pop(0)

def first_star(boards):
    # lets mark all the winning numbers inside the tables 
    for point in winning_numbers:
        for i,board in enumerate(boards):
            for j,line in enumerate(board):
                for k,number in enumerate(line):
                    column = 0 
                    row =0  
                    if ( number[0] == point ):
                        boards[i][j][k][1] =  1 # highlights the number 
                        for offset in range(5): # checks rows and columns  
                            column += boards[i][offset][k][1] 
                            row += boards[i][j][offset][1] 
                        if (number[0] == point and (column==5 or row==5)): # returns value if the column is comleted
                            not_called =0 
                            for l in boards[i]:
                                for n in l:
                                    if (n[1]==0): not_called+=n[0]
                            return(not_called*point)
                        break

def second_star(boards):
    # lets mark all the winning numbers inside the tables 
    for point in winning_numbers:
        # for i,board in enumerate(boards):
        i=0
        while i< len(boards):
            for j,line in enumerate(boards[i]):
                exit_board = False 
                for k,number in enumerate(line):
                    column = 0 
                    row =0  
                    if ( number[0] == point ):
                        boards[i][j][k][1] =  1 # highlights the number 
                        for offset in range(5): # checks rows and columns  
                            column += boards[i][offset][k][1] 
                            row += boards[i][j][offset][1] 
                        if (number[0] == point and (column==5 or row==5)): # returns value if the column is comleted
                            if(len(boards)==1):
                                not_called =0 
                                for l in boards[i]:
                                    for n in l:
                                        if (n[1]==0): not_called+=n[0]
                                return(not_called*point)
                            else:
                                boards.pop(i)
                                i-=1
                                exit_board= True 
                        break
                if (exit_board == True ): break 
            i+=1

print("First Star:",first_star(Boards))
print("Second Star:",second_star(Boards))

