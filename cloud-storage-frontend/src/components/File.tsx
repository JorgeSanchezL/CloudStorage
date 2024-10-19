import React, { useState } from 'react'
import './File.css'

interface FileProps {
    text: string
    onDownloadClick: () => void
}

const File: React.FC<FileProps> = ({ text, onDownloadClick }) => {
    const [isHovered, setIsHovered] = useState(false);

    return (
        <div 
            className="file-container"
            onMouseEnter={() => setIsHovered(true)}
            onMouseLeave={() => setIsHovered(false)}
        >
            <div className="file-icon">📄</div>
            <div className="file-text">{text}</div>
            {isHovered && (
                <button className="download-button" onClick={(e) => {
                    e.stopPropagation();
                    onDownloadClick();
                }}>
                    ⬇️
                </button>
            )}
        </div>
    )
}

export default File