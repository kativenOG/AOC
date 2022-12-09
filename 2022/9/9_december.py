import copy 
commands = open("./input.txt","r").read().splitlines()

def rope_problem(Commands):
    rope = [[0,0] for _ in range(10)]
    t_visited_pos1 = set()
    t_visited_pos9 = set()

    for command in Commands:
        direction , steps = command.split()[0],int(command.split()[1])
        sum = True if (direction=="U" or direction=="R") else False
        axis = 1 if (direction=="U" or direction=="D") else 0 
        for _ in range(steps):
            if sum: rope[0][axis] += 1 
            else: rope[0][axis] -=1
            for knot in range(1,10):
                x_diff, y_diff = rope[knot-1][0]-rope[knot][0],rope[knot-1][1]-rope[knot][1]
                if (x_diff> 1 or x_diff<-1) and (y_diff == 1 or y_diff==-1):
                    rope[knot][0] += int(x_diff/abs(x_diff))
                    rope[knot][1] += y_diff
                elif (y_diff> 1 or y_diff<-1) and (x_diff == 1 or x_diff==-1):
                    rope[knot][0] += x_diff 
                    rope[knot][1] += int(y_diff/abs(y_diff))
                else:
                    if (x_diff>1): rope[knot][0]+=1 # X coordinate change
                    elif (x_diff<-1): rope[knot][0]-=1 
                    if (y_diff>1): rope[knot][1]+=1 # Y coordinate change
                    elif (y_diff<-1): rope[knot][1]-=1
                t_visited_pos1.add(tuple(rope[1]))
                t_visited_pos9.add(tuple(rope[9]))
    return len(t_visited_pos1),len(t_visited_pos9)

# new_rope = copy.deepcopy(commands)
first,second = rope_problem(commands)
print(f"First Star: {first}\n")
print(f"Second Star: {second}\n")
