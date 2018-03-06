use std::fs::File;
use std::io::prelude::*;

pub type FileContents = Iterator<Item=String>;

pub trait FileIO {
    fn read(&self, filename: &str) -> Box<FileContents>;
    fn write(&self, filename: &str, contents: Box<FileContents>);
}

pub struct StandardFileIO;

impl StandardFileIO {
    pub fn new() -> StandardFileIO {
        StandardFileIO {}
    }
}

impl FileIO for StandardFileIO {
    fn read(&self, filename: &str) -> Box<FileContents> {
        let mut file = File::open(filename).expect("file not found");
        let mut contents = String::default();
        file.read_to_string(&mut contents).expect("could not read file");
        Box::new(FileContentIterator::new(filename))
    }

    fn write(&self, filename: &str, contents: Box<FileContents>) {
        let mut file = File::create(filename).expect("could not create file");
        for data in contents {
            file.write_all(data.as_bytes()).expect("failed to write file contents");
        }
    }
}

pub struct FileContentIterator {
    file: File,
}

impl FileContentIterator {
    fn new(filename: &str) -> FileContentIterator {
        FileContentIterator {
            file: File::open(filename).expect("file not found"),
        }
    }
}

impl Iterator for FileContentIterator {
    type Item = String;

    fn next(&mut self) -> Option<Self::Item> {
        let mut buffer = [0; 256];
        let bytes_read = self.file.read(&mut buffer).expect("error reading file");
        if 0 < bytes_read {
            Some(String::from_utf8(buffer[0..bytes_read].to_vec()).expect("failed to parse bytes"))
        } else {
            None
        }
    }
}

pub struct MockFileContents {
    contents: String
}

impl MockFileContents {
    pub fn new(contents: &str) -> MockFileContents {
        MockFileContents {
            contents: String::from(contents)
        }
    }
}

impl Iterator for MockFileContents {
    type Item = String;

    fn next(&mut self) -> Option<Self::Item> {
        if 0 < self.contents.len() {
            let remainder = self.contents.split_off(2);
            let prefix = self.contents.clone();
            self.contents = remainder;
            Some(prefix)
        } else {
            None
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn read_returns_contents_of_file() {
        let fileio = StandardFileIO::new();
        let contents: String = fileio.read("in.txt").collect();
        assert_eq!(String::from("The dog barks at midnight."), contents);
    }

    #[test]
    fn write_writes_contents_to_file() {
        let filename = "out.txt";
        ::std::fs::remove_file(&filename).ok();

        let fileio = StandardFileIO::new();
        let contents = "The dog barks at midnight.";
        let contents_iter = MockFileContents::new(contents);
        fileio.write(filename, Box::new(contents_iter));
        assert_eq!(String::from(contents), fileio.read(filename).collect::<String>());

        ::std::fs::remove_file(filename).ok();
    }
}
