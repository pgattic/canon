use dioxus::prelude::*;
use dioxus_logger::tracing::Level;

mod components;
mod pages;
mod views;

use pages::*;

#[derive(Clone, Routable, Debug, PartialEq)]
pub enum Route {
    #[route("/")]
    Reading {},
    #[route("/search")]
    Search {},
    #[route("/store")]
    Store {},
}

const MAIN_CSS: Asset = asset!("/assets/main.css");

fn main() {
    dioxus_logger::init(Level::INFO).expect("failed to init logger");
    dioxus::launch(App);
}

#[component]
fn App() -> Element {
    rsx! {
        document::Link { rel: "stylesheet", href: MAIN_CSS }
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
                    Router::<Route> {}
                }
                nav {
                    r#style: "
                        width: 100%;
                        height: 60px;
                        padding: 8px;
                        background-color: #282c34;
                        color: white;
                        box-shadow: 0 -2px 5px rgba(0,0,0,0.2);
                    ",
                    //Link { to: Route::Reading {}, "Home" }
                    //Link { to: Route::Search {}, "Search" }
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

//#[component]
//pub fn BottomNavBar() -> Element {
//    rsx! {
//    }
//}
//
