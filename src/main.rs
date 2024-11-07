#![allow(non_snake_case)]

use dioxus::desktop::WindowBuilder;
use dirs::home_dir;
use std::path::PathBuf;
use dioxus::prelude::*;
//use dioxus_logger::tracing::{info, Level};
use libcanon::reference::Reference;
use libcanon::*;
use dioxus::desktop::tao::dpi::PhysicalPosition;

#[derive(Clone, Routable, Debug, PartialEq)]
enum Route {
    #[route("/")]
    Home {},
    #[route("/blog/:id")]
    Blog { id: i32 },
}

fn main() {
    // Init logger
    //dioxus_logger::init(Level::INFO).expect("failed to init logger");
    //info!("starting app");

    let cfg = dioxus::desktop::Config::new()
        .with_custom_head(r#"<style>
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
        </style>"#.to_string())
        .with_window(
            WindowBuilder::new()
                .with_title("Canon")
                .with_decorations(false)
                .with_position(PhysicalPosition::new(0, 0)));

    LaunchBuilder::desktop().with_cfg(cfg).launch(App);
    //dioxus::launch(App);
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
fn PackagesView() -> Element {
    let canon_path: PathBuf = home_dir().unwrap().join(".canon");
    
    let mut pkgs = use_signal(|| pkg_mgr::list(&canon_path).unwrap());

    rsx! {
        for pkg in pkgs.iter() {
            span {
                p {"{pkg}",}
                //button {
                //    onclick: |_| {
                //        pkg_mgr::remove(&pkg, &canon_path);
                //        pkgs.set(pkg_mgr::list(&canon_path).unwrap());
                //    },
                //    "delete",
                //}
            }
        }
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
                //h1 { "{citation.book_name}" }
                for ch in citation.chapters {
                    if ch.entire_chapter {
                        h2 {
                            style: "text-align: center",
                            "Chapter {ch.path.file_name().unwrap().to_str().unwrap()}"
                        }
                    }
                    for v in ch.verses {
                        p { b {"{v.verse} "} "{v.content}" }
                    }
                }
            }
        }
        Err(_problem) => {
            rsx! {
                p { "Reference not found" }
            }
        }
    }
}

#[component]
fn Home() -> Element {
    let mut query = use_signal(|| String::from("1ne3"));

    rsx! {
        nav {
            r#style: "
                position: sticky;
                top: 0;
                width: 100%;
                padding: 8px;
                background: #444444;
            ",
            input {
                r#style: "
                    background: #f00;
                ",
                r#type: "text",
                value: "{query}",
                oninput: move |e| {
                    query.set(e.value());
                },
            }
        }
        div {
            r#style: "
                margin: 0 auto;
                max-width: 800px;
            ",
            PackagesView{},
            ScriptureView { query: "{query}" }
        }
    }
}
