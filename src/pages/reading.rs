use dioxus::prelude::*;
use crate::views::ScriptureView;
use crate::components::TopBar;

#[component]
pub fn Reading() -> Element {
    let mut query = use_signal(|| String::from("1ne3"));
    let mut show_numbers = use_signal(|| true);

    rsx! {
        TopBar {
            bar: rsx! {
                input {
                    r#type: "text",
                    value: "{query}",
                    oninput: move |e| {query.set(e.value());},
                }
                button {
                    onclick: move |_| {show_numbers.toggle();},
                    "Show/hide numbers"
                }
            },
            content: rsx! {
                ScriptureView { query: query, show_numbers: show_numbers() }
            }
        }
    }
}

