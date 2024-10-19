import React, { useEffect, useState } from 'react'

import { useNavigate, useLocation } from 'react-router-dom'

import File from '../components/File'
import Folder from '../components/Folder'
import { getDirectoryInfo, getFile } from './backend'

type DirectoryInfo = {
    Files: FileInfo[]
    Directories: string[]
}

type FileInfo = {
    Name: string;
    Size: number;
    LastModification: string;
};

const FileSystem = () => {
    const { pathname } = useLocation()

    let path = pathname.replace(/^\/+/, '')
    let fileSystemPath = path.includes('/') ? path.substring(path.indexOf('/') + 1) : ""

    const navigate = useNavigate()

    const [directoryInfo, setDirectoryInfo] = useState<DirectoryInfo | null>(null)


    useEffect(() => {
        getDirectoryInfo(fileSystemPath).then((res) => {
            setDirectoryInfo(res)
        })
    }, [pathname])

    return (
        <div style={{ display: "grid", gridTemplateColumns: "repeat(7, 1fr)", rowGap: "32px", justifyItems: "center", marginTop: "calc(5vh + 32px)" }}>
            {
                fileSystemPath ?
                    <Folder key={".."} name={"↩"} onClick={
                        () => {
                            let newPath = path.endsWith('/') ? path.slice(0, -1) : path;
                            newPath = newPath.substring(0, newPath.lastIndexOf('/'));
                            navigate(`/${newPath}`)
                        }
                    } />
                    : 
                    null
            }
            {
                directoryInfo?.Directories.map((directory, index) => {
                    return <Folder key={index} name={directory} onClick={
                        () => {
                            let currentPath = fileSystemPath == "" ? "" : fileSystemPath + '/';
                            navigate(currentPath + directory)
                        }
                    } />
                })
            }
            {
                directoryInfo?.Files.map((fileInfo, index) => {
                    return <File key={index} fileInfo={fileInfo} onDownloadClick={
                        () => {
                            getFile(fileSystemPath + "/" + fileInfo.Name)
                        }
                    } />
                })
            }
        </div>
    )
}

export default FileSystem
