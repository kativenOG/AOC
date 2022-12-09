forest = []
file = open("./input.txt","r").read().splitlines()
for line in file:
    forest.append(list(map(int,[*line])))
# print(forest)

def three_finder(Forest): # Non stai controllando le 2 collone e righe per le altre 2 direzioni 
    # visible = len(Forest)*2 + len(Forest[0])*2  # Border threes
    visible=0 
    bottom_limit = len(Forest)-1
    right_limit = len(Forest[0])-1
    top_max_array  = [ x for x in Forest[0] ]
    # Cycling through each line in the forest except the first and last (always visible)
    for line_n,line in enumerate(Forest):
        left_max = line[0]
        # Calculating a max for each direction and then comparing it to the value of the selected tree 
        for pos,tree in enumerate(line): # We Cycle left to right !!! (easy left maximum value)
            # Left:
            if pos!=0: 
                if tree > left_max: visible, left_max = visible +1, tree 
            else: visible+=1
            # Top:
            if line_n != 0: 
                if tree > top_max_array[pos]: visible, top_max_array[pos] = visible+1, tree 
            else: visible+=1
            # Right:
            if pos != right_limit:
                right_max = max(line[pos:]) 
                visible+= 1 if(tree>right_max) else 0
            else: visible+=1
            # Down: ( the most expensive (from a computational point of view))
            if line_n != bottom_limit:
                down_max =  max([Forest[val][pos] for val in range(line_n,len(Forest))])
                visible+= 1 if(tree>down_max) else 0
            else: visible+=1

    return visible


print(f"First Star: {three_finder(forest)}")
