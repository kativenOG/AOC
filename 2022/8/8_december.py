forest = []
file = open("./input.txt","r").read().splitlines()
for line in file:
    forest.append(list(map(int,[*line])))
# print(forest)

def three_finder(Forest): # Non stai controllando le 2 collone e righe per le altre 2 direzioni 
    # visible = len(Forest)*2 + len(Forest[0])*2  # Border threes
    visible=(len(Forest)+len(Forest[0]))*2 - 4
    bottom_limit,right_limit = len(Forest)-1,len(Forest[0])-1
    
    top_max_array  = [ x for x in Forest[0] ]
    # Cycling through each line in the forest except the first and last (always visible)
    for line_n,line in enumerate(Forest):
        left_max = line[0]
        # Calculating a max for each direction and then comparing it to the value of the selected tree 
        for pos,tree in enumerate(line): # We Cycle left to right !!! (easy left maximum value)
            if pos != 0 and pos!=right_limit and line_n!=0 and line_n!=bottom_limit: # not counting borders 
                right_max = max(line[int(pos+1):]) 
                down_max =  max([Forest[val][pos] for val in range(int(line_n+1),len(Forest))]) 
                print(f"Tree: {tree}, TopMax:{top_max_array[pos]}, BottomMax: {down_max}, LeftMax:{left_max}, RightMax: {right_max}")
                # Left:
                if (tree > left_max):  
                    print("Increased on Left\n")
                    visible, left_max = visible +1, tree 
                # Top:
                elif (tree > top_max_array[pos]): 
                    print("Increased on Top\n")
                    visible, top_max_array[pos] = visible+1, tree 
                # Right:
                elif (tree > right_max): 
                    print("Increased on Right\n")
                    visible+= 1 
                # Down: 
                elif(tree>down_max):   
                    print("Increased on Down\n")
                    visible+= 1

    return visible


print(f"First Star: {three_finder(forest)}")
