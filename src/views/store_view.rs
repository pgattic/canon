use libcanon::*;
use dioxus::prelude::*;
use crate::constants::canon_home;

#[component]
pub fn StoreView() -> Element {

    let installed_packages = use_signal(|| pkg_mgr::list(&canon_home()).unwrap());
    let catalogue = pkg_mgr::get_catalogue();
    let mut install_state = use_signal(|| "".to_string());
    rsx! {
        div {
            max_width: "800px",
            h2 { "Installed packages:" }
            for pkg in installed_packages() {
                span {
                    p { "{pkg}" }
                    button {
                        onclick: move |_| {let _ = pkg_mgr::remove(&pkg, &canon_home());},
                        "Delete"
                    }
                }
            }
            h2 { "Catalogue:" }
            for pkg in catalogue {
                if !installed_packages().contains(&pkg.0) {
                    span {
                        p {"{pkg.0}"}
                        button {
                            onclick: move |_| {
                                //let can_p = canon_path.clone();
                                install_state.set(format!("Installing {}...", pkg.0));
                                match pkg_mgr::install(&pkg.1, &canon_home()) {
                                    Err(message) => {install_state.set(message.to_string())},
                                    _ => {install_state.set(format!("{} installed successfully!", pkg.0))}
                                }
                                //installed_packages.set(pkg_mgr::list(&can_p).unwrap())
                            },
                            "Download",
                        }
                    }
                }
            }
            p { "{install_state}" }
        }
    }
}

