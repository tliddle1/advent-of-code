max_possible = {"green":13,"red":12,"blue":14}

def validGameNumber(line):
    info = line.split(": ")
    game_number = int(info[0].split(" ")[1])
    rounds = info[1].split("; ")
    result = 0
    for r in rounds:
        groups = r.split(", ")
        for g in groups:
            num,color = g.split(' ')
            num = int(num)
            if num > max_possible[color]:
                return 0
    return game_number
    

def solve(lines):
    result = 0
    for line in lines:
        result += validGameNumber(line)
    
    return result

lines = inp.split("\n")
print(solve(lines))
