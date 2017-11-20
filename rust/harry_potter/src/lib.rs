use std::collections::HashMap;

const BASE_PRICE: f64 = 8.0;
pub const FIRST: &str = "First";
pub const SECOND: &str = "Second";
pub const THIRD: &str = "Third";
pub const FOURTH: &str = "Fourth";
pub const FIFTH: &str = "Fifth";

pub fn calc_cost(mut books: HashMap<&str, u32>) -> f64 {
    let mut cost: f64 = 0.0;
    loop {
        let mut set = vec![];
        for (title, count) in &mut books {
            if count > &mut 0 {
                set.push(title);
                *count -= 1;
            }
        }
        println!("Found set: {:?}", set);
        if 0 != set.len() {
            cost += (set.len() as f64) * BASE_PRICE * (1.0 - get_discount(set.len()));
        } else {
            break;
        }
    }
    cost
}

fn get_discount(num_distinct: usize) -> f64 {
    if 5 == num_distinct {
        0.25
    } else if 4 == num_distinct {
        0.20
    } else if 3 == num_distinct {
        0.10
    } else if 2 == num_distinct {
        0.05
    } else {
        0.0
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn one_book_costs_8() {
        let books = build_basket(vec![FIRST]);
        assert_eq!(8.0, calc_cost(books));
    }

    #[test]
    fn two_identical_books_costs_16() {
        let books = build_basket(vec![FIRST, FIRST]);
        assert_eq!(16.0, calc_cost(books));
    }

    #[test]
    fn two_different_books_get_5_percent_discount() {
        let books = build_basket(vec![FIRST, SECOND]);
        assert_eq!(16.0 * 0.95, calc_cost(books));
    }

    #[test]
    fn three_different_books_get_10_percent_discount() {
        let books = build_basket(vec![FIRST, SECOND, THIRD]);
        assert_eq!(24.0 * 0.9, calc_cost(books));
    }

    #[test]
    fn four_different_books_get_20_percent_discount() {
        let books = build_basket(vec![FIRST, SECOND, THIRD, FOURTH]);
        assert_eq!(32.0 * 0.8, calc_cost(books));
    }

    #[test]
    fn three_different_books_get_25_percent_discount() {
        let books = build_basket(vec![FIRST, SECOND, THIRD, FOURTH, FIFTH]);
        assert_eq!(40.0 * 0.75, calc_cost(books));
    }

    #[test]
    fn title_outside_of_set_does_not_get_discount() {
        let books = build_basket(vec![FIRST, SECOND, THIRD, FIRST]);
        assert_eq!(8.0 + (24.0 * 0.9), calc_cost(books));
    }

    #[test]
    fn complex_example() {
        let books = build_basket(vec![FIRST, FIRST, SECOND, SECOND, THIRD, THIRD, FOURTH, FIFTH]);
        assert_eq!(51.2, calc_cost(books));
    }

    fn build_basket(books: Vec<&str>) -> HashMap<&str, u32> {
        let mut basket = HashMap::new();
        for book in books {
            let count = basket.entry(book).or_insert(0);
            *count += 1;
        }
        basket
    }
}
