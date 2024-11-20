use dioxus::prelude::*;
use crate::views::StoreView;
use crate::components::TopBar;
use crate::Route;

#[component]
pub fn Store() -> Element {

    rsx! {
        TopBar {
            bar: rsx! {
                p {"Download somethin'!"}
            },
            content: rsx! {
                Link { to: Route::Reading {}, "Read" }
                Link { to: Route::Search {}, "Search" }
                StoreView {}
            }
        }
    }
}

