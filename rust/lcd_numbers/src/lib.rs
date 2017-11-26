const NUMBERS: &'static [&'static str] = &[
" ___ |   ||   ||   ||___|", //0
"         |    |    |    |", //1
" ___     | ___||    |___ ", //2
" ___     | ___|    | ___|", //3
"     |   ||___|    |    |", //4
" ___ |    |___     | ___|", //5
" ___ |    |___ |   ||___|", //6
" ___     |    |    |    |", //7
" ___ |   ||___||   ||___|", //8
" ___ |   ||___|    | ___|"  //9
 ];
 const DEFAULT_SIZE: u32 = 5;

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
        let mut result = String::new();
        for row in 0..self.char_height {
            for digit in &mut digits {
                let row: String = digit.take(self.char_width as usize).collect();
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

    fn corrected_index(&self) -> u32 {
        let corrected_x = correct_coord(self.x(), self.width);
        let corrected_y = correct_coord(self.y(), self.height);
        let corrected_index = (DEFAULT_SIZE * corrected_y) + corrected_x;
        corrected_index
    }
}

fn correct_coord(coord: u32, size: u32) -> u32 {
    if 0 == coord {
        0
    } else if size/2 == coord {
        DEFAULT_SIZE / 2
    } else if size-1 == coord {
        DEFAULT_SIZE - 1
    } else if size/2 > coord {
        1
    } else {
        DEFAULT_SIZE - 2
    }
}

impl Iterator for DigitVisualizer {
    type Item = char;

    fn next(&mut self) -> Option<Self::Item> {
        if self.index >= self.width * self.height {
            return None;
        }
        let digit_viz = NUMBERS[self.digit as usize];
        let symbol = digit_viz.chars().nth(self.corrected_index() as usize).unwrap();
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
