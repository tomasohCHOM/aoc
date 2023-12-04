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


def play(line_idx):
    res = 0
    game = lines[line_idx].split(": ")[1]
    winning_nums = set(game.split(" | ")[0].split())
    i = 0
    for num in game.split(" | ")[1].split():
        if num in winning_nums:
            i += 1
    for j in range(i):
        res += play(line_idx + j + 1)
    return res + 1


def part2():
    ans = 0
    for i in range(len(lines)):
        ans += play(i)
    return ans


if __name__ == "__main__":
    print(part1())
    print(part2())
