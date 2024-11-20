use dirs::home_dir;
use std::path::PathBuf;
use libcanon::*;
use dioxus::prelude::*;

#[component]
pub fn StoreView() -> Element {
    let canon_path: PathBuf = home_dir().unwrap().join(".canon").join("texts");

    let installed_packages = pkg_mgr::list(&canon_path);
    match installed_packages {
        Ok(pkgs) => {
            rsx! {
                for pkg in pkgs {
                    p {"{pkg}"}
                }
            }
        }
        Err(problem) => {
            rsx! {
                p { "Error: {problem}" }
            }
        }
    }
}

