
/* Project Modules */
mod components;
mod pages;
mod views;
mod constants;


use dioxus::prelude::*;
use dioxus_logger::tracing::Level;

use pages::*;
use constants::canon_home;

#[derive(Clone, Routable, Debug, PartialEq)]
pub enum Route {
    #[layout(Navbar)]
    #[route("/")]
    Reading {},
    #[route("/search")]
    Search {},
    #[route("/store")]
    Store {},
}

const MAIN_CSS: Asset = asset!("/assets/main.css");

fn main() {
    if !canon_home().exists() {
        let _ = std::fs::create_dir_all(&canon_home());
        println!("Canon home dir created at {:?}", canon_home());
    }
    dioxus_logger::init(Level::INFO).expect("failed to init logger");
    dioxus::launch(App);
}

#[component]
fn App() -> Element {
    rsx! {
        document::Link { rel: "stylesheet", href: MAIN_CSS }
        Router::<Route> {}
    }
}

#[component]
pub fn Navbar() -> Element {
    rsx! {
        div {
            style: "height: 100vh",
            div {
                r#style: "
                    margin: 0;
                    display: grid;
                    grid-template-rows: 1fr auto;
                    height: 100%;
                ",
                div {
                    style: "overflow: scroll;",
                    Outlet::<Route> {}
                }
                nav {
                    r#style: "
                        width: 100%;
                        padding: 8px;
                        background-color: #333333;
                        color: white;
                        box-shadow: 0 -2px 5px rgba(0,0,0,0.2);
                        display: flex;
                        justify-content: space-around;
                    ",
                    Link { to: Route::Reading {}, "Home" }
                    Link { to: Route::Search {}, "Search" }
                    Link { to: Route::Store {}, "Store" }
                }
            }
        }
        //div {
        //    r#style: "height: 100vh;",
        //    BottomBar {
        //        content: rsx! {
        //            Router::<Route> {}
        //        },
        //        bar: rsx! {
        //            p {"hello"}
        //            Link { to: Route::Reading {}, "Home" }
        //            Link { to: Route::Search {}, "Search" }
        //        }
        //    }
        //}
    }
}

