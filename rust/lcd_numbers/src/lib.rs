const NUMBERS: &'static [&'static str] = &[
" _ 
| |
|_|",
"   
  |
  |",
" _ 
 _|
|_ ",
" _ 
 _|
 _|",
 "   
|_|
  |",
" _ 
|_ 
 _|",
" _ 
|_ 
|_|",
" _ 
  |
  |",
" _ 
|_|
|_|",
" _ 
|_|
 _|"
 ];

pub fn lcdify(number: u32) -> String {
    let digits: Vec<_> = digits(number).iter().cloned().map(visualize_digit).collect();
    combine(digits)
}

fn digits(mut number: u32) -> Vec<u32> {
    let mut digits = vec![];
    loop {
        digits.insert(0, number % 10);
        number /= 10;
        if 1 > number {
            break;
        }
    }
    digits
}

fn visualize_digit(digit: u32) -> String {
    NUMBERS[digit as usize].to_string()
}

fn combine(digits: Vec<String>) -> String {
    let mut result = String::new();
    for row in 0..3 {
        for digit in &digits {
            let rows: Vec<_> = digit.split('\n').collect();
            result.push_str(rows[row]);
        }
        if 2 != row {
            result.push('\n');
        }
    }
    result
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn one() {
        verify_digits(1, "   \n  |\n  |");
    }

    #[test]
    fn two() {
        verify_digits(2, " _ \n _|\n|_ ");
    }

    #[test]
    fn three() {
        verify_digits(3, " _ \n _|\n _|");
    }

    #[test]
    fn four() {
        verify_digits(4, "   \n|_|\n  |");
    }

    #[test]
    fn five() {
        verify_digits(5, " _ \n|_ \n _|");
    }

    #[test]
    fn six() {
        verify_digits(6, " _ \n|_ \n|_|");
    }

    #[test]
    fn seven() {
        verify_digits(7, " _ \n  |\n  |");
    }

    #[test]
    fn eight() {
        verify_digits(8, " _ \n|_|\n|_|");
    }

    #[test]
    fn nine() {
        verify_digits(9, " _ \n|_|\n _|");
    }

    #[test]
    fn zero() {
        verify_digits(0, " _ \n| |\n|_|");
    }

    #[test]
    fn ten() {
        verify_digits(10, "    _ \n  || |\n  ||_|");
    }

    #[test]
    fn twenty_five() {
        verify_digits(25, " _  _ \n _||_ \n|_  _|");
    }

    fn verify_digits(digit: u32, visualization: &str) {
        assert_eq!(visualization.to_string(), lcdify(digit));
        println!("{}:\n{}", digit, visualization);
    }
}
