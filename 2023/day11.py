def findDistance(k,l,galaxies,s,empty_rows,empty_cols,multiplier):
    a1,b1 = galaxies[k]
    a2,b2 = galaxies[l]
    result = abs(a1-a2) + abs(b1-b2)
    for r in empty_rows:
        if a1 < r < a2 or a2 < r < a1:
            result += (multiplier-1)
    for c in empty_cols:
        if b1 < c < b2 or b2 < c < b1:
            result += (multiplier-1)
    return result

def solve(s):
    galaxies = []
    empty_rows = []
    occupied_cols = set()
    n = len(s)
    m = len(s[0])
    for i in range(n):
        occupied = False
        for j in range(m):
            if s[i][j] == '#':
                galaxies.append((i,j))
                occupied = True
                occupied_cols.add(j)
        if not occupied:
            empty_rows.append(i)
    empty_cols = []
    for j in range(m):
        if j not in occupied_cols:
            empty_cols.append(j)
    part1 = 0
    part2 = 0
    for k in range(len(galaxies)):
        for l in range(k+1,len(galaxies)):
            part1 += findDistance(k,l,galaxies,s,empty_rows,empty_cols,2)
            part2 += findDistance(k,l,galaxies,s,empty_rows,empty_cols,1000000)
    return part1,part2

inp = """...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#....."""

lines = inp.split("\n")
print(solve(lines))
