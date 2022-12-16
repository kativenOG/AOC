from copy import deepcopy
data = open("./input.txt","r").read().splitlines()
valves = []


# Simple valve class
class valve():
    def __init__(self,name,flow_rate,tunnels) -> None:
        self.name = name
        self.flow_rate = flow_rate
        self.tunnels = tunnels
    
    def release(self,minutes)->int: # flow_rate is the pressure expected to be released per minute
        release_value = minutes * self.flow_rate
        return release_value
     
    def print(self)->None:
        print(f"Valve {self.name} has flow rate={self.flow_rate}; tunnels lead to valves {self.tunnels}")

# Creating a map of the valves 
for line in data:
    splitted_line = line.split()
    name = splitted_line[1]
    flow_rate = int(splitted_line[4].split("=")[1].split(";")[0])
    tunnels = list(map(lambda x: x.split(",")[0],splitted_line[9:]))
    valves.append(valve(name,flow_rate,tunnels))
    
# for valve in valves:
#     valve.print()


# Best first search kinda, no heuristic  
def optimal_path(environment):
    total_release = 0

    current_valve = environment[0]
    total_release =  current_valve.release(30)
    visited_valves = set()

    for time in range(29,0,-1):
        possibilities = []
        for action in current_valve.tunnels:
            target = [x for x in environment if x.name == action] 
            if target[0].name not in visited_valves: possibilities.append(deepcopy(target))

        current_valve =  sorted(possibilities,key=lambda x: int(x.flow_rate),reverse=True)[0]

        total_release += current_valve.release(time)  
        visited_valves.add(current_valve.name)

    return total_release

print(optimal_path(valves))
