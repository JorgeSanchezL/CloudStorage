import './App.css'
import NavBar from '../components/NavBar'
import File from '../components/File'
import Folder from '../components/Folder'

const App = () => {
    return (
        <div>
            <NavBar text='CloudStorage' />
            <div style={{ display: "grid", gridTemplateColumns: "repeat(7, 1fr)", rowGap: "32px", justifyItems: "center", marginTop: "calc(5vh + 32px)" }}>
                <File text='File1' />
                <File text='File2' />
                <File text='File3' />
                <File text='File4' />
                <File text='File5' />
                <File text='File6' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <File text='File7' />
                <Folder text='File7' />
                <Folder text='File7' />
                <Folder text='File7' />
                <Folder text='File7' />
                <Folder text='File7' />
            </div>
        </div>
    )
}

export default App
