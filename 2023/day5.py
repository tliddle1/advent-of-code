def convert(src,m):
    dst = []
    for x in m.split('\n')[1:]:
        mappings = x.split()
        destination_start = int(mappings[0])
        source_start = int(mappings[1])
        range_length = int(mappings[2])
        for i in src:
            if source_start <= i < source_start+range_length:
                dst.append((destination_start+i-source_start))
                src.remove(i)
    return dst + src

def solve(lines):
    seeds, seeds_soil, soil_fert, fert_water, water_light, light_temp, temp_humid, humid_loc = lines
    seeds = seeds.split()[1:]
    seeds = [int(x) for x in seeds]
    result = []
    for s in seeds:
        soils = convert([s],seeds_soil)
        ferts = convert(soils,soil_fert)
        waters = convert(ferts,fert_water)
        lights = convert(waters,water_light)
        temps = convert(lights,light_temp)
        humids = convert(temps,temp_humid)
        locs = convert(humids,humid_loc)
        result.append(locs[0])
    return min(result)


inp = """seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4"""

lines = inp.split("\n\n")
print(solve(lines))
