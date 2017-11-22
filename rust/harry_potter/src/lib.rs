use std::collections::HashMap;
use std::collections::HashSet;

pub const FIRST: &str = "First";
pub const SECOND: &str = "Second";
pub const THIRD: &str = "Third";
pub const FOURTH: &str = "Fourth";
pub const FIFTH: &str = "Fifth";

const BASE_PRICE: f64 = 8.0;

pub fn calc_cost(mut books: HashMap<&str, u32>) -> f64 {
    let mut sets: Vec<HashSet<String>> = vec![];
    loop {
        let mut set = HashSet::new();
        for (title, count) in &mut books {
            if count > &mut 0 {
                set.insert(title.to_string());
                *count -= 1;
            }
        }
        let set_size = set.len();
        sets.push(set);
        if 0 == set_size {
            break;
        }
    }
    cost_of(ensure_best_discount(sets))
}

/*fn ensure_best_discount(sets: &mut Vec<HashSet<String>>) {
    loop {
        let mut fiver = match sets.iter().filter(|s| 5 == s.len()).next() {
            Some(s) => s,
            None => break
        };
        
        for threer in sets.iter().filter(|s| 3 == s.len()) {
            let title = fiver.difference(threer).next().unwrap();
            fiver.remove(title);
            threer.insert(title);
            break;
        }
    }
}*/

/*fn ensure_best_discount(sets: Vec<HashSet<String>>) -> Vec<HashSet<String>> {
    let five_sets = sets.iter().filter(|s| 5 == s.len());
    let three_sets = sets.iter().filter(|s| 3 == s.len());
    let other_sets = sets.iter().filter(|s| 5 != s.len() && 3 != s.len());
    let (five_sets, three_sets): (Vec<HashSet<String>>, Vec<HashSet<String>>) = five_sets.zip(three_sets).map(|(fiver, threer)| {
        let title = fiver.difference(threer).next().unwrap().to_string();
        let mut fiver = fiver.clone();
        fiver.remove(&title);
        let mut threer = threer.clone();
        threer.insert(title);
        (fiver, threer)
    }).unzip();
    other_sets.chain(five_sets.iter()).chain(three_sets.iter()).collect()
}*/

fn ensure_best_discount(sets: Vec<HashSet<String>>) -> Vec<HashSet<String>> {
    let mut result = vec![];
    let five_sets = sets.iter().filter(|s| 5 == s.len());
    let three_sets = sets.iter().filter(|s| 3 == s.len());
    five_sets.zip(three_sets).map(|(fiver, threer)| {
        let title = fiver.difference(threer).next().unwrap().to_string();
        let mut fiver = fiver.clone();
        fiver.remove(&title);
        let mut threer = threer.clone();
        threer.insert(title);
        (fiver, threer)
    }).for_each(|(fiver, threer)| {
        result.push(fiver);
        result.push(threer);
    });

    let other_sets = sets.iter().filter(|s| 5 != s.len() && 3 != s.len());
    other_sets.for_each(|s| result.push(s.clone()));
    println!("final sets: {:?}", result);
    result
}

fn cost_of(sets: Vec<HashSet<String>>) -> f64 {
    sets.iter().map(cost_of_set).sum()
}

fn cost_of_set(set: &HashSet<String>) -> f64 {
    (set.len() as f64) * BASE_PRICE * (1.0 - get_discount(set.len()))
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

    #[test]
    fn another_complex_example() {
        let books = build_basket(vec![FIRST, SECOND, THIRD, FOURTH, FIFTH, FIRST, SECOND]);
        assert_eq!(45.2, calc_cost(books));
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
