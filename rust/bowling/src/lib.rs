const MAX_PINS:u32 = 10;
const NUM_FRAMES:u32 = 10;

pub struct Throw {
    pins:u32
}

impl Throw {
    pub fn new(pins:u32) -> Throw {
        if pins > MAX_PINS {
            panic!("Number of pins must be {} or less", MAX_PINS);
        }
        Throw {
            pins
        }
    }
}

pub struct Game {
    throws:Vec<Throw>
}

impl Game {
    pub fn new() -> Game {
        Game {
            throws: vec![]
        }
    }
    
    pub fn throw(&mut self, throw:Throw) {
        self.throws.push(throw);
    }

    pub fn score(&self) -> u32 {
        let mut score = 0;
        let mut frame_index = 0;
        for _ in 0..NUM_FRAMES {
            let first_throw = &self.throws[frame_index];
            let second_throw = &self.throws[frame_index + 1];
            let mut throws_consumed = 2;

            score += first_throw.pins + second_throw.pins;

            if Game::is_strike(&first_throw) {
                score += self.throws[frame_index + 2].pins;
                throws_consumed = 1;
            } else if Game::is_spare(&first_throw, &second_throw) {
                score += self.throws[frame_index + 2].pins;
            }

            frame_index += throws_consumed;
        }
        score
    }

    fn is_strike(first_throw:&Throw) -> bool {
        MAX_PINS <= first_throw.pins
    }

    fn is_spare(first_throw:&Throw, second_throw:&Throw) -> bool {
        MAX_PINS <= first_throw.pins + second_throw.pins
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn all_zeroes_scores_0() {
        let mut game = Game::new();
        throw(&mut game, 20, 0);
        assert_eq!(0, game.score());
    }

    #[test]
    fn all_ones_scores_20() {
        let mut game = Game::new();
        throw(&mut game, 20, 1);
        assert_eq!(20, game.score());
    }

    #[test]
    #[should_panic(expected = "Number of pins must be 10 or less")]
    fn cannot_throw_more_than_ten_pins() {
        let mut game = Game::new();
        game.throw(Throw::new(11));
    }

    #[test]
    fn spare_in_first_frame_counts_next_throw() {
        let mut game = Game::new();
        game.throw(Throw::new(3));
        game.throw(Throw::new(7));
        game.throw(Throw::new(5));
        throw(&mut game, 17, 0);
        assert_eq!(20, game.score());
    }

    #[test]
    fn multiple_spares_each_count_next_throw() {
        let mut game = Game::new();
        throw(&mut game, 6, 5);
        throw(&mut game, 14, 0);
        assert_eq!(15 + 15 + 10, game.score());
    }

    #[test]
    fn spare_in_last_frame_counts_extra_throw() {
        let mut game = Game::new();
        throw(&mut game, 18, 0);
        game.throw(Throw::new(5));
        game.throw(Throw::new(5));
        game.throw(Throw::new(7));
        assert_eq!(17, game.score());
    }

    #[test]
    fn strike_in_first_frame_counts_next_two_throws() {
        let mut game = Game::new();
        game.throw(Throw::new(10));
        game.throw(Throw::new(2));
        game.throw(Throw::new(6));
        throw(&mut game, 16, 0);
        assert_eq!(26, game.score());
    }

    #[test]
    fn multiple_strikes_each_count_next_two_throws() {
        let mut game = Game::new();
        game.throw(Throw::new(10));
        game.throw(Throw::new(10));
        game.throw(Throw::new(10));
        game.throw(Throw::new(10));
        throw(&mut game, 12, 0);
        assert_eq!(90, game.score());
    }

    #[test]
    fn strike_in_the_last_frame_counts_next_two_throw() {
        let mut game = Game::new();
        throw(&mut game, 18, 0);
        game.throw(Throw::new(10));
        game.throw(Throw::new(10));
        game.throw(Throw::new(10));
        assert_eq!(30, game.score());
    }

    #[test]
    fn complex_game_scores_correctly() {
        let mut game = Game::new();
        game.throw(Throw::new(10));

        game.throw(Throw::new(5));
        game.throw(Throw::new(5));

        game.throw(Throw::new(7));
        game.throw(Throw::new(1));

        throw(&mut game, 10, 0);

        game.throw(Throw::new(10));

        game.throw(Throw::new(10));
        game.throw(Throw::new(6));
        game.throw(Throw::new(1));
        assert_eq!(20 + 17 + 8 + 26 + 17, game.score());
    }

    fn throw(game:&mut Game, num_throws:u32, num_pins:u32) {
        for _ in 1..num_throws+1 {
            game.throw(Throw::new(num_pins));
        }
    }
}
