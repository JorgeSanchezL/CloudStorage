import Card from './Card'

const File = (props: { text: string }) => {
    return (
        <Card text={props.text} color='#141414' onClick={() => {/*select file*/ }} />
    )
}

export default File