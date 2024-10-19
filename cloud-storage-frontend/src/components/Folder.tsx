import React from 'react'
import './Folder.css'

interface FolderProps {
    name: string
    onClick: () => void
}

const Folder: React.FC<FolderProps> = ({ name, onClick }) => {
    return (
        <div onClick={onClick} className="folder-container">
            <div className="folder-icon">📁</div>
            <div className="folder-text">{name}</div>
        </div>
    )
}

export default Folder