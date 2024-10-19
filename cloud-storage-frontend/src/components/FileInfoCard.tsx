import React from 'react';

import './FileInfoCard.css';

type FileInfo = {
    Name: string;
    Size: number;
    LastModification: string;
};

interface FileInfoCardProps {
    fileInfo: FileInfo;
    onClose: () => void;
}

const formatSize = (size: number) => {
    if (size < 1024) {
        return `${size} B`;
    } else if (size < 1024 * 1024) {
        return `${(size / 1024).toFixed(2)} KB`;
    } else {
        return `${(size / (1024 * 1024)).toFixed(2)} MB`;
    }
}

const formatDateTime = (date: string) => {
    return new Date(date).toLocaleString();
}

const FileInfoCard: React.FC<FileInfoCardProps> = ({ fileInfo, onClose }) => {
    return (
        <div className="modal-overlay">
            <div className="modal-content">
                <button className="close-button" onClick={onClose}>✖️</button>
                <h2>File Information</h2>
                <p><strong>Name:</strong> {fileInfo.Name}</p>
                <p><strong>Size:</strong> {formatSize(fileInfo.Size)}</p>
                <p><strong>Date of last modification:</strong> {formatDateTime(fileInfo.LastModification)}</p>
            </div>
        </div>
    );
};

export default FileInfoCard;