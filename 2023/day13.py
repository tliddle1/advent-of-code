def transpose(x):
	return [[row[j] for row in x] for j in range(len(x[0]))]

def distance(a,b):
    dist = 0
    for c1, c2 in zip(a, b):
        if c1 != c2:
            dist += 1
    return dist

def calc(group, dist):
    for i in range(1, len(group)):
        for i in range(1, len(group)):
            count = 0
            for top, bottom in zip(group[i-1::-1], group[i:]):
                count += distance(top,bottom)
            if count == dist:
                return i
    else: 
        return 0
        
def solve(groups, dist):
    result = 0
    for group in groups:
        group = group.split("\n")
        r = calc(group,dist)
        c = calc(transpose(group),dist)
        result += c+ r * 100
    return result

groups = inp.split('\n\n')

print(solve(groups, 0))
print(solve(groups, 1))
