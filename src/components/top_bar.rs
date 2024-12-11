use dioxus::prelude::*;

#[derive(PartialEq, Clone, Props)]
pub struct TopBarProps {
    bar: Element,
    content: Element
}

#[component]
pub fn TopBar(props: TopBarProps) -> Element {
    rsx! {
        div {
            r#style: "
                margin: 0;
                display: grid;
                grid-template-rows: auto 1fr;
                min-height: 0;
                width: 100vw;
                height: 100%;
            ",
            nav {
                r#style: "
                    padding: 8px;
                    background: #333333;
                    box-shadow: 0 2px 5px rgba(0,0,0,0.2);
                ",
                {props.bar}
            }
            div {
                r#style: "
                    margin: 0 auto;
                    max-width: 800px;
                    padding: 0 32px 16px;
                    overflow: scroll;
                ",
                {props.content}
            }
        }
    }
}

