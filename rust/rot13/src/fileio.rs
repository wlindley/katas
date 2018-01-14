use std::fs::File;
use std::io::prelude::*;

pub trait FileIO {
    fn read(&self, filename: &str) -> String;
    fn write(&self, filename: &str, contents: String);
}

pub struct StandardFileIO;

impl StandardFileIO {
    pub fn new() -> StandardFileIO {
        StandardFileIO {}
    }
}

impl FileIO for StandardFileIO {
    fn read(&self, filename: &str) -> String {
        let mut file = File::open(filename).expect("file not found");
        let mut contents = String::default();
        file.read_to_string(&mut contents).expect("could not read file");
        contents
    }

    fn write(&self, filename: &str, contents: String) {
        let mut file = File::create(filename).expect("could not create file");
        file.write_all(contents.as_bytes()).expect("failed to write file contents");
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn read_returns_contents_of_file() {
        let fileio = StandardFileIO::new();
        let contents = fileio.read("in.txt");
        assert_eq!(String::from("The dog barks at midnight."), contents);
    }

    #[test]
    fn write_writes_contents_to_file() {
        let filename = "out.txt";
        ::std::fs::remove_file(&filename).ok();

        let fileio = StandardFileIO::new();
        let contents = String::from("The dog barks at midnight.");
        fileio.write(filename, contents.clone());
        assert_eq!(contents, fileio.read(filename));

        ::std::fs::remove_file(filename).ok();
    }
}
