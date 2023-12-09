inp = "3,4,3,1,2"

def iterate(d):
    n_new = d[0]
    for i in range(8):
        d[i] = d[i+1]
    d[8] = n_new
    d[6] += n_new
    return d

def solve(lanternfish, iterations):
    d = {0:0,1:0,2:0,3:0,4:0,5:0,6:0,7:0,8:0}
    for f in lanternfish:
        d[int(f)] += 1
    for i in range(iterations):
        d = iterate(d)
    return sum(d.values())

print(solve(inp.split(","),80))
print(solve(inp.split(","),256))
