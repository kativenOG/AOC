MAX_SIZE = 100000
data = open("./input.txt","r").read().splitlines()

class Node():
    def __init__(self,name,parent_node):
        self.name = name
        self.size = 0 
        self.parent = parent_node if (parent_node!=None) else None
        self.childs = []
    
    def update(self,size) -> None:
        self.size+= size
        if self.parent!= None: self.parent.update(size)

    def insert(self,name) :
        self.childs.append(Node(name,self))
        return self.childs[len(self.childs)-1]
    
    def backward(self):
        return self.parent
    
    def sizes(self)-> int:
        counter = 0 if self.size > MAX_SIZE else 1 
        print(f"\nName: {self.name} Size: {self.size}")
        if len(self.childs)>0:
            for val in self.childs:
                counter += val.sizes() 
        return counter

    # def print(self):
    #     print(f"Name: {self.name} Size: {self.size} Father: {self.parent}")

# Ignora le linee dir <name> (crei solo con cd)  e ls !! 
def first_star(Data)-> int:
    root= Node("root",None)
    current_pos = root
    for line in Data:
        if line[0] == "$": # Commands: (non serve controllare ls)
            command = line.split()[1:]
            if command[0] == "cd":
                if command[1] == "..":
                    current_pos = current_pos.backward()
                else:
                    current_pos = current_pos.insert(command[1])
        else:
            command = line.split()
            if command[0].isnumeric():
                current_pos.update(int(command[0]))
         
    targets = root.sizes() 
    return targets

print(f"First Star: {first_star(data)}\n")
