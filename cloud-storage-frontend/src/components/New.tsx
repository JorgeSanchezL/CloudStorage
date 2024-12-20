import React from 'react'
import './New.css'

interface NewProps {
    onClick: () => void
}

const New: React.FC<NewProps> = ({ onClick }) => {
    return (
        <div onClick={onClick} className="new-container">
            <div className="new-icon">➕</div>
        </div>
    )
}

export default New