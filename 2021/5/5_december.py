# Loading input data in pairs(start-finish) of pairs (x-y) of coordinates
import copy 
file = open("./input.txt","r")
data = file.read().split("\n")
data.pop(len(data)-1)
coordinates = []
max_x,max_y = 0,0

# add vertical line support 
for line in data:
    coordinate = line.split(" -> ") 
    start = coordinate[0].split(",")
    finish = coordinate[1].split(",")
    if ((start[0]==finish[0]) or (start[1] == finish[1])):
        max_x = int(start[0]) if (int(start[0]) > max_x) else int(finish[0]) if (int(finish[0]) > max_x) else max_x
        max_y = int(start[1]) if (int(start[1]) > max_y) else int(finish[1]) if (int(finish[1]) > max_y) else max_y
        f,s = [int(start[0]),int(start[1])],[int(finish[0]),int(finish[1])]
        command = [f,s]
        coordinates.append(command)

def first_star(commands):
    counter = 0 
    field_matrix = [[0 for _ in range(max_x)] for _ in range(max_y)]

    for command in commands:
        if (command[0][0]==command[1][0]): # y lines
            x = command[0][0]
            start_y = command[0][1] if (command[0][1] < command[1][1] ) else command[1][1] -1
            y_diff = abs(command[0][1]-command[1][1]) + start_y
            print(start_y,(start_y<y_diff),(y_diff-start_y))
            # if(start_y == 990): print("MAX")
            while(start_y < y_diff):
                start_y+=1
                field_matrix[start_y][x]+=1
                
        else: # x lines
            y = command[0][1]
            start_x = command[0][0] if (command[0][0] < command[1][0] ) else command[1][0] -1
            x_diff = abs(command[0][0] - command[1][0]) + start_x
            print(start_x,(start_x<x_diff),(x_diff-start_x))
            if(start_x == 990): print("MAX")
            while(start_x < x_diff):
                start_x+=1
                field_matrix[y][start_x]+=1

    # output = open("output.txt","w")  
    # output.write(str(field_matrix))
    for x in field_matrix:
        for y in x:
            if y > 1 and y<5: counter+=1
    return counter 

print("First Star result:",first_star(coordinates))
