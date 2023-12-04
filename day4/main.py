f = open("input.dat", "r")
lines = [x[:-1] for x in f.readlines()]
f.close()


def part1():
    ans = 0
    for line in lines:
        game = line.split(": ")[1]
        winning_nums = set(game.split(" | ")[0].split())
        i = 0
        for num in game.split(" | ")[1].split():
            if num in winning_nums:
                i += 1
        points = 2 ** (i - 1) if i > 0 else 0

        ans += points
    return ans


def part2():
    ans = 0
    for i, line in enumerate(lines):
        id = i + 1
        game = line.split(" ")[1]
        winning_nums = set(game.split(" | ")[0].split())
        i = 0
        for num in game.split(" | ")[1].split():
            if num in winning_nums:
                i += 1

    return ans


if __name__ == "__main__":
    print(part1())
    print(part2())
