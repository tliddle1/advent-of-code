inp = """467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598.."""

digits = ['1','2','3','4','5','6','7','8','9','0']

def parseNum(line,x):
    num = ""
    EON = False # End of Number
    while not EON:
        num += line[x]
        x += 1
        if x == len(line) or line[x] not in digits:
            EON = True
    return int(num), len(num)

# Part 1

def isAdj(i,j,symbol_locations):
    # This works because i,j will never be a symbol since it's in the number
    for x in range(i-1,i+2):
        for y in range(j-1,j+2):
            if (x,y) in symbol_locations:
                return True
    return False

def isPart(x,y,length,symbol_locations):
    for i in range(x,length+x):
        if isAdj(i,y,symbol_locations):
            return True
    return False

def solve(lines):
    symbol_locations = []
    
    y = 0
    for line in lines:
        x = 0
        for c in line:
            if c != '.' and c not in digits: # is symbol
                symbol_locations.append((x,y))
            x += 1
        y += 1

    result = 0
    parsingNumber = False
    y = 0
    for line in lines:
        x = 0
        for c in line:
            if c in digits and not parsingNumber:
                number, length = parseNum(line,x)
                if isPart(x,y,length, symbol_locations):
                    result += number
                parsingNumber = True
            elif c not in digits:
                parsingNumber = False
            x += 1
        y += 1
    return result

lines = inp.split("\n")
print(solve(lines))

# Part 2

d = {}
count = {}

def isAdj(i,j,gear_locations):
    for x in range(i-1,i+2):
        for y in range(j-1,j+2):
            if (x,y) in gear_locations:
                return True, (x,y)
    return False, None

def recordGearRatio(x,y,length,gear_locations, num):
    for i in range(x,length+x):
        isGear, coord = isAdj(i,y,gear_locations)
        if isGear:
            k = str(coord[0])+','+str(coord[1])
            if k in d.keys():
                if count[k] < 2:
                    d[k] *= num
                    count[k] += 1
                else:
                    d[k] = 0
            else:
                d[k] = num
                count[k] = 1
            return

def solve(lines):
    gear_locations = []
    y = 0
    for line in lines:
        x = 0
        for c in line:
            if c == "*": # is gear
                gear_locations.append((x,y))
            x += 1
        y += 1

    y = 0
    inNum = False
    for line in lines:
        x = 0
        for c in line:
            if c in digits and not inNum:
                number, length = parseNum(line,x)
                recordGearRatio(x,y,length,gear_locations,number)
                inNum = True
            elif c not in digits:
                inNum = False
            x += 1
        y += 1
    
    result = 0
    for k, v in d.items():
        if count[k] == 2:
            result += v
    return result

lines = inp.split("\n")
print(solve(lines))
