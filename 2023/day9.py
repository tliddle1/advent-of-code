def diff(x):
    diffs = []
    for i in range(len(x)-1):
        diffs.append(x[i+1]-x[i])
    return diffs

def allZero(x):
    for i in x:
        if i != 0:
            return False
    return True

def extrapolateLeft(x):
    diffs = []
    d = x
    diffs.append(d)
    while not allZero(d):
        d = diff(d)
        diffs.append(d)
    inc = 0
    for i in range(len(diffs)-1,-1,-1):
        inc = diffs[i][0]-inc
    return inc

def extrapolateRight(x):
    diffs = []
    d = x
    diffs.append(d)
    while not allZero(d):
        d = diff(d)
        diffs.append(d)
    inc = 0
    for i in range(len(diffs)-1,-1,-1):
        inc = diffs[i][-1]+inc
    return inc

def solve(lines):
    part1 = 0
    part2 = 0
    for x in lines:
        part1 += extrapolateRight([int(a) for a in x.split()])
        part2 += extrapolateLeft([int(a) for a in x.split()])
        
    return part1, part2

inp = """0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45"""

lines = inp.split("\n")
print(solve(lines))
