#![allow(non_snake_case)]

use dirs::home_dir;
use std::path::PathBuf;
use dioxus::prelude::*;
use dioxus_logger::tracing::{info, Level};
use libcanon::reference::Reference;
use libcanon::*;

#[derive(Clone, Routable, Debug, PartialEq)]
enum Route {
    #[route("/")]
    Home {},
    #[route("/blog/:id")]
    Blog { id: i32 },
}

fn main() {
    // Init logger
    dioxus_logger::init(Level::INFO).expect("failed to init logger");
    info!("starting app");

    dioxus::launch(App);
}

#[component]
fn App() -> Element {
    rsx! {
        Router::<Route> {}
    }
}

#[component]
fn Blog(id: i32) -> Element {
    rsx! {
        Link { to: Route::Home {}, "Go to counter" }
        "Blog post {id}"
    }
}

#[component]
fn ScriptureView(query: String) -> Element {
    let canon_path: PathBuf = home_dir().unwrap().join(".canon");

    // Parse the reference
    let reference = Reference::from_str(&query).unwrap();
    let result = citation::cite(&canon_path, &reference);
    match result {
        Ok(citation) => {
            rsx! {
                for ch in citation.chapters {
                    for v in ch.verses {
                        p { "{v.verse} {v.content}" }
                    }
                }
            }
        }
        Err(problem) => {
            rsx! {
                p { "Gotta problem? Hello World" }
            }
        }
    }
}

#[component]
fn Home() -> Element {
    let mut count = use_signal(|| 0);

    rsx! {
        Link {
            to: Route::Blog {
                id: count()
            },
            "Go to blog"
        }
        div {
            h1 { "High-Five counter: {count}" }
            button { onclick: move |_| count += 1, "Up high!" }
            button { onclick: move |_| count -= 1, "Down low!" }
            ScriptureView { query: "1ne3" }
        }
    }
}
