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

fn main() {
    part1();
}
