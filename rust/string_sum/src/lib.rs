const CUSTOM_DELIMETER_BEGIN:&str = "//";
const CUSTOM_DELIMETER_END:&str = "\n";

pub fn add(numbers:&str) -> Result<u32, String> {
    let input = decide_delimeters(numbers);

    let numbers:Vec<&str> = input.numbers.split(|c:char| input.delimeters.contains(&c)).collect();
    let mut negatives = String::from("");
    let mut total = 0;
    for number in numbers {
        if number.contains('-') {
            negatives.push_str(&format!(" {}", &number));
        } else {
            total += safe_parse(&number);
        }
    }
    if negatives.len() > 0 {
        return Err(String::from(format!("negatives not allowed:{}", &negatives)));
    }

    Ok(total)
}

struct ParsedInput<'a> {
    delimeters:Vec<char>,
    numbers:&'a str
}

impl<'a> ParsedInput<'a> {
    fn default(numbers:&'a str) -> ParsedInput {
        ParsedInput {
            delimeters: vec![',', '\n'],
            numbers
        }
    }

    fn new(numbers:&'a str, delimeters:Vec<char>) -> ParsedInput {
        if numbers.contains(CUSTOM_DELIMETER_END) {
            panic!("Specified delimeter not used");
        }

        if delimeters.len() == 0 {
            panic!("No delimeter specified");
        }

        ParsedInput {
            delimeters,
            numbers
        }
    }
}

fn decide_delimeters(input:&str) -> ParsedInput {
    if !input.starts_with(CUSTOM_DELIMETER_BEGIN) {
        return ParsedInput::default(input);
    }
    
    if !input.contains(CUSTOM_DELIMETER_END) {
        panic!("Invalid delimeter statement");
    }

    let lines:Vec<&str> = input.splitn(2, CUSTOM_DELIMETER_END).collect();
    let delimeter = lines[0].replace(CUSTOM_DELIMETER_BEGIN, "").replace(CUSTOM_DELIMETER_END, "");
    ParsedInput::new(lines[1], delimeter.chars().collect())
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
        assert_eq!(0, add("").unwrap());
    }

    #[test]
    fn returns_1_for_string_with_1() {
        assert_eq!(1, add("1").unwrap());
    }

    #[test]
    fn returns_3_for_string_with_1_and_2() {
        assert_eq!(3, add("1,2").unwrap());
    }

    #[test]
    fn supports_newline_delimeter() {
        assert_eq!(6, add("2\n3,1").unwrap());
    }

    #[test]
    fn supports_delimeter_specification() {
        assert_eq!(12, add("//;\n4;6;2").unwrap());
    }

    #[test]
    #[should_panic(expected = "No delimeter specified")]
    fn panics_when_delimeter_specification_has_no_delimeter() {
        let _ = add("//\n3,2,1");
    }

    #[test]
    #[should_panic(expected = "Specified delimeter not used")]
    fn panics_when_delimeter_specification_is_not_used() {
        let _ = add("//3\n2\n1");
    }

    #[test]
    #[should_panic(expected = "Invalid delimeter statement")]
    fn panics_when_delimeter_specification_is_invalid() {
        let _ = add("//|3|2|1");
    }

    #[test]
    fn returns_error_when_negative_number_is_in_input() {
        assert_eq!(Err(String::from("negatives not allowed: -1")), add("1,4,-1"));
    }

    #[test]
    fn returns_error_with_all_negative_numbers_when_negative_numbers_are_in_input() {
        assert_eq!(Err(String::from("negatives not allowed: -4 -1")), add("1,-4,-1"));
    }
}
