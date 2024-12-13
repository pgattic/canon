use dioxus::prelude::*;
use crate::views::StoreView;
use crate::components::TopBar;

#[component]
pub fn Store() -> Element {

    rsx! {
        TopBar {
            bar: rsx! {
                "Package Management"
            },
            content: rsx! {
                StoreView {}
            }
        }
    }
}

