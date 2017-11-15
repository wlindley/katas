const CUSTOM_DELIMETER_BEGIN:&str = "//";
const CUSTOM_DELIMETER_END:&str = "\n";

pub fn add(numbers:&str) -> u32 {
    let (delimeters, numbers) = decide_delimeters(numbers);
    numbers.split(|c:char| delimeters.contains(&c))
           .map(safe_parse)
           .sum()
}

fn decide_delimeters(input:&str) -> (Vec<char>, &str) {
    if !input.starts_with(CUSTOM_DELIMETER_BEGIN) {
        return (vec![',', '\n'], input);
    }
    
    if !input.contains(CUSTOM_DELIMETER_END) {
        panic!("Invalid delimeter statement");
    }

    let lines:Vec<&str> = input.splitn(2, CUSTOM_DELIMETER_END).collect();
    if lines[1].contains(CUSTOM_DELIMETER_END) {
        panic!("Specified delimeter not used");
    }

    let delimeter = lines[0].replace(CUSTOM_DELIMETER_BEGIN, "").replace(CUSTOM_DELIMETER_END, "");
    if delimeter.len() == 0 {
        panic!("No delimeter specified");
    }
    (delimeter.chars().collect(), lines[1])
}

fn safe_parse(number:&str) -> u32 {
    match number.parse::<u32>() {
        Ok(v) => v,
        Err(_) => 0
    }
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

    #[test]
    #[should_panic(expected = "No delimeter specified")]
    fn panics_when_delimeter_specification_has_no_delimeter() {
        add("//\n3,2,1");
    }

    #[test]
    #[should_panic(expected = "Specified delimeter not used")]
    fn panics_when_delimeter_specification_is_not_used() {
        add("//3\n2\n1");
    }

    #[test]
    #[should_panic(expected = "Invalid delimeter statement")]
    fn panics_when_delimeter_specification_is_invalid() {
        add("//|3|2|1");
    }
}
