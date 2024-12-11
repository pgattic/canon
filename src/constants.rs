use std::path::PathBuf;
use dirs::home_dir;

pub fn canon_home() -> PathBuf {
    #[cfg(target_os = "android")]
    return PathBuf::from("/data/data/com.pgattic.Canon/texts");
    #[cfg(not(target_os = "android"))]
    return home_dir().unwrap().join(".canon").join("texts");
}

