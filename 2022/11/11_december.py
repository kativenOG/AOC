from copy import deepcopy

class Monkey:
    def __init__(self,text_input):  #,starting_items,worry_factor,dividendo,true_value,false_value):
        text_input = text_input.split("\n")
        self.counter = 0
        numbers = text_input[1].split()[2:]
        numbers = list(map(lambda x: x.replace(",",""),numbers))
        self.items = list(map(int,numbers))

        self.worry_factor = [0 for _ in range(2)] # [Operation , Value] 
        self.worry_factor[0] = text_input[2].split()[4] # Operation 
        self.worry_factor[1] = text_input[2].split()[5] # Value

        self.test_dividendo = int(text_input[3].split()[3])
        self.true_value = int(text_input[4].split()[5])
        self.false_value = int(text_input[5].split()[5])

    def show(self):
        print(f"MONKEY!!!\n  Items: {self.items}\n  Operation: {[*self.worry_factor]}\n  Test: divisible by: {self.test_dividendo}\n    If True: {self.true_value}\n    If False: {self.false_value}\n") 

    def throw(self):
        if len(self.items) != 0:
            return_array = [] # array filled with arrays of monkey-items
            for item in self.items:
                self.counter+=1
                inspected = item 
                inspected = round(inspected/3) # Worry Level Decreases 

                # Monkey Inspection 
                if(self.worry_factor[1]!="old"): number = int(self.worry_factor[1])  
                else: number = inspected
                if self.worry_factor[0] == "*": inspected = number * inspected # MULTIPLICATION
                else: inspected = number + inspected  # SUMMATION
            
                # Monkey Test:
                if inspected % self.test_dividendo == 0: return_array.append(list([self.true_value, inspected]))
                else: return_array.append(list([self.false_value, inspected]))
            self.items = []
            return return_array
        else: return  None
    
    def catch(self,item):
        self.items.append(item)


# Loading Data
data = open("./input.txt","r").read().split("\n\n")
monkeys = []
for monkey in data:
    monkeys.append(Monkey(monkey))


def first_star(Monkeys)->int:
    for _ in range(20): # Rounds
        for monkey in Monkeys:
            monkey_item_array = monkey.throw()
            if (monkey_item_array!=None):
                for object in monkey_item_array:
                    Monkeys[object[0]].catch(object[1])


    counters = sorted(list(map(lambda x: x.counter,Monkeys)),reverse=True)
    print(counters)
    return  counters[0]*counters[1]


monkeys2 = deepcopy(monkeys)
print(f"First Star: {first_star(monkeys)}")
# print(f"Second Star: {second_star(monkeys2)}")

