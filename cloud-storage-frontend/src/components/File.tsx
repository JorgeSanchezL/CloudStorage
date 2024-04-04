import Card from './Card'

const File = (props: { text: string, onClick: () => void  }) => {
    return (
        <Card text={props.text} color='#141414' onClick={props.onClick} />
    )
}

export default File