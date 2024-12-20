import React, { useContext, useState } from 'react';
import { ModalContext } from '../modal/modalProvider';
import './File.css';

import DeleteConfirmationCard from './DeleteConfirmationCard';
import FileInfoCard from './FileInfoCard';

type FileInfo = {
    Name: string;
    Size: number;
    LastModification: string;
};

interface FileProps {
    fileInfo: FileInfo;
    onDownloadClick: () => void;
    onRemoveClick: () => void;
}

const File: React.FC<FileProps> = ({ fileInfo, onDownloadClick, onRemoveClick }) => {

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
            onClick={(e) => {
                e.stopPropagation();
                isModalOpen ? closeModal : () => openModal(
                    <FileInfoCard 
                        fileInfo={fileInfo} 
                        onClose={() => closeModal()}
                    />)
                }
            }>

            <div className="file-icon">📄</div>
            <div className="file-text">{fileInfo.Name}</div>
            {isHovered && (
                <button 
                    className="remove-button"
                    onClick={(e) => {
                        e.stopPropagation();
                        openModal(
                            <DeleteConfirmationCard 
                                fileName={fileInfo.Name}
                                onAccept={() => {
                                    onRemoveClick()
                                    closeModal()
                                }} 
                                onCancel={() => closeModal()}
                            />
                        );
                    }}>
                    🗑️
                </button>
            )}
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