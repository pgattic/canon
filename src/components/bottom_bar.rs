use dioxus::prelude::*;

#[derive(PartialEq, Clone, Props)]
pub struct BottomBarProps {
    content: Element,
    bar: Element,
}

#[component]
pub fn BottomBar(props: BottomBarProps) -> Element {
    rsx! {
        div {
            r#style: "
                margin: 0;
                display: grid;
                grid-template-rows: 1fr auto;
                height: 100%;
            ",
            div {
                style: "overflow: scroll;",
                {props.content}
            }
            nav {
                r#style: "
                    width: 100vw;
                    height: 60px;
                    padding: 8px;
                    background-color: #333333;
                    color: white;
                    box-shadow: 0 -2px 5px rgba(0,0,0,0.2);
                ",
                {props.bar}
            }
        }
    }
}

