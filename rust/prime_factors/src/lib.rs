pub fn generate(mut number:u32) -> Vec<u32> {
    let mut primes = vec![];
    while number > 1 {
        for divisor in 2..number+1 {
            if number % divisor == 0 {
                primes.push(divisor);
                number /= divisor;
                break;
            }
        }
    }
    primes
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn zero_returns_empty_list() {
        let empty: Vec<u32> = vec![];
        assert_eq!(empty, generate(0));
    }

    #[test]
    fn one_returns_empty_list() {
        let empty: Vec<u32> = vec![];
        assert_eq!(empty, generate(1));
    }

    #[test]
    fn two_returns_2() {
        assert_eq!(vec![2], generate(2));
    }

    #[test]
    fn three_returns_3() {
        assert_eq!(vec![3], generate(3));
    }

    #[test]
    fn four_returns_2_2() {
        assert_eq!(vec![2, 2], generate(4));
    }

    #[test]
    fn five_returns_5() {
        assert_eq!(vec![5], generate(5));
    }

    #[test]
    fn eight_returns_2_2_2() {
        assert_eq!(vec![2, 2, 2], generate(8));
    }

    #[test]
    fn nine_ninety_eight_returns_2_499() {
        assert_eq!(vec![2, 499], generate(998));
    }

    #[test]
    fn nine_ninety_nine_returns_3_3_3_37() {
        assert_eq!(vec![3, 3, 3, 37], generate(999));
    }
}
