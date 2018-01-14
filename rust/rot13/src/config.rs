pub struct Config {
    pub is_production: bool
}

pub fn debug() -> Config {
	Config {
		is_production: false
	}
}

pub fn production() -> Config {
	Config {
		is_production: true
	}
}

pub fn new(is_production: bool) -> Config {
	Config {
		is_production
	}
}
