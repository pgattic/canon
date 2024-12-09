use dioxus::prelude::*;
use crate::views::SearchResultsView;
use crate::components::TopBar;

#[component]
pub fn Search() -> Element {
    let mut query = use_signal(|| String::new());
    let mut input = use_signal(|| String::new());

    rsx! {
        TopBar {
            bar: rsx! {
                input {
                    r#type: "text",
                    value: "{query}",
                    oninput: move |e| {input.set(e.value());},
                }
                button {
                    onclick: move |_| {query.set(input.to_string());},
                    "Search!"
                }
            },
            content: rsx! {
                SearchResultsView { query: query }
            }
        }
    }
}

