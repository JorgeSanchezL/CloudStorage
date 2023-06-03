import Card from './Card'

const Folder = (props: { text: string }) => {
    return (
        <Card text={props.text} color='#f5c242' onClick={() => {/*navigate to the folder*/ }} />
    )
}

export default Folder