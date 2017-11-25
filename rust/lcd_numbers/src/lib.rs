pub struct LcdPrinter {
    char_width: u32,
    char_height: u32
}

impl LcdPrinter {
    pub fn default() -> LcdPrinter {
        Self::new(3, 3)
    }

    pub fn new(width: u32, height: u32) -> LcdPrinter {
        if 3 > width {
            panic!("Width is less than minimum of 3");
        }
        if 3 > height {
            panic!("Height is less than minimum of 3");
        }
        LcdPrinter {
            char_width: width,
            char_height: height
        }
    }

    pub fn lcdify(&self, number: u32) -> String {
        let digit_iterators = Digitizer::new(number).map(|d| DigitVisualizer::new(d, self.char_width, self.char_height)).collect();
        self.combine(digit_iterators)
    }

    fn combine<T>(&self, mut digits: Vec<T>) -> String
    where T : Iterator<Item = char> {
        println!(" ");
        let mut result = String::new();
        for row in 0..self.char_height {
            for digit in &mut digits {
                let row: String = digit.take(self.char_width as usize).collect();
                println!("Adding row: {}", row);
                result.push_str(&row);
            }
            if self.char_height-1 != row {
                result.push('\n');
            }
        }
        result
    }
}

struct Digitizer {
    digits: Vec<u32>
}

impl Digitizer {
    fn new(mut number: u32) -> Digitizer {
        let mut digits = vec![];
        loop {
            digits.push(number % 10);
            number /= 10;
            if 1 > number {
                break;
            }
        }
        Digitizer {
            digits : digits
        }
    }
}

impl Iterator for Digitizer {
    type Item = u32;

    fn next(&mut self) -> Option<Self::Item> {
        self.digits.pop()
    }
}

struct DigitVisualizer {
    digit: u32,
    width: u32,
    height: u32,
    index: u32
}

impl DigitVisualizer {
    fn new(digit: u32, width: u32, height: u32) -> DigitVisualizer {
        DigitVisualizer { digit, width, height, index: 0 }
    }

    fn x(&self) -> u32 {
        self.index % self.width
    }

    fn y(&self) -> u32 {
        self.index / self.width
    }

    fn next_zero(&self) -> char {
        let x = self.x();
        let y = self.y();
        if (0 == x && 0 != y) || (self.width-1 == x && 0 != y) {
            '|'
        } else if 0 == y && 0 != x && self.width-1 != x {
            '_'
        } else if self.height-1 == y {
            '_'
        } else {
            ' '
        }
    }

    fn next_one(&self) -> char {
        let x = self.x();
        let y = self.y();
        if 0 != y && self.width-1 == x {
            '|'
        } else {
            ' '
        }
    }

    fn next_two(&self) -> char {
        let x = self.x();
        let y = self.y();
        if (0 == y || self.height-1 == y) && (0 != x && self.width-1 != x) {
            '_'
        } else if self.height/2 == y {
            if 0 == x {
                ' '
            } else if self.width-1 == x {
                '|'
            } else {
                '_'
            }
        } else if self.width-1 == x && self.height/2 >= y && 0 < y {
            '|'
        } else if 0 == x && self.height/2 < y {
            '|'
        } else {
            ' '
        }
    }

    fn next_three(&self) -> char {
        let x = self.x();
        let y = self.y();
        if (0 != x && self.width-1 != x) && (0 == y || self.height/2 == y || self.height-1 == y) {
            '_'
        } else if self.width-1 == x && 0 != y {
            '|'
        } else {
            ' '
        }
    }

    fn next_four(&self) -> char {
        let x = self.x();
        let y = self.y();
        if 0 < y && self.width-1 == x {
            '|'
        } else if 0 < y && self.height/2 >= y && 0 == x {
            '|'
        } else if self.height/2 == y {
            '_'
        } else {
            ' '
        }
    }

    fn next_five(&self) -> char {
        let x = self.x();
        let y = self.y();
        if (0 == y || self.height-1 == y) && (0 != x && self.width-1 != x) {
            '_'
        } else if self.height/2 == y {
            if self.width-1 == x {
                ' '
            } else if 0 == x {
                '|'
            } else {
                '_'
            }
        } else if 0 == x && self.height/2 >= y && 0 < y {
            '|'
        } else if self.width-1 == x && self.height/2 < y {
            '|'
        } else {
            ' '
        }
    }

    fn next_six(&self) -> char {
        ' '
    }

    fn next_seven(&self) -> char {
        ' '
    }

    fn next_eight(&self) -> char {
        ' '
    }

    fn next_nine(&self) -> char {
        ' '
    }
}

impl Iterator for DigitVisualizer {
    type Item = char;

    fn next(&mut self) -> Option<Self::Item> {
        if self.index >= self.width * self.height {
            return None;
        }
        let symbol = if 1 == self.digit {
            self.next_one()
        } else if 2 == self.digit {
            self.next_two()
        } else if 3 == self.digit {
            self.next_three()
        } else if 4 == self.digit {
            self.next_four()
        } else if 5 == self.digit {
            self.next_five()
        } else if 6 == self.digit {
            self.next_six()
        } else if 7 == self.digit {
            self.next_seven()
        } else if 8 == self.digit {
            self.next_eight()
        } else if 9 == self.digit {
            self.next_nine()
        } else {
            self.next_zero()
        };
        self.index += 1;
        Some(symbol)
    }
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

    #[test]
    fn zero_in_5_by_4() {
        verify_custom_digits(0, 5, 4, " ___ \n|   |\n|   |\n|___|");
    }

    #[test]
    fn four_in_3_by_4() {
        verify_custom_digits(4, 3, 4, "   \n| |\n|_|\n  |");
    }

    #[test]
    fn four_in_3_by_5() {
        verify_custom_digits(4, 3, 5, "   \n| |\n|_|\n  |\n  |");
    }

    #[test]
    #[should_panic(expected = "Width is less than minimum of 3")]
    fn cannot_make_printer_with_width_less_than_3() {
        LcdPrinter::new(2, 4);
    }

    #[test]
    #[should_panic(expected = "Height is less than minimum of 3")]
    fn cannot_make_printer_with_height_less_than_3() {
        LcdPrinter::new(5, 1);
    }

    fn verify_digits(digit: u32, visualization: &str) {
        let printer = LcdPrinter::default();
        assert_eq!(visualization.to_string(), printer.lcdify(digit));
        println!("{}:\n{}", digit, visualization);
    }

    fn verify_custom_digits(digit: u32, width: u32, height: u32, visualization: &str) {
        let printer = LcdPrinter::new(width, height);
        assert_eq!(visualization.to_string(), printer.lcdify(digit));
        println!("{}:\n{}", digit, visualization);
    }
}
