import './App.css'

import { useState, useEffect } from 'react'
import { Route, Routes, useNavigate, useParams } from 'react-router-dom'

import File from '../components/File'
import Folder from '../components/Folder'
import { getDirectoryInfo} from './backend'

type DirectoryInfo = {
    Files: string[]
    Directories: string[]
}


const FileSystem = (props: {parentPath: string, setParentLoad: (value: boolean) => void}) => {

    let [shouldLoad, setShouldLoad] = useState<boolean>(true)

    const { path } = useParams()
    let realPath = (props.parentPath + "/" + path).replace(/^\/+/, '')

    const navigate = useNavigate()

    const [directoryInfo, setDirectoryInfo] = useState<DirectoryInfo | null>(null)

    useEffect(() => {
        props.setParentLoad(false)
        getDirectoryInfo(realPath).then((res) => {
            setDirectoryInfo(res)
            console.log(res)
        })
    }, [])

    return (
        <div>
            {
                shouldLoad ? 
                    <div style={{ display: "grid", gridTemplateColumns: "repeat(7, 1fr)", rowGap: "32px", justifyItems: "center", marginTop: "calc(5vh + 32px)" }}>
                        {
                            directoryInfo?.Directories.map((directory, index) => {
                                return <Folder key={index} name={directory} onClick={
                                    () => {
                                        navigate(directory)
                                    }
                                }/>
                            })
                        }
                        {
                            directoryInfo?.Files.map((file, index) => {
                                return <File key={index} text={file} />
                            })
                        }
                    </div>
                : null
            }
        <Routes>
            <Route path="/:path/*" element={<FileSystem setParentLoad={(value: boolean) => setShouldLoad(value)} parentPath={realPath}/>} />
        </Routes>
        </div>
    )
}

export default FileSystem
