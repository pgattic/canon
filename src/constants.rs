use std::path::PathBuf;
use dirs::home_dir;

pub fn canon_home() -> PathBuf {
    //PathBuf::from("/data/data/com.pgattic.Canon/texts")
    home_dir().unwrap().join(".canon").join("texts")
}

