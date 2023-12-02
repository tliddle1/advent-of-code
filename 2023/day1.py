inp = """Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"""

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
