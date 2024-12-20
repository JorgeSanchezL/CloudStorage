import React from 'react';
import './DeleteConfirmationCard.css';

interface DeleteConfirmationCardProps {
    fileName: string;
    onAccept: () => void;
    onCancel: () => void;
}

const DeleteConfirmationCard: React.FC<DeleteConfirmationCardProps> = ({ fileName, onAccept, onCancel }) => {
    return (
        <div className="modal-overlay">
            <div className="modal-content">
                <h1 className="modal-title">Confirm Deletion</h1>
                <p>Are you sure you want to delete the file <strong>{fileName}</strong>?</p>
                <div className="button-group">
                    <button className="cancel-button" onClick={onCancel}>Cancel</button>
                    <button className="accept-button" onClick={onAccept}>Accept</button>
                </div>
            </div>
        </div>
    );
}

export default DeleteConfirmationCard;