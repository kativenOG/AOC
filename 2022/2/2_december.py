# A/X-ROCK  B/Y-PAPER  C/Z-SCIZZOR 
moveValues= {"Rock":1 ,"Paper":2 ,"Scizzor":3}
moveMapFirst = {"A":"Rock","B":"Paper","C":"Scizzor"}



data = open("input.txt","r").read().split("\n")
data.pop(len(data) -1 )

# First Star
moveMapSecond= {"X":"Rock","Y":"Paper","Z":"Scizzor"}

def game_result(first,second):
    if first == "Rock":
        return True if second=="Paper" else False 
    if first == "Paper":
        return True if second=="Scizzor" else False 
    if first == "Scizzor":
        return True if second=="Rock" else False 
    else: return None

def first_star(Data):
    score = 0 
    for game in Data:
        moves = game.split()
        first,second = moveMapFirst[moves[0]],moveMapSecond[moves[1]]
        score += 3 if ( first == second ) else 6 if (game_result(first,second)) else 0 
        score += moveValues[second]
    return score


# Second Star 
resultMapSecond = {"X":"Lose","Y":"Draw","Z":"Win"} 

def move_extractor(first,result):
    if result == "Win":
        return "Rock" if (first == "Scizzor") else "Paper" if (first == "Rock") else "Scizzor"
    else:
        return "Rock" if (first == "Paper") else "Paper" if (first == "Scizzor") else "Scizzor"

def second_star(Data):
    score = 0 
    for game in Data:
        moves = game.split()
        first,second = moveMapFirst[moves[0]],resultMapSecond[moves[1]]
        score += 3 if ( second== "Draw" ) else 6 if (second == "Win") else 0 
        extracted = first if second == "Draw" else move_extractor(first,second)
        score += moveValues[extracted]
    return score

print(f"First star Result: {first_star(data)}\nSecond star Result: {second_star(data)}")
