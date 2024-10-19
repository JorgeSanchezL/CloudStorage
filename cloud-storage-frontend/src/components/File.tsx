import React, { useContext, useState } from 'react';
import './File.css';
import { ModalContext } from '../modal/modalProvider'; // Adjust the import path as necessary

import FileInfoCard from './FileInfoCard';
import ReactDOM from 'react-dom';

type FileInfo = {
    Name: string;
    Size: number;
    LastModification: string;
};

interface FileProps {
    fileInfo: FileInfo;
    onDownloadClick: () => void;
}

const File: React.FC<FileProps> = ({ fileInfo, onDownloadClick }) => {

    const modalContext = useContext(ModalContext);

    if (!modalContext) {
        throw new Error("ModalContext must be used within a ModalProvider");
    }

    const {isModalOpen, openModal, closeModal} = modalContext;

    const [isHovered, setIsHovered] = useState(false);

    return (
        <div 
            className="file-container"
            onMouseEnter={() => setIsHovered(true)}
            onMouseLeave={() => setIsHovered(false)}
            onClick={isModalOpen ? closeModal : () => openModal(
                <FileInfoCard 
                    fileInfo={fileInfo} 
                    onClose={() => closeModal()}
                />)
            }>

            <div className="file-icon">📄</div>
            <div className="file-text">{fileInfo.Name}</div>
            {isHovered && (
                <button className="download-button" onClick={(e) => {
                    e.stopPropagation();
                    onDownloadClick();
                }}>
                    ⬇️
                </button>
            )}
        </div>
    );
};

export default File;