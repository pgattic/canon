use dioxus::prelude::*;
use crate::views::ScriptureView;
use crate::components::TopBar;
use libcanon::reference::Reference;
use libcanon::citation;
use crate::canon_home;

#[component]
pub fn Reading() -> Element {
    let mut show_numbers = use_signal(|| true);

    // Parse the reference
    let reference = Reference::from_str(&"1ne3").unwrap();
    //println!("Reference: {:?}", reference);
    let mut result = use_signal(|| citation::cite(&canon_home(), &reference).unwrap());
    //let result = use_resource(move || async move { citation::cite(&canon_home(), &reference).await });
    //println!("Result: {:?}", result);
    rsx! {
        TopBar {
            bar: rsx! {
                input {
                    //border: "none",
                    //background: "none",
                    padding: "4px",
                    type: "text",
                    value: "1ne3",
                    size: 12,
                    oninput: move |e| {
                        let reference = Reference::from_str(&e.value()).unwrap();
                        result.set(citation::cite(&canon_home(), &reference).unwrap());
                    },
                }
                button {
                    onclick: move |_| {show_numbers.toggle();},
                    "Show/hide numbers"
                }
            },
            content: rsx! {
                ScriptureView { content: result(), show_numbers: show_numbers() }
            }
        }
    }
}

