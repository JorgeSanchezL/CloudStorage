const NavBar = (props: { text: string }) => {
    return (
        <div style={{ display: "flex", justifyContent: "center", alignItems: "center", flexDirection: "column", position: "fixed", top: 0, backgroundColor: "#141414", height: "5vh", width: "100vw" }}>
            {props.text}
        </div>
    )
}

export default NavBar
