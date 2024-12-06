#![allow(non_snake_case)]

mod components;
mod pages;
mod views;

use pages::*;
use dioxus::desktop::WindowBuilder;
use dioxus::prelude::*;
use dioxus::desktop::tao::dpi::PhysicalPosition;

//use crate::components::BottomBar;

fn main() {
    let cfg = dioxus::desktop::Config::new()
        .with_custom_head(r#"
<style>
  body {
    margin: 0;
    overflow-x: hidden;
    /*padding: 0;*/
  }

  @media (prefers-color-scheme: dark) {
    body {
      background-color: #222222;
      color: #dddddd;
    }
  }

  * {
    -webkit-user-select: none;
    -ms-user-select: none;
    user-select: none;
  }
</style>"#.to_string())
        .with_window(
            WindowBuilder::new()
                .with_title("Canon")
                //.with_decorations(false)
                .with_position(PhysicalPosition::new(0, 0)));

    LaunchBuilder::desktop().with_cfg(cfg).launch(App);
}

#[component]
fn App() -> Element {
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
                    p {"hello"}
                    Link { to: Route::Reading {}, "Home" }
                    Link { to: Route::Search {}, "Search" }
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

#[derive(Clone, Routable, Debug, PartialEq)]
pub enum Route {
    #[route("/")]
    Reading {},
    #[route("/search")]
    Search {},
    #[route("/store")]
    Store {},
}

//#[component]
//pub fn BottomNavBar() -> Element {
//    rsx! {
//    }
//}
//
