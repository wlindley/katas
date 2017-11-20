pub fn run(max: u32) -> Vec<String> {
    (1..max+1).map(process).collect()
}

fn process(number: u32) -> String {
    if matches(3, number) && matches(5, number) {
        String::from("FizzBuzz")
    } else if matches(3, number) {
        String::from("Fizz")
    } else if matches(5, number) {
        String::from("Buzz")
    } else {
        format!("{}", number)
    }
}

fn matches(target: u32, number: u32) -> bool {
    0 == number % target || contains_digit(target, number)
}

fn contains_digit(digit: u32, number: u32) -> bool {
    let mut number = number;
    while number > 0 {
        let cur = number % 10;
        if digit == cur {
            return true;
        }
        number /= 10;
    }
    false
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn run_returns_numbers_from_1_to_max_inclusive_with_modifications() {
        assert_eq!(
            vec!["1", "2", "Fizz", "4", "Buzz"],
            run(5)
        );
    }

    #[test]
    fn process_returns_given_number_as_string() {
        assert_eq!("2", process(2));
    }

    #[test]
    fn process_returns_fizz_when_number_is_divisible_by_3() {
        assert_eq!("Fizz", process(3));
    }

    #[test]
    fn process_returns_buzz_when_number_is_divisible_by_5() {
        assert_eq!("Buzz", process(5));
    }

    #[test]
    fn process_returns_fizzbuzz_when_number_is_divisible_by_15() {
        assert_eq!("FizzBuzz", process(15));
    }

    #[test]
    fn process_returns_fizz_when_number_includes_3() {
        assert_eq!("Fizz", process(301));
    }

    #[test]
    fn process_returns_buzz_when_number_includes_5() {
        assert_eq!("Buzz", process(52));
    }
}
