use std::fs;

fn part1() {
    let mut result = 0;
    for line in fs::read_to_string("input.dat").unwrap().lines() {
        let nums = line
            .chars()
            .filter(|c| c.is_ascii_digit())
            .map(|c| c as u8 - b'0')
            .collect::<Vec<_>>();
        result +=
            (*nums.iter().nth(0).unwrap() as i32) * 10 + (*nums.iter().last().unwrap() as i32);
    }
    println!("{}", result);
}

fn part2() {
    // Functional programming... yay
    let result: u32 = fs::read_to_string("input.dat")
        .unwrap()
        .lines()
        .filter(|line| !line.is_empty())
        .map(|line| {
            line.to_string()
                .replace("one", "one1one")
                .replace("two", "two2two")
                .replace("three", "three3three")
                .replace("four", "four4four")
                .replace("five", "five5five")
                .replace("six", "six6six")
                .replace("seven", "seven7seven")
                .replace("eight", "eight8eight")
                .replace("nine", "nine9nine")
        })
        .map(|line| {
            line.chars()
                .filter_map(|c| c.to_digit(10))
                .collect::<Vec<u32>>()
        })
        .map(|vec| 10 * vec.first().unwrap() + vec.last().unwrap())
        .sum();

    println!("{}", result);
}

fn main() {
    part1();
    part2();
}
