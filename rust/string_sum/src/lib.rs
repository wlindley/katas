fn delimeter(c:char) -> bool {
    ',' == c || '\n' == c
}

fn safe_parse(number:&str) -> u32 {
    match number.parse::<u32>() {
        Ok(v) => v,
        Err(_) => 0
    }
}

pub fn add(numbers:&str) -> u32 {
    numbers.split(delimeter)
           .map(safe_parse)
           .sum()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn returns_0_for_empty_string() {
        assert_eq!(0, add(""));
    }

    #[test]
    fn returns_1_for_string_with_1() {
        assert_eq!(1, add("1"));
    }

    #[test]
    fn returns_3_for_string_with_1_and_2() {
        assert_eq!(3, add("1,2"));
    }

    #[test]
    fn supports_newline_delimeter() {
        assert_eq!(6, add("2\n3,1"));
    }
}
