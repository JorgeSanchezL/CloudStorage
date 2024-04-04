import Card from './Card'

const Folder = (props: { name: string, onClick: () => void }) => {
    return (
        <Card text={props.name} color='#f5c242' onClick={props.onClick}/>
    )
}

export default Folder