const Card = (props: { text: string, color: string, onClick: () => void }) => {
    return (
        <div
            style={{
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
                flexDirection: "column",
                backgroundColor: props.color,
                height: "10vh", width: "10vw",
                borderRadius: "16px"
            }}
            onClick={props.onClick}
        >
            {props.text}
        </div>
    )
}

export default Card
