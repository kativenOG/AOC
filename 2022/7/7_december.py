MAX_SIZE = 100000
UPDATE_SIZE= 30000000  
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
        child = Node(name,self)
        self.childs.append(child)
        return child
    
    def sizes(self)-> int:
        total = 0 if self.size > MAX_SIZE else self.size 
        # print(f"\nName: {self.name} Size: {self.size}")
        for val in self.childs:
            total+= val.sizes() 
        return total 
    
    def update_dir(self,space_to_free)->int:
        target = 70000000
        if self.size> space_to_free:
            target = self.size
        for val in self.childs:
                target = min(target,val.update_dir(space_to_free)) 
        return target

    def print(self):
        print(f"\nName: {self.name} Size: {self.size} Father: {self.parent}")

# Ignora le linee dir <name> (crei solo con cd)  e ls !! 
def results(Data):
    root= Node("root",None)
    current_pos = root
    for line in Data:

        if line[0] == "$": # Commands: (non serve controllare ls)
            command = line.split()[1:]
            if command[0] == "cd":
                if command[1] == "..":
                    current_pos = current_pos.parent
                else:
                    current_pos = current_pos.insert(command[1])

        else:
            command = line.split()
            if command[0].isnumeric():
                current_pos.update(int(command[0]))
         
    unused = 70000000 - root.size 
    missing_space = 30000000 - unused
    return root.sizes(), root.update_dir(missing_space)

targets,update_size =  results(data)
print(f"\nFirst Star: {targets}\nSecond Star: {update_size}")
