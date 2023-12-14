inp = """O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#...."""

def transpose(x):
	return [[row[j] for row in x] for j in range(len(x[0]))]

def turnLeft(x):
	return [[row[j] for row in x] for j in range(len(x[0])-1,-1,-1)]

def turnRight(x):
    return [[row[j] for row in x[::-1]] for j in range(len(x[0]))]

def toString(rows):
    result = ""
    for row in rows:
        for c in row:
            result += c
    return result

def tilt(rows):
    for row in rows:
        ground = 0
        for i, col in enumerate(row):
            if col == '.':
                continue
            elif col == '#':
                ground = i+1
            elif col == 'O':
                row[ground] = 'O'
                if ground != i:
                    row[i] = '.'
                ground += 1
    return rows

west = inp.split("\n")
num_iterations = 1000000000
states = {}
for i in range(num_iterations):
    north = turnLeft(west)
    north_tilted = tilt(north)

    west = turnRight(north_tilted)
    west_tilted = tilt(west)

    south = turnRight(west_tilted)
    south_tilted = tilt(south)

    east = turnRight(south_tilted)
    east_tilted = tilt(east)

    west = turnLeft(turnLeft(east_tilted))
    key = toString(west)
    if key in states:
        cycle = i - states[key]
        sub = states[key]
        first_instance = (num_iterations-sub)%cycle+sub
        break
    else:
        states[key] = i

west = inp.split("\n")
states = {}
for i in range(first_instance):
    north = turnLeft(west)
    north_tilted = tilt(north)
    west = turnRight(north_tilted)
    west_tilted = tilt(west)
    south = turnRight(west_tilted)
    south_tilted = tilt(south)
    east = turnRight(south_tilted)
    east_tilted = tilt(east)
    west = turnLeft(turnLeft(east_tilted))

cols = turnLeft(west)
result = 0
for col in cols:
    inc = len(col)
    for row in col:
        if row == 'O':
            result += inc
        inc -= 1
print(result)
