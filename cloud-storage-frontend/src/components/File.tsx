import React from 'react'
import './File.css'

interface FileProps {
    text: string
    onClick: () => void
}

const File: React.FC<FileProps> = ({ text, onClick }) => {
    return (
        <div onClick={onClick} className="file-container">
            <div className="file-icon">📄</div>
            <div className="file-text">{text}</div>
        </div>
    )
}

export default File