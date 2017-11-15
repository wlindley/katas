const CUSTOM_DELIMETER_BEGIN:&str = "//";
const CUSTOM_DELIMETER_END:&str = "\n";

fn safe_parse(number:&str) -> u32 {
    match number.parse::<u32>() {
        Ok(v) => v,
        Err(_) => 0
    }
}

fn decide_delimeters(input:&str) -> (Vec<char>, &str) {
    if !input.starts_with(CUSTOM_DELIMETER_BEGIN) {
        return (vec![',', '\n'], input);
    }
    
    let lines:Vec<&str> = input.split(CUSTOM_DELIMETER_END).collect();
    let delimeter = lines[0].replace(CUSTOM_DELIMETER_BEGIN, "").replace(CUSTOM_DELIMETER_END, "");
    (delimeter.chars().collect(), lines[1])
}

pub fn add(numbers:&str) -> u32 {
    let (delimeters, numbers) = decide_delimeters(numbers);
    numbers.split(|c:char| delimeters.contains(&c))
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

    #[test]
    fn supports_delimeter_specification() {
        assert_eq!(12, add("//;\n4;6;2"));
    }
}
