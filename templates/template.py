import argparse


def part1(input):
    return -1


def part2(input):
    return -1


if __name__ == "__main__":
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument(
        "--sample",
        "-s",
        action="store_true",
        help="Use sample.txt (manually filled in)",
    )
    args = arg_parser.parse_args()
    input_file = "input.txt"
    if args.sample:
        print("USING SAMPLE INPUT FILE")
        input_file = "sample.txt"

    parsed_input = None  # Use some parsing method here

    print("Part 1:", part1(parsed_input))
    print("Part 2:", part2(parsed_input))
