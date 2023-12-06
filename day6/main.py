f = open("input.dat", "r")
lines = [x[:-1] for x in f.readlines()]
f.close()


def part1():
    times = [int(x) for x in lines[0].split(":")[1].strip().split()]
    distances = [int(x) for x in lines[1].split(":")[1].strip().split()]
    ans = 1
    for time, dist in zip(times, distances):
        ways = 0
        for i in range(time + 1):
            if i * (time - i) > dist:
                ways += 1
        ans *= ways
    return ans


def part2():
    time = int("".join(lines[0].split(":")[1].strip().split()))
    distance = int("".join(lines[1].split(":")[1].strip().split()))

    ways = 0
    for i in range(time + 1):
        if i * (time - i) > distance:
            ways += 1
    return ways


if __name__ == "__main__":
    print(part1())
    print(part2())
