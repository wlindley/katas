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
    let index = number as usize;
    NUMBERS[index].to_string()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn one() {
        verify_digit(1, "   \n  |\n  |");
    }

    #[test]
    fn two() {
        verify_digit(2, " _ \n _|\n|_ ");
    }

    #[test]
    fn three() {
        verify_digit(3, " _ \n _|\n _|");
    }

    #[test]
    fn four() {
        verify_digit(4, "   \n|_|\n  |");
    }

    #[test]
    fn five() {
        verify_digit(5, " _ \n|_ \n _|");
    }

    #[test]
    fn six() {
        verify_digit(6, " _ \n|_ \n|_|");
    }

    #[test]
    fn seven() {
        verify_digit(7, " _ \n  |\n  |");
    }

    #[test]
    fn eight() {
        verify_digit(8, " _ \n|_|\n|_|");
    }

    #[test]
    fn nine() {
        verify_digit(9, " _ \n|_|\n _|");
    }

    #[test]
    fn zero() {
        verify_digit(0, " _ \n| |\n|_|");
    }

    fn verify_digit(digit: u32, visualization: &str) {
        assert_eq!(visualization.to_string(), lcdify(digit));
        println!("{}:\n{}", digit, visualization);
    }
}
