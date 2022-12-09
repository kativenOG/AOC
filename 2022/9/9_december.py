import copy 
commands = open("./input.txt","r").read().splitlines()

def first_star(Commands):
    t,h = [0,0],[0,0] # s = 0,0
    t_visited_pos = set()
    t_visited_pos.add(tuple(t))
    # t_visited_pos = []
    # t_visited_pos.append(tuple(t))
    for command in Commands:
        # Parsing the command
        direction , steps = command.split()[0],int(command.split()[1])
        sum = True if (direction=="U" or direction=="R") else False
        axis = 1 if (direction=="U" or direction=="D") else 0 
        for _ in range(steps):
            # Moving the head
            if sum: h[axis] += 1 
            else: h[axis] -=1

            # Moving the tail: im not moving diagonally when needed 
            x_diff, y_diff = h[0]-t[0],h[1]-t[1]
            if (x_diff> 1 or x_diff<-1) and (y_diff == 1 or y_diff==-1):
                t[0] += int(x_diff/abs(x_diff))
                t[1] += y_diff
            elif (y_diff> 1 or y_diff<-1) and (x_diff == 1 or x_diff==-1):
                t[0] += x_diff 
                t[1] += int(y_diff/abs(y_diff))
            else:
                if (x_diff>1): t[0]+=1 # X coordinate change
                elif (x_diff<-1): t[0]-=1 
                if (y_diff>1): t[1]+=1 # Y coordinate change
                elif (y_diff<-1): t[1]-=1 

            # Updating tail previous positions array:
            t_visited_pos.add(tuple(t))
            # print(f"Tail: {tuple(t)}, Head: {tuple(h)}, DistanceX {h[0]-t[0]}, Distance Y: {h[1]-t[1]}")

    return len(t_visited_pos)

def second_star(Commands):
    rope = [[0,0] for _ in range(10)]
    t_visited_pos = set()

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
                t_visited_pos.add(tuple(rope[9]))
    return len(t_visited_pos)
new_rope = copy.deepcopy(commands)
print(f"First Star: {first_star(commands)}\n")
print(f"Second Star: {second_star(new_rope)}\n")
